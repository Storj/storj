// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package segments

import (
	"context"
	"time"

	"github.com/zeebo/errs"

	"storj.io/storj/pkg/eestream"
	"storj.io/storj/pkg/identity"
	"storj.io/storj/pkg/overlay"
	"storj.io/storj/pkg/pb"
	ecclient "storj.io/storj/pkg/storage/ec"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/satellite/metainfo"
	"storj.io/storj/satellite/orders"
)

// Repairer for segments
type Repairer struct {
	metainfo *metainfo.Service
	orders   *orders.Service
	cache    *overlay.Cache
	ec       ecclient.Client
	identity *identity.FullIdentity
	timeout  time.Duration
}

// NewSegmentRepairer creates a new instance of SegmentRepairer
func NewSegmentRepairer(metainfo *metainfo.Service, orders *orders.Service, cache *overlay.Cache, ec ecclient.Client, identity *identity.FullIdentity, timeout time.Duration) *Repairer {
	return &Repairer{
		metainfo: metainfo,
		orders:   orders,
		cache:    cache,
		ec:       ec,
		identity: identity,
		timeout:  timeout,
	}
}

// Repair retrieves an at-risk segment and repairs and stores lost pieces on new nodes
func (repairer *Repairer) Repair(ctx context.Context, path storj.Path) (err error) {
	defer mon.Task()(&ctx)(&err)

	// Read the segment pointer from the metainfo
	pointer, err := repairer.metainfo.Get(path)
	if err != nil {
		return Error.Wrap(err)
	}

	if pointer.GetType() != pb.Pointer_REMOTE {
		return Error.New("cannot repair inline segment %s", path)
	}

	redundancy, err := eestream.NewRedundancyStrategyFromProto(pointer.GetRemote().GetRedundancy())
	if err != nil {
		return Error.Wrap(err)
	}

	pieceSize := eestream.CalcPieceSize(pointer.GetSegmentSize(), redundancy)
	expiration := pointer.GetExpirationDate()

	var excludeNodeIDs storj.NodeIDList
	var healthyPieces []*pb.RemotePiece
	pieces := pointer.GetRemote().GetRemotePieces()
	missingPieces, err := repairer.cache.GetMissingPieces(ctx, pieces)
	if err != nil {
		return Error.New("error getting missing pieces %s", err)
	}

	numHealthy := len(pieces) - len(missingPieces)
	// irreparable piece
	if int32(numHealthy) < pointer.Remote.Redundancy.MinReq {
		return Error.New("piece %v cannot be repaired", path)
	}

	// repair not needed
	if int32(numHealthy) > pointer.Remote.Redundancy.RepairThreshold {
		return Error.New("piece %v with %d pieces above repiar threshold %d", path, numHealthy, pointer.Remote.Redundancy.RepairThreshold)
	}

	lostPiecesSet := sliceToSet(missingPieces)

	// Populate healthyPieces with all pieces from the pointer except those correlating to indices in lostPieces
	for _, piece := range pieces {
		excludeNodeIDs = append(excludeNodeIDs, piece.NodeId)
		if _, ok := lostPiecesSet[piece.GetPieceNum()]; !ok {
			healthyPieces = append(healthyPieces, piece)
		}
	}

	bucketID, err := createBucketID(path)
	if err != nil {
		return Error.Wrap(err)
	}

	// Create the order limits for the GET_REPAIR action
	getOrderLimits, err := repairer.orders.CreateGetRepairOrderLimits(ctx, repairer.identity.PeerIdentity(), bucketID, pointer, healthyPieces)
	if err != nil {
		return Error.Wrap(err)
	}

	// Request Overlay for n-h new storage nodes
	request := overlay.FindStorageNodesRequest{
		RequestedCount: redundancy.TotalCount() - len(healthyPieces),
		FreeBandwidth:  pieceSize,
		FreeDisk:       pieceSize,
		ExcludedNodes:  excludeNodeIDs,
	}
	newNodes, err := repairer.cache.FindStorageNodes(ctx, request)
	if err != nil {
		return Error.Wrap(err)
	}

	// Create the order limits for the PUT_REPAIR action
	putLimits, err := repairer.orders.CreatePutRepairOrderLimits(ctx, repairer.identity.PeerIdentity(), bucketID, pointer, getOrderLimits, newNodes)
	if err != nil {
		return Error.Wrap(err)
	}

	// Download the segment using just the healthy pieces
	rr, err := repairer.ec.Get(ctx, getOrderLimits, redundancy, pointer.GetSegmentSize())
	if err != nil {
		return Error.Wrap(err)
	}

	r, err := rr.Range(ctx, 0, rr.Size())
	if err != nil {
		return Error.Wrap(err)
	}
	defer func() { err = errs.Combine(err, r.Close()) }()

	// Upload the repaired pieces
	successfulNodes, hashes, err := repairer.ec.Repair(ctx, putLimits, redundancy, r, convertTime(expiration), repairer.timeout, path)
	if err != nil {
		return Error.Wrap(err)
	}

	// Add the successfully uploaded pieces to the healthyPieces
	for i, node := range successfulNodes {
		if node == nil {
			continue
		}
		healthyPieces = append(healthyPieces, &pb.RemotePiece{
			PieceNum: int32(i),
			NodeId:   node.Id,
			Hash:     hashes[i],
		})
	}

	// Update the remote pieces in the pointer
	pointer.GetRemote().RemotePieces = healthyPieces

	// Update the segment pointer in the metainfo
	return repairer.metainfo.Put(path, pointer)
}

// sliceToSet converts the given slice to a set
func sliceToSet(slice []int32) map[int32]struct{} {
	set := make(map[int32]struct{}, len(slice))
	for _, value := range slice {
		set[value] = struct{}{}
	}
	return set
}

func createBucketID(path storj.Path) ([]byte, error) {
	comps := storj.SplitPath(path)
	if len(comps) < 3 {
		return nil, Error.New("no bucket component in path: %s", path)
	}
	return []byte(storj.JoinPaths(comps[0], comps[2])), nil
}
