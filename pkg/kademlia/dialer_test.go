// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package kademlia_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeebo/errs"
	"go.uber.org/zap/zaptest"
	"golang.org/x/sync/errgroup"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/pkg/kademlia"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
)

func TestDialer(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 4, UplinkCount: 3,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		time.Sleep(2 * time.Second)
		expectedKademliaEntries := 1 + len(planet.Satellites) + len(planet.StorageNodes)

		// TODO: also use satellites
		peers := planet.StorageNodes

		{ // Ping
			self := planet.StorageNodes[0]

			dialer := kademlia.NewDialer(zaptest.NewLogger(t), self.Transport)
			defer ctx.Check(dialer.Close)

			var group errgroup.Group

			for i := range peers {
				peer := peers[i]
				group.Go(func() error {
					pinged, err := dialer.Ping(ctx, peer.Local())
					var pingErr error
					if !pinged {
						pingErr = fmt.Errorf("ping to %s should have succeeded", peer.ID())
					}
					return errs.Combine(pingErr, err)
				})
			}
			defer ctx.Check(group.Wait)
		}

		{ // Lookup
			self := planet.StorageNodes[1]
			dialer := kademlia.NewDialer(zaptest.NewLogger(t), self.Transport)
			defer ctx.Check(dialer.Close)

			var group errgroup.Group

			for i := range peers {
				peer := peers[i]
				group.Go(func() error {
					for _, target := range peers {
						errTag := fmt.Errorf("lookup peer:%s target:%s", peer.ID(), target.ID())
						peer.Local().Type.DPanicOnInvalid("test client peer")
						target.Local().Type.DPanicOnInvalid("test client target")

						results, err := dialer.Lookup(ctx, self.Local(), peer.Local(), target.Local())
						if err != nil {
							return errs.Combine(errTag, err)
						}

						if containsResult(results, target.ID()) {
							continue
						}

						// with small network we expect to return everything
						if len(results) != expectedKademliaEntries {
							return errs.Combine(errTag, fmt.Errorf("expected %d got %d: %s", expectedKademliaEntries, len(results), pb.NodesToIDs(results)))
						}

						return nil
					}
					return nil
				})
			}

			defer ctx.Check(group.Wait)
		}

		{ // Lookup
			self := planet.StorageNodes[2]
			dialer := kademlia.NewDialer(zaptest.NewLogger(t), self.Transport)
			defer ctx.Check(dialer.Close)

			targets := []storj.NodeID{
				{},    // empty target
				{255}, // non-empty
			}

			var group errgroup.Group

			for i := range targets {
				target := targets[i]
				for i := range peers {
					peer := peers[i]
					group.Go(func() error {
						errTag := fmt.Errorf("invalid lookup peer:%s target:%s", peer.ID(), target)
						peer.Local().Type.DPanicOnInvalid("peer info")
						results, err := dialer.Lookup(ctx, self.Local(), peer.Local(), pb.Node{Id: target, Type: pb.NodeType_STORAGE})
						if err != nil {
							return errs.Combine(errTag, err)
						}

						// with small network we expect to return everything
						if len(results) != expectedKademliaEntries {
							return errs.Combine(errTag, fmt.Errorf("expected %d got %d: %s", expectedKademliaEntries, len(results), pb.NodesToIDs(results)))
						}

						return nil
					})
				}
			}

			defer ctx.Check(group.Wait)
		}
	})
}

func containsResult(nodes []*pb.Node, target storj.NodeID) bool {
	for _, node := range nodes {
		if node.Id == target {
			return true
		}
	}
	return false
}
