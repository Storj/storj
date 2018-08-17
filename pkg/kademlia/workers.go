// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package kademlia

import (
	"container/heap"
	"context"
	"log"
	"math/big"
	"sync"
	"time"

	"storj.io/storj/pkg/dht"

	"github.com/zeebo/errs"
	"storj.io/storj/pkg/node"
	proto "storj.io/storj/protos/overlay"
)

var (
	// WorkerError is the class of errors for the worker struct
	WorkerError = errs.Class("worker error")
)

type worker struct {
	pq             PriorityQueue
	mu             *sync.Mutex
	maxResponse    time.Duration
	cancel         context.CancelFunc
	nodeClient     node.Client
	find           dht.NodeID
	workInProgress int
	k              int
}

func newWorker(ctx context.Context, rt *RoutingTable, nodes []*proto.Node, nc node.Client, target dht.NodeID, k int) *worker {
	t, ok := new(big.Int).SetString(target.String(), 2)
	if !ok {
		// TODO(coyle): return error
		return nil
	}

	pq := func(nodes []*proto.Node) PriorityQueue {
		pq := make(PriorityQueue, len(nodes))
		for i, node := range nodes {
			bnode, ok := new(big.Int).SetString(node.GetId(), 2)
			if !ok {
				// TODO(coyle): return error
			}
			pq[i] = &Item{
				value:    node,
				priority: new(big.Int).Xor(t, bnode),
				index:    i,
			}

		}
		heap.Init(&pq)

		return pq
	}(nodes)

	return &worker{
		pq:             pq,
		mu:             &sync.Mutex{},
		maxResponse:    0 * time.Millisecond,
		nodeClient:     nc,
		find:           target,
		workInProgress: 0,
		k:              k,
	}
}

func (w *worker) work(ctx context.Context, ch chan []*proto.Node) error {
	// grab uncontacted node from working set
	// change status to inprogress
	// ask node for target
	// if node has target cancel ctx and send node
	for {
		if ctx.Err() != nil {
			return nil
		}

		nodes := w.lookup(ctx, w.getWork())
		if nodes == nil {
			continue
		}

		if err := w.update(nodes); err != nil {
			//TODO(coyle): determine best way to handle this error
		}

		return nil
	}

}

func (w *worker) getWork() *proto.Node {
	w.mu.Lock()
	if w.pq.Len() <= 0 && w.workInProgress <= 0 {
		w.mu.Unlock()
		time.AfterFunc(2*w.maxResponse, w.cancel)
		return nil
	}

	defer w.mu.Unlock()
	w.workInProgress++
	return w.pq.Pop().(*Item).value
}

func (w *worker) lookup(ctx context.Context, node *proto.Node) []*proto.Node {
	start := time.Now()
	nodes, err := w.nodeClient.Lookup(ctx, *node, proto.Node{Id: w.find.String()})
	if err != nil {
		// TODO(coyle): I think we might want to do another look up on this node or update something
		// but for now let's just log and ignore.
		log.Printf("Error occured during lookup for %s on %s :: error = %s", w.find.String(), node.GetId(), err.Error())
		return nil
	}

	latency := time.Now().Sub(start)
	if latency > w.maxResponse {
		w.maxResponse = latency
	}

	return nodes
}

func (w *worker) update(nodes []*proto.Node) error {
	if len(nodes) == 0 {
		return WorkerError.New("nodes must not be empty")
	}
	t, ok := new(big.Int).SetString(w.find.String(), 2)
	if !ok {
		return WorkerError.New("Unable to parse target ID")
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	for _, v := range nodes {
		bn, ok := new(big.Int).SetString(w.find.String(), 2)
		if !ok {
			return WorkerError.New("Unable to parse node ID")
		}

		w.pq.Push(&Item{
			value:    v,
			priority: new(big.Int).Xor(t, bn),
		})
	}
	// only keep the k closest nodes
	w.pq = w.pq[:w.k]
	// reinitialize heap
	heap.Init(&w.pq)

	return nil
}
