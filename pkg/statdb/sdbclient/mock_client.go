// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package sdbclient

import (
	"context"

	pb "storj.io/storj/pkg/statdb/proto"
)

// MockStatDB creates a noop Mock Statdb Client
type MockStatDB struct{}

// NewMockClient initializes a new mock statdb client
func NewMockClient() Client {
	return &MockStatDB{}
}

// a compiler trick to make sure *MockStatDB implements Client
var _ Client = (*MockStatDB)(nil)

// Create is used for creating a new entry in the stats db with default reputation
func (sdb *MockStatDB) Create(ctx context.Context, nodeID []byte) (err error) {
	defer mon.Task()(&ctx)(&err)

	return nil
}

// CreateWithStats is used for creating a new entry in the stats db with a specific reputation
func (sdb *MockStatDB) CreateWithStats(ctx context.Context, nodeID []byte, stats *pb.NodeStats) (err error) {
	defer mon.Task()(&ctx)(&err)

	return nil
}

// Get is used for retrieving an entry from statdb or creating a new one if one does not exist
func (sdb *MockStatDB) Get(ctx context.Context, nodeID []byte) (stats *pb.NodeStats, err error) {
	defer mon.Task()(&ctx)(&err)

	stats = &pb.NodeStats{
		AuditSuccessRatio: 0,
		UptimeRatio:       0,
		AuditCount:        0,
	}
	return stats, nil
}

// FindValidNodes is used for retrieving a subset of nodes that meet a minimum reputation requirement
func (sdb *MockStatDB) FindValidNodes(ctx context.Context, nodeIDs [][]byte, minStats *pb.NodeStats) (passedIDs [][]byte, err error) {
	return nil, nil
}

// Update is used for updating a node's stats in the stats db
func (sdb *MockStatDB) Update(ctx context.Context, nodeID []byte,
	auditSuccess, isUp bool, latencyList []int64,
	updateAuditSuccess, updateUptime, updateLatency bool) (stats *pb.NodeStats, err error) {
	stats = &pb.NodeStats{
		AuditSuccessRatio: 0,
		UptimeRatio:       0,
		AuditCount:        0,
	}
	return stats, nil
}

// UpdateUptime is used for updating a node's uptime in statdb
func (sdb *MockStatDB) UpdateUptime(ctx context.Context, nodeID []byte,
	isUp bool) (stats *pb.NodeStats, err error) {
	stats = &pb.NodeStats{
		AuditSuccessRatio: 0,
		UptimeRatio:       0,
		AuditCount:        0,
	}
	return stats, nil
}

// UpdateBatch is used for updating multiple nodes' stats in the stats db
func (sdb *MockStatDB) UpdateBatch(ctx context.Context, nodes []*pb.Node) (statsList []*pb.NodeStats, failedNodes []*pb.Node, err error) {
	return nil, nil, nil
}

// CreateEntryIfNotExists creates a db entry for a node if entry doesn't already exist
func (sdb *MockStatDB) CreateEntryIfNotExists(ctx context.Context, nodeID []byte) (stats *pb.NodeStats, err error) {
	stats = &pb.NodeStats{
		AuditSuccessRatio: 0,
		UptimeRatio:       0,
		AuditCount:        0,
	}
	return stats, nil
}
