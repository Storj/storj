// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package gc

import (
	"context"
	"time"

	"go.uber.org/zap"

	"storj.io/storj/internal/memory"
	"storj.io/storj/pkg/bloomfilter"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
)

// Observer implements the observer interface for gc
type Observer struct {
	log          *zap.Logger
	config       Config
	creationDate time.Time
	pieceCounts  map[storj.NodeID]int

	retainInfos map[storj.NodeID]*RetainInfo
}

// NewObserver instantiates a gc Observer
func NewObserver(log *zap.Logger, config Config, pieceCounts map[storj.NodeID]int) *Observer {
	return &Observer{
		log:          log,
		config:       config,
		creationDate: time.Now().UTC(),
		pieceCounts:  pieceCounts,

		retainInfos: make(map[storj.NodeID]*RetainInfo),
	}
}

// RemoteSegment takes a remote segment found in metainfo and adds pieces to bloom filters
func (observer *Observer) RemoteSegment(ctx context.Context, path storj.Path, pointer *pb.Pointer) (err error) {
	defer mon.Task()(&ctx)(&err)

	remote := pointer.GetRemote()
	pieces := remote.GetRemotePieces()

	for _, piece := range pieces {
		pieceID := remote.RootPieceId.Derive(piece.NodeId, piece.PieceNum)
		observer.add(ctx, piece.NodeId, pieceID)
	}
	return nil
}

// RemoteObject returns nil because gc does not interact with remote objects
func (observer *Observer) RemoteObject(ctx context.Context, path storj.Path, pointer *pb.Pointer) (err error) {
	defer mon.Task()(&ctx)(&err)
	return nil
}

// InlineSegment returns nil because we're only doing gc for storage nodes for now
func (observer *Observer) InlineSegment(ctx context.Context, path storj.Path, pointer *pb.Pointer) (err error) {
	defer mon.Task()(&ctx)(&err)
	return nil
}

// adds a pieceID to the relevant node's RetainInfo
func (observer *Observer) add(ctx context.Context, nodeID storj.NodeID, pieceID storj.PieceID) {
	var filter *bloomfilter.Filter

	if _, ok := observer.retainInfos[nodeID]; !ok {
		// If we know how many pieces a node should be storing, use that number. Otherwise use default.
		numPieces := observer.config.InitialPieces
		if observer.pieceCounts[nodeID] > 0 {
			numPieces = observer.pieceCounts[nodeID]
		}
		// limit size of bloom filter to ensure we are under the limit for GRPC
		filter = bloomfilter.NewOptimalMaxSize(numPieces, observer.config.FalsePositiveRate, 2*memory.MiB)
		observer.retainInfos[nodeID] = &RetainInfo{
			Filter:       filter,
			CreationDate: observer.creationDate,
		}
	}

	observer.retainInfos[nodeID].Filter.Add(pieceID)
	observer.retainInfos[nodeID].Count++
}
