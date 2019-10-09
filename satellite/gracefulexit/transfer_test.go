// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package gracefulexit_test

import (
	"io"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeebo/errs"
	"google.golang.org/grpc"

	"storj.io/storj/internal/memory"
	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/internal/testrand"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/uplink"
)

const numMessages = 6

func TestInitiate(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 8,
		UplinkCount:      1,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		uplinkPeer := planet.Uplinks[0]
		satellite := planet.Satellites[0]
		exitingNode := planet.StorageNodes[0]

		satellite.GracefulExit.Chore.Loop.Pause()

		rs := &uplink.RSConfig{
			MinThreshold:     4,
			RepairThreshold:  6,
			SuccessThreshold: 8,
			MaxThreshold:     8,
		}

		for i := 0; i < numMessages; i++ {
			err := uplinkPeer.UploadWithConfig(ctx, satellite, rs, "testbucket", "test/path"+strconv.Itoa(i), testrand.Bytes(5*memory.KiB))
			require.NoError(t, err)
		}
		// check that there are no exiting nodes.
		exitingNodeIDs, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodeIDs, 0)

		conn, err := exitingNode.Dialer.DialAddressID(ctx, satellite.Addr(), satellite.Identity.ID)
		require.NoError(t, err)
		defer func() {
			err = errs.Combine(err, conn.Close())
		}()

		client := conn.SatelliteGracefulExitClient()

		c, err := client.Process(ctx, grpc.EmptyCallOption{})
		require.NoError(t, err)

		response, err := c.Recv()
		require.NoError(t, err)

		switch response.GetMessage().(type) {
		case *pb.SatelliteMessage_NotReady:
			// now check that the exiting node is initiated.
			exitingNodeIDs, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
			require.NoError(t, err)
			require.Len(t, exitingNodeIDs, 1)

			require.Equal(t, exitingNode.ID(), exitingNodeIDs[0])
		default:
			t.FailNow()
		}
		// close the old client
		require.NoError(t, c.CloseSend())

		// trigger the metainfo loop chore so we can get some pieces to transfer
		satellite.GracefulExit.Chore.Loop.TriggerWait()

		incompleteTransfers, err := satellite.DB.GracefulExit().GetIncomplete(ctx, exitingNode.ID(), numMessages+1, 0)
		require.NoError(t, err)
		require.Len(t, incompleteTransfers, numMessages)

		c, err = client.Process(ctx, grpc.EmptyCallOption{})
		require.NoError(t, err)
		defer func() {
			err = errs.Combine(err, c.CloseSend())
		}()

		var pieceID storj.PieceID
		for {
			response, err = c.Recv()
			if err == io.EOF {
				// Done
				break
			}
			require.NoError(t, err)

			switch m := response.GetMessage().(type) {
			case *pb.SatelliteMessage_TransferPiece:
				require.NotNil(t, m)

				// pick the first one to fail
				if pieceID.IsZero() {
					pieceID = m.TransferPiece.PieceId
				}

				if pieceID != m.TransferPiece.PieceId {
					success := &pb.StorageNodeMessage{
						Message: &pb.StorageNodeMessage_Succeeded{
							Succeeded: &pb.TransferSucceeded{
								OriginalPieceHash: &pb.PieceHash{PieceId: m.TransferPiece.PieceId},
								AddressedOrderLimit: &pb.AddressedOrderLimit{
									Limit: &pb.OrderLimit{
										PieceId: m.TransferPiece.PieceId,
									},
								},
							},
						},
					}
					c.Send(success)
				} else {
					failed := &pb.StorageNodeMessage{
						Message: &pb.StorageNodeMessage_Failed{
							Failed: &pb.TransferFailed{
								PieceId: m.TransferPiece.PieceId,
								Error:   pb.TransferFailed_UNKNOWN,
							},
						},
					}
					c.Send(failed)
				}
			case *pb.SatelliteMessage_ExitCompleted:
				break
			default:
				t.FailNow()
			}
		}

		progress, err := satellite.DB.GracefulExit().GetProgress(ctx, exitingNode.ID())
		require.NoError(t, err)

		require.Equal(t, int64(numMessages), progress.PiecesTransferred)
		// even though we failed 1, it eventually succeeded, so the count should be 0
		require.Equal(t, int64(0), progress.PiecesFailed)

	})
}
