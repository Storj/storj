// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package planneddowntime

import (
	"context"
	"time"

	"go.uber.org/zap"
	"storj.io/common/pb"
	"storj.io/common/storj"
)

type DB interface {
	Add(id storj.NodeID, req *pb.ScheduleDowntimeRequest) (*pb.DowntimeWindow, error)
	Delete(id storj.NodeID, req *pb.CancelRequest) error
}

type Service struct {
	log *zap.Logger
	db  *DB
}

func NewService(log *zap.Logger, db *DB) *Service {
	return &Service{
		log: log,
		db:  db,
	}
}

// ScheduleDowntime inserts a downtime into the DB.
func (service *Service) ScheduleDowntime(ctx context.Context, id storj.NodeID, req *pb.ScheduleDowntimeRequest) (_ *pb.ScheduleDowntimeResponse, err error) {
	defer mon.Task()(&ctx)(&err)

	// downtime, err := service.db.Add(ctx, peer.ID, req)
	// if err != nil {
	// 	return nil, err
	// }

	return &pb.ScheduleDowntimeResponse{
		Window: &pb.DowntimeWindow{
			Id: []byte{'a', 'b', 'c'},
			Timeframe: &pb.Timeframe{
				Start: time.Time{},
				End:   time.Now(),
			},
		},
	}, nil
}

// Cancel deletes a scheduled timeframe from the DB.
func (service *Service) Cancel(ctx context.Context, id storj.NodeID, req *pb.CancelRequest) (_ *pb.CancelResponse, err error) {
	defer mon.Task()(&ctx)(&err)

	// err := service.db.Delete(ctx, peer.ID, req.Id)
	// if err != nil {
	// 	return nil, err
	// }
	return &pb.CancelResponse{}, nil
}

// Close closes resources.
func (service *Service) Close() error { return nil }
