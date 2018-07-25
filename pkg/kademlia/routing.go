// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package kademlia

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	// bkad "github.com/coyle/kademlia"

	"storj.io/storj/pkg/dht"
	proto "storj.io/storj/protos/overlay"
)
//Move RoutingTable struct here once merged in

//RoutingOptions for configuring RoutingTable
type RoutingOptions struct { 
	kpath 		string
	npath 		string 
	idLength 	int 
	bucketSize 	int
}

// Local returns the local nodes ID
func (rt RoutingTable) Local() proto.Node {
	return proto.Node{
		Id: string(rt.dht.HT.Self.ID),
		Address: &proto.NodeAddress{
			Transport: defaultTransport, // TODO(coyle): this should be stored on the route table
			Address:   fmt.Sprintf("%s:%d", rt.dht.HT.Self.IP.String(), rt.dht.HT.Self.Port),
		},
	}

}

// K returns the currently configured maximum of nodes to store in a bucket
func (rt RoutingTable) K() int {
	return rt.dht.NumNodes()
}

// CacheSize returns the total current size of the cache
func (rt RoutingTable) CacheSize() int {
	// TODO: How is this calculated ? size of the routing table ? is it total bytes, mb, kb etc .?
	return 0
}

// GetBucket retrieves a bucket from the local node
func (rt RoutingTable) GetBucket(id string) (bucket dht.Bucket, ok bool) {
	i, err := hex.DecodeString(id)
	if err != nil {
		return &KBucket{}, false
	}
	b := rt.ht.GetBucket(i)
	if b == nil {
		return &KBucket{}, false
	}

	return &KBucket{
		nodes: convertNetworkNodes(b),
	}, true
}

// GetBuckets retrieves all buckets from the local node
func (rt RoutingTable) GetBuckets() (k []dht.Bucket, err error) {
	bs := []dht.Bucket{}
	b := rt.ht.GetBuckets()

	for _, v := range b {
		bs = append(bs, &KBucket{nodes: convertNetworkNodes(v)})
	}

	return bs, nil
}

// FindNear finds all Nodes near the provided nodeID up to the provided limit
func (rt RoutingTable) FindNear(id dht.NodeID, limit int) ([]*proto.Node, error) {
	return convertNetworkNodes(rt.ht.GetClosestContacts(id.Bytes(), limit)), nil
}

// ConnectionSuccess handles the details of what kademlia should do when
// a successful connection is made to node on the network
func (rt RoutingTable) ConnectionSuccess(id string, address proto.NodeAddress) {
	// TODO: What should we do ?
	return
}

// ConnectionFailed handles the details of what kademlia should do when
// a connection fails for a node on the network
func (rt RoutingTable) ConnectionFailed(id string, address proto.NodeAddress) {
	// TODO: What should we do ?
	return
}

// SetBucketTimestamp updates the last updated time for a bucket
func (rt RoutingTable) SetBucketTimestamp(id string, now time.Time) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		return NodeErr.New("unable to convert id to int")
	}

	rt.ht.SetBucketTime(i, now)

	return nil
}

// GetBucketTimestamp retrieves the last updated time for a bucket
func (rt RoutingTable) GetBucketTimestamp(id string, bucket dht.Bucket) (time.Time, error) {
	return rt.dht.GetExpirationTime([]byte(id)), nil
}

// GetNodeRoutingTable gets a routing table for a given node rather than the local node's routing table
func GetNodeRoutingTable(ctx context.Context, ID NodeID) (RoutingTable, error) {
	return RouteTable{}, nil
}
