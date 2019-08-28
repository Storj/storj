// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package audit

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"storj.io/storj/satellite/metainfo"
)

// Config2 contains configurable values for audit 2.0 service
type Config2 struct {
	QueueInterval time.Duration `help:"how often to repopulate the audit queue" default:"30s"`
	Slots         int           `help:"number of reservoir slots allotted for nodes, currently capped at 2" default:"1"`
	WorkerCount   int           `help:"number of workers to run audits on paths" default:"1"`
}

// Service2 helps coordinate Cursor and Verifier to run the audit process continuously
type Service2 struct {
	log *zap.Logger

	reservoirChore *ReservoirChore
	workers        []*worker
	queue          *queue
}

// NewService instantiates a Service with access to a Cursor and Verifier
func NewService2(log *zap.Logger, config Config2, metaloop *metainfo.Loop) (*Service2, error) {
	queue := newQueue(*sync.NewCond(&sync.Mutex{}), make(chan struct{}))
	var workers []*worker
	for i := 0; i < config.WorkerCount; i++ {
		workers = append(workers, newWorker(queue))
	}
	return &Service2{
		log: log,

		reservoirChore: NewReservoirChore(log.Named("reservoir chore"), queue, metaloop, config),
		workers:        workers,
		queue:          queue,
	}, nil
}

// Run runs auditing service
func (service *Service2) Run(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)
	service.log.Info("audit 2.0 is starting up")

	var group errgroup.Group
	group.Go(func() error {
		return service.reservoirChore.populateQueueJob(ctx)
	})

	for _, worker := range service.workers {
		group.Go(func() error {
			return worker.run(ctx)
		})
	}

	return group.Wait()

}

// Close halts the reservoir chore and audit workers.
func (service *Service2) Close() error {
	close(service.queue.closed)
	return nil
}

// worker processes items on the audit queue.
type worker struct {
	queue *queue
}

// newWorker instantiates a worker.
func newWorker(queue *queue) *worker {
	return &worker{
		queue: queue,
	}
}

// worker removes an item from the queue and runs an audit.
func (w *worker) run(ctx context.Context) error {
	for {
		_, err := w.queue.next(ctx)
		if err != nil {
			return err
		}
		// TODO: audit the path
	}
}
