// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package gracefulexit

import (
	"context"

	"github.com/zeebo/errs"
	"go.uber.org/zap"

	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/satellite/metainfo"
)

var _ metainfo.Observer = (*PathCollector)(nil)

// PathCollector uses the metainfo loop to add paths to node reservoirs
//
// architecture: Observer
type PathCollector struct {
	db              DB
	metainfoService *metainfo.Service
	nodeIDMap       map[storj.NodeID]struct{}
	buffer          []TransferQueueItem
	log             *zap.Logger
	batchSize       int
}

// NewPathCollector instantiates a path collector.
func NewPathCollector(db DB, metainfoService *metainfo.Service, nodeIDs storj.NodeIDList, log *zap.Logger, batchSize int) *PathCollector {
	buffer := make([]TransferQueueItem, 0, batchSize)
	collector := &PathCollector{
		db:              db,
		metainfoService: metainfoService,
		log:             log,
		buffer:          buffer,
		batchSize:       batchSize,
	}

	if len(nodeIDs) > 0 {
		collector.nodeIDMap = make(map[storj.NodeID]struct{}, len(nodeIDs))
		for _, nodeID := range nodeIDs {
			collector.nodeIDMap[nodeID] = struct{}{}
		}
	}

	return collector
}

// Flush persists the current buffer items to the database.
func (collector *PathCollector) Flush(ctx context.Context) (err error) {
	return collector.flush(ctx, 1)
}

// RemoteSegment takes a remote segment found in metainfo and creates a graceful exit transfer queue item if it doesn't exist already
func (collector *PathCollector) RemoteSegment(ctx context.Context, path metainfo.ScopedPath, pointer *pb.Pointer) (err error) {
	if len(collector.nodeIDMap) == 0 {
		return nil
	}

	numPieces := int32(len(pointer.GetRemote().GetRemotePieces()))
	for _, piece := range pointer.GetRemote().GetRemotePieces() {
		_, ok := collector.nodeIDMap[piece.NodeId]
		if ok {
			item := TransferQueueItem{
				NodeID:          piece.NodeId,
				Path:            []byte(path.Raw),
				PieceNum:        piece.PieceNum,
				DurabilityRatio: 1.0,
			}
			collector.log.Debug("adding piece to transfer queue.",
				zap.String("path", path.Raw), zap.Int32("piece num", piece.GetPieceNum()),
				zap.Int32("num pieces", numPieces), zap.Int32("required pieces", pointer.GetRemote().GetRedundancy().GetSuccessThreshold()))

			collector.buffer = append(collector.buffer, item)
			err = collector.flush(ctx, collector.batchSize)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Object returns nil because the audit service does not interact with objects
func (collector *PathCollector) Object(ctx context.Context, path metainfo.ScopedPath, pointer *pb.Pointer) (err error) {
	return nil
}

// InlineSegment returns nil because we're only auditing for storage nodes for now
func (collector *PathCollector) InlineSegment(ctx context.Context, path metainfo.ScopedPath, pointer *pb.Pointer) (err error) {
	return nil
}

func (collector *PathCollector) flush(ctx context.Context, limit int) (err error) {
	if len(collector.buffer) >= limit {
		err = collector.db.Enqueue(ctx, collector.buffer)
		collector.buffer = collector.buffer[:0]

		return errs.Wrap(err)
	}
	return nil
}
