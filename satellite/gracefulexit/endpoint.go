// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package gracefulexit

import (
	"context"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"storj.io/storj/pkg/identity"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/rpc/rpcstatus"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/satellite/metainfo"
	"storj.io/storj/satellite/orders"
	"storj.io/storj/satellite/overlay"
	"storj.io/storj/uplink/eestream"
)

// millis for the transfer queue building ticker
const buildQueueMillis = 100

// drpcEndpoint wraps streaming methods so that they can be used with drpc
type drpcEndpoint struct{ *Endpoint }

// processStream is the minimum interface required to process requests.
type processStream interface {
	Context() context.Context
	Send(*pb.SatelliteMessage) error
	Recv() (*pb.StorageNodeMessage, error)
}

// Endpoint for handling the transfer of pieces for Graceful Exit.
type Endpoint struct {
	log       *zap.Logger
	interval  time.Duration
	db        DB
	overlaydb overlay.DB
	overlay   *overlay.Service
	metainfo  *metainfo.Service
	orders    *orders.Service
	config    Config
}

type pendingTransfer struct {
	path             []byte
	pieceSize        int64
	satelliteMessage *pb.SatelliteMessage
}

// DRPC returns a DRPC form of the endpoint.
func (endpoint *Endpoint) DRPC() pb.DRPCSatelliteGracefulExitServer {
	return &drpcEndpoint{Endpoint: endpoint}
}

// NewEndpoint creates a new graceful exit endpoint.
func NewEndpoint(log *zap.Logger, db DB, overlaydb overlay.DB, overlay *overlay.Service, metainfo *metainfo.Service, orders *orders.Service, config Config) *Endpoint {
	return &Endpoint{
		log:       log,
		interval:  time.Millisecond * buildQueueMillis,
		db:        db,
		overlaydb: overlaydb,
		overlay:   overlay,
		metainfo:  metainfo,
		orders:    orders,
		config:    config,
	}
}

// Process is called by storage nodes to receive pieces to transfer to new nodes and get exit status.
func (endpoint *Endpoint) Process(stream pb.SatelliteGracefulExit_ProcessServer) error {
	return endpoint.doProcess(stream)
}

// Process is called by storage nodes to receive pieces to transfer to new nodes and get exit status.
func (endpoint *drpcEndpoint) Process(stream pb.DRPCSatelliteGracefulExit_ProcessStream) error {
	return endpoint.doProcess(stream)
}

func (endpoint *Endpoint) doProcess(stream processStream) (err error) {
	ctx := stream.Context()
	defer mon.Task()(&ctx)(&err)

	peer, err := identity.PeerIdentityFromContext(ctx)
	if err != nil {
		return rpcstatus.Error(rpcstatus.Unauthenticated, err.Error())
	}

	nodeID := peer.ID
	endpoint.log.Debug("graceful exit process.", zap.String("nodeID", nodeID.String()))

	eofHandler := func(err error) error {
		if err == io.EOF {
			endpoint.log.Debug("received EOF when trying to receive messages from storage node.", zap.String("nodeID", nodeID.String()))
			return nil
		}
		return rpcstatus.Error(rpcstatus.Unknown, err.Error())
	}

	exitStatus, err := endpoint.overlaydb.GetExitStatus(ctx, nodeID)
	if err != nil {
		return Error.Wrap(err)
	}

	if exitStatus.ExitFinishedAt != nil {
		// TODO revisit this. Should check if signature was sent
		completed := &pb.SatelliteMessage{Message: &pb.SatelliteMessage_ExitCompleted{ExitCompleted: &pb.ExitCompleted{}}}
		err = stream.Send(completed)
		return Error.Wrap(err)
	}

	if exitStatus.ExitInitiatedAt == nil {
		request := &overlay.ExitStatusRequest{NodeID: nodeID, ExitInitiatedAt: time.Now().UTC()}
		_, err = endpoint.overlaydb.UpdateExitStatus(ctx, request)
		if err != nil {
			return Error.Wrap(err)
		}

		err = stream.Send(&pb.SatelliteMessage{Message: &pb.SatelliteMessage_NotReady{NotReady: &pb.NotReady{}}})
		return Error.Wrap(err)
	}

	if exitStatus.ExitLoopCompletedAt == nil {
		err = stream.Send(&pb.SatelliteMessage{Message: &pb.SatelliteMessage_NotReady{NotReady: &pb.NotReady{}}})
		return Error.Wrap(err)
	}

	// TODO possibly switch out for custom queue
	var pending sync.Map
	pendingLength := func() int {
		count := 0
		pending.Range(func(key interface{}, value interface{}) bool {
			count++
			return true
		})

		return count
	}
	// this will be 1 until GetIncomplete* methods no longer return values
	var morePiecesFlag int32 = 1

	var group errgroup.Group
	group.Go(func() error {
		ticker := time.NewTicker(endpoint.interval)
		defer ticker.Stop()

		for range ticker.C {
			if pendingLength() == 0 {
				incomplete, err := endpoint.db.GetIncompleteNotFailed(ctx, nodeID, endpoint.config.EndpointBatchSize, 0)
				if err != nil {
					return Error.Wrap(err)
				}

				if len(incomplete) == 0 {
					incomplete, err = endpoint.db.GetIncompleteFailed(ctx, nodeID, endpoint.config.EndpointMaxFailures+1, endpoint.config.EndpointBatchSize, 0)
					if err != nil {
						return Error.Wrap(err)
					}
				}

				if len(incomplete) == 0 {
					endpoint.log.Debug("no more pieces to transfer for node.", zap.String("node ID", nodeID.String()))
					atomic.StoreInt32(&morePiecesFlag, 0)
					break
				}

				for _, inc := range incomplete {
					pointer, err := endpoint.metainfo.Get(ctx, string(inc.Path))
					if err != nil {
						return Error.Wrap(err)
					}
					remote := pointer.GetRemote()

					found := false
					for _, piece := range remote.GetRemotePieces() {
						if piece.NodeId == nodeID && piece.PieceNum == inc.PieceNum {
							found = true
						}
					}
					if !found {
						endpoint.log.Debug("piece no longer held by node.", zap.String("node ID", nodeID.String()), zap.ByteString("path", inc.Path), zap.Int32("piece num", inc.PieceNum))
						continue
					}

					redundancy, err := eestream.NewRedundancyStrategyFromProto(pointer.GetRemote().GetRedundancy())
					if err != nil {
						return Error.Wrap(err)
					}
					pieceSize := eestream.CalcPieceSize(pointer.GetSegmentSize(), redundancy)

					request := overlay.FindStorageNodesRequest{
						RequestedCount: 1,
						FreeBandwidth:  pieceSize,
						FreeDisk:       pieceSize,
						ExcludedNodes:  []storj.NodeID{nodeID},
					}

					newNodes, err := endpoint.overlay.FindStorageNodes(ctx, request)
					if err != nil {
						return Error.Wrap(err)
					}

					if len(newNodes) == 0 {
						return Error.New("could not find a node to transfer this piece to. nodeID %v, path %v, pieceNum %v.", nodeID.String(), zap.ByteString("path", inc.Path), inc.PieceNum)
					}

					endpoint.log.Debug("found new node for piece transfer.", zap.String("node ID", newNodes[0].Id.String()), zap.ByteString("path", inc.Path), zap.Int32("piece num", inc.PieceNum))

					parts := storj.SplitPath(string(inc.Path))
					if len(parts) < 2 {
						return Error.New("invalid path %v.", zap.ByteString("path", inc.Path))
					}

					bucketID := []byte(storj.JoinPaths(parts[0], parts[1]))
					limit, privateKey, err := endpoint.orders.CreateGracefulExitPutOrderLimit(ctx, bucketID, newNodes[0].Id, inc.PieceNum, remote.RootPieceId, remote.Redundancy.GetErasureShareSize())
					if err != nil {
						return Error.Wrap(err)
					}

					transferMsg := &pb.SatelliteMessage{
						Message: &pb.SatelliteMessage_TransferPiece{
							TransferPiece: &pb.TransferPiece{
								PieceId:             limit.Limit.PieceId,
								AddressedOrderLimit: limit,
								PrivateKey:          privateKey,
							},
						},
					}
					err = stream.Send(transferMsg)
					if err != nil {
						return Error.Wrap(err)
					}
					pending.Store(limit.Limit.PieceId, pendingTransfer{
						path:             inc.Path,
						pieceSize:        pieceSize,
						satelliteMessage: transferMsg,
					})
				}
			}
		}
		return nil
	})

	for {
		// if there are no more transfers and the pending queue is empty, send complete
		if atomic.LoadInt32(&morePiecesFlag) == 0 && pendingLength() == 0 {
			// TODO check whether failure threshold is met before sending completed
			// TODO needs exit signature
			transferMsg := &pb.SatelliteMessage{
				Message: &pb.SatelliteMessage_ExitCompleted{
					ExitCompleted: &pb.ExitCompleted{},
				},
			}
			err = stream.Send(transferMsg)
			if err != nil {
				return Error.Wrap(err)
			}
			break
		}
		// skip if there are none pending
		if pendingLength() == 0 {
			continue
		}

		request, err := stream.Recv()
		if err != nil {
			return eofHandler(err)
		}

		switch m := request.GetMessage().(type) {
		case *pb.StorageNodeMessage_Succeeded:
			err = endpoint.handleSucceeded(ctx, &pending, nodeID, m)
			if err != nil {
				return Error.Wrap(err)
			}

		case *pb.StorageNodeMessage_Failed:
			err = endpoint.handleFailed(ctx, &pending, nodeID, m)
			if err != nil {
				return Error.Wrap(err)
			}
		default:
			return Error.New("unknown storage node message: %v", m)
		}
	}

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}

func (endpoint *Endpoint) sendPiecesToTransfer(ctx context.Context, stream processStream, nodeID storj.NodeID) (err error) {
	defer mon.Task()(&ctx)(&err)
	endpoint.log.Debug("sending pieces to transfer.", zap.String("nodeID", nodeID.String()))

	return nil
}

func (endpoint *Endpoint) handleSucceeded(ctx context.Context, pending *sync.Map, nodeID storj.NodeID, message *pb.StorageNodeMessage_Succeeded) (err error) {
	defer mon.Task()(&ctx)(&err)
	if message.Succeeded.GetAddressedOrderLimit() == nil {
		return Error.New("Addressed order limit cannot be nil.")
	}
	if message.Succeeded.GetOriginalPieceHash() == nil {
		return Error.New("Original piece hash cannot be nil.")
	}
	endpoint.log.Debug("transfer succeeded.", zap.String("piece ID", message.Succeeded.AddressedOrderLimit.Limit.PieceId.String()))

	// TODO validation

	transfer, ok := pending.Load(message.Succeeded.OriginalPieceHash.PieceId)
	if !ok {
		endpoint.log.Debug("could not find transfer message in pending queue. skipping .", zap.String("piece ID", message.Succeeded.AddressedOrderLimit.Limit.PieceId.String()))
	}

	transferQueueItem, err := endpoint.db.GetTransferQueueItem(ctx, nodeID, transfer.(pendingTransfer).path)
	if err != nil {
		return Error.Wrap(err)
	}

	var failed int64
	if transferQueueItem.FailedCount != nil && *transferQueueItem.FailedCount > 0 {
		failed = -1
	}

	err = endpoint.db.IncrementProgress(ctx, nodeID, transfer.(pendingTransfer).pieceSize, 1, failed)
	if err != nil {
		return Error.Wrap(err)
	}

	err = endpoint.db.DeleteTransferQueueItem(ctx, nodeID, transfer.(pendingTransfer).path)
	if err != nil {
		return Error.Wrap(err)
	}

	pending.Delete(message.Succeeded.GetAddressedOrderLimit().GetLimit().PieceId)

	return nil
}

func (endpoint *Endpoint) handleFailed(ctx context.Context, pending *sync.Map, nodeID storj.NodeID, message *pb.StorageNodeMessage_Failed) (err error) {
	defer mon.Task()(&ctx)(&err)
	endpoint.log.Warn("transfer failed.", zap.String("piece ID", message.Failed.PieceId.String()), zap.String("transfer error", message.Failed.GetError().String()))
	pieceID := message.Failed.PieceId
	transfer, ok := pending.Load(pieceID)
	if !ok {
		endpoint.log.Debug("could not find transfer message in pending queue. skipping .", zap.String("piece ID", pieceID.String()))
	}
	transferQueueItem, err := endpoint.db.GetTransferQueueItem(ctx, nodeID, transfer.(pendingTransfer).path)
	if err != nil {
		return Error.Wrap(err)
	}
	now := time.Now().UTC()
	failedCount := 1
	if transferQueueItem.FailedCount != nil {
		failedCount = *transferQueueItem.FailedCount + 1
	}

	errorCode := int(pb.TransferFailed_Error_value[message.Failed.Error.String()])

	// TODO if error code is NOT_FOUND, the node no longer has the piece. remove the queue item and the pointer

	transferQueueItem.LastFailedAt = &now
	transferQueueItem.FailedCount = &failedCount
	transferQueueItem.LastFailedCode = &errorCode
	err = endpoint.db.UpdateTransferQueueItem(ctx, *transferQueueItem)
	if err != nil {
		return Error.Wrap(err)
	}

	// only increment failed if it hasn't failed before
	if failedCount == 1 {
		err = endpoint.db.IncrementProgress(ctx, nodeID, 0, 0, 1)
		if err != nil {
			return Error.Wrap(err)
		}
	}

	pending.Delete(pieceID)

	return nil
}
