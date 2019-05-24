// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package audit_test

import (
	"crypto/rand"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"storj.io/storj/internal/memory"
	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/pkg/audit"
	"storj.io/storj/pkg/auth/signing"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/pkcrypto"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/uplink"
)

// This is a bulky test but all it's doing is:
// - uploading random data
// - using the cursor to get a stripe
// - creating pending audits for all nodes holding pieces for that stripe
//     - the actual shares are downloaded to make sure ExpectedShareHash is correct
// - calling reverify on that same stripe
// - expect all six storage nodes to be marked as successes in the audit report
func TestReverifySuccess(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 6, UplinkCount: 1,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		err := planet.Satellites[0].Audit.Service.Close()
		require.NoError(t, err)

		ul := planet.Uplinks[0]
		testData := make([]byte, 1*memory.MiB)
		_, err = rand.Read(testData)
		require.NoError(t, err)

		err = ul.UploadWithConfig(ctx, planet.Satellites[0], &uplink.RSConfig{
			MinThreshold:     4,
			RepairThreshold:  5,
			SuccessThreshold: 6,
			MaxThreshold:     6,
		}, "testbucket", "test/path", testData)
		require.NoError(t, err)

		metainfo := planet.Satellites[0].Metainfo.Service
		cursor := audit.NewCursor(metainfo)

		var stripe *audit.Stripe
		stripe, _, err = cursor.NextStripe(ctx)
		require.NoError(t, err)
		require.NotNil(t, stripe)

		overlay := planet.Satellites[0].Overlay.Service
		transport := planet.Satellites[0].Transport
		orders := planet.Satellites[0].Orders.Service
		containment := planet.Satellites[0].DB.Containment()
		minBytesPerSecond := 128 * memory.B
		reporter := audit.NewReporter(overlay, containment, 1)
		verifier := audit.NewVerifier(zap.L(), reporter, transport, overlay, containment, orders, planet.Satellites[0].Identity, minBytesPerSecond)
		require.NotNil(t, verifier)

		for _, piece := range stripe.Segment.GetRemote().GetRemotePieces() {
			rootPieceID := stripe.Segment.GetRemote().RootPieceId
			redundancy := stripe.Segment.GetRemote().GetRedundancy()

			orderLimit, err := signing.SignOrderLimit(signing.SignerFromFullIdentity(planet.Satellites[0].Identity), &pb.OrderLimit2{
				SerialNumber:    storj.SerialNumber{},
				SatelliteId:     planet.Satellites[0].ID(),
				UplinkId:        ul.ID(),
				StorageNodeId:   piece.NodeId,
				PieceId:         rootPieceID.Derive(piece.NodeId),
				Action:          pb.PieceAction_GET_AUDIT,
				Limit:           int64(redundancy.ErasureShareSize),
				PieceExpiration: &timestamp.Timestamp{Seconds: time.Now().Unix() + 3000},
				OrderExpiration: &timestamp.Timestamp{Seconds: time.Now().Unix() + 3000},
			})
			require.NoError(t, err)

			var nodeAddr *pb.NodeAddress

			for i := range planet.StorageNodes {
				if planet.StorageNodes[i].ID() == piece.NodeId {
					nodeAddr = &pb.NodeAddress{
						Address:   planet.StorageNodes[i].Addr(),
						Transport: pb.NodeTransport_TCP_TLS_GRPC,
					}
				}
			}

			limit := &pb.AddressedOrderLimit{
				Limit:              orderLimit,
				StorageNodeAddress: nodeAddr,
			}

			share, err := verifier.GetShare(ctx, limit, stripe.Index, redundancy.ErasureShareSize, int(piece.PieceNum))
			require.NoError(t, err)

			pending := &audit.PendingAudit{
				NodeID:            piece.NodeId,
				PieceID:           stripe.Segment.GetRemote().RootPieceId,
				StripeIndex:       stripe.Index,
				ShareSize:         stripe.Segment.GetRemote().GetRedundancy().ErasureShareSize,
				ExpectedShareHash: pkcrypto.SHA256Hash(share.Data),
				ReverifyCount:     0,
			}

			err = planet.Satellites[0].DB.Containment().IncrementPending(ctx, pending)
			require.NoError(t, err)
		}

		report, err := verifier.Reverify(ctx, stripe)
		require.NoError(t, err)

		successes := make(map[string]bool)
		for _, nodeID := range report.Successes {
			successes[nodeID.String()] = true
		}

		for _, piece := range stripe.Segment.GetRemote().GetRemotePieces() {
			require.True(t, successes[piece.NodeId.String()])
		}
	})
}

func TestContainIncrementAndGet(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 4,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {

		randBytes := make([]byte, 10)
		_, err := rand.Read(randBytes)
		require.NoError(t, err)
		someHash := pkcrypto.SHA256Hash(randBytes)

		input := &audit.PendingAudit{
			NodeID:            planet.StorageNodes[0].ID(),
			PieceID:           storj.PieceID{},
			StripeIndex:       0,
			ShareSize:         0,
			ExpectedShareHash: someHash,
			ReverifyCount:     0,
		}

		err = planet.Satellites[0].DB.Containment().IncrementPending(ctx, input)
		require.NoError(t, err)

		output, err := planet.Satellites[0].DB.Containment().Get(ctx, input.NodeID)
		require.NoError(t, err)

		require.Equal(t, input, output)

		// check contained flag set to true
		node, err := planet.Satellites[0].DB.OverlayCache().Get(ctx, input.NodeID)
		require.NoError(t, err)
		require.True(t, node.Contained)

		nodeID1 := planet.StorageNodes[1].ID()
		_, err = planet.Satellites[0].DB.Containment().Get(ctx, nodeID1)
		require.Error(t, err, audit.ErrContainedNotFound.New(nodeID1.String()))
		require.True(t, audit.ErrContainedNotFound.Has(err))
	})
}

func TestContainIncrementPendingEntryExists(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 4,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {

		randBytes := make([]byte, 10)
		_, err := rand.Read(randBytes)
		require.NoError(t, err)
		hash1 := pkcrypto.SHA256Hash(randBytes)

		info1 := &audit.PendingAudit{
			NodeID:            planet.StorageNodes[0].ID(),
			PieceID:           storj.PieceID{},
			StripeIndex:       0,
			ShareSize:         0,
			ExpectedShareHash: hash1,
			ReverifyCount:     0,
		}

		err = planet.Satellites[0].DB.Containment().IncrementPending(ctx, info1)
		require.NoError(t, err)

		randBytes = make([]byte, 10)
		_, err = rand.Read(randBytes)
		require.NoError(t, err)
		hash2 := pkcrypto.SHA256Hash(randBytes)

		info2 := &audit.PendingAudit{
			NodeID:            info1.NodeID,
			PieceID:           storj.PieceID{},
			StripeIndex:       1,
			ShareSize:         1,
			ExpectedShareHash: hash2,
			ReverifyCount:     0,
		}

		// expect failure when an entry with the same nodeID but different expected share data already exists
		err = planet.Satellites[0].DB.Containment().IncrementPending(ctx, info2)
		require.Error(t, err)
		require.True(t, audit.ErrAlreadyExists.Has(err))

		// expect reverify count for an entry to be 0 after first IncrementPending call
		pending, err := planet.Satellites[0].DB.Containment().Get(ctx, info1.NodeID)
		require.NoError(t, err)
		require.EqualValues(t, 0, pending.ReverifyCount)

		// expect reverify count to be 1 after second IncrementPending call
		err = planet.Satellites[0].DB.Containment().IncrementPending(ctx, info1)
		require.NoError(t, err)
		pending, err = planet.Satellites[0].DB.Containment().Get(ctx, info1.NodeID)
		require.NoError(t, err)
		require.EqualValues(t, 1, pending.ReverifyCount)
	})
}

func TestContainDelete(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 4,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {

		randBytes := make([]byte, 10)
		_, err := rand.Read(randBytes)
		require.NoError(t, err)
		hash1 := pkcrypto.SHA256Hash(randBytes)

		info1 := &audit.PendingAudit{
			NodeID:            planet.StorageNodes[0].ID(),
			PieceID:           storj.PieceID{},
			StripeIndex:       0,
			ShareSize:         0,
			ExpectedShareHash: hash1,
			ReverifyCount:     0,
		}

		err = planet.Satellites[0].DB.Containment().IncrementPending(ctx, info1)
		require.NoError(t, err)

		// check contained flag set to true
		node, err := planet.Satellites[0].DB.OverlayCache().Get(ctx, info1.NodeID)
		require.NoError(t, err)
		require.True(t, node.Contained)

		isDeleted, err := planet.Satellites[0].DB.Containment().Delete(ctx, info1.NodeID)
		require.NoError(t, err)
		require.True(t, isDeleted)

		// check contained flag set to false
		node, err = planet.Satellites[0].DB.OverlayCache().Get(ctx, info1.NodeID)
		require.NoError(t, err)
		require.False(t, node.Contained)

		// get pending audit that doesn't exist
		_, err = planet.Satellites[0].DB.Containment().Get(ctx, info1.NodeID)
		require.Error(t, err, audit.ErrContainedNotFound.New(info1.NodeID.String()))
		require.True(t, audit.ErrContainedNotFound.Has(err))

		// delete pending audit that doesn't exist
		isDeleted, err = planet.Satellites[0].DB.Containment().Delete(ctx, info1.NodeID)
		require.NoError(t, err)
		require.False(t, isDeleted)
	})
}
