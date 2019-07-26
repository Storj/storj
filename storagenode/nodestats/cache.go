// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package nodestats

import (
	"context"
	"time"

	"storj.io/storj/storagenode/trust"

	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"storj.io/storj/internal/date"
	"storj.io/storj/internal/sync2"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/storagenode/console"
	"storj.io/storj/storagenode/reputation"
)

var (
	// NodeStatsCacheErr defines node stats cache loop error
	NodeStatsCacheErr = errs.Class("node stats cache error")
)

// Cache runs cache loop and stores reputation stats
// and storage usage into db
type Cache struct {
	log *zap.Logger

	service      *Service
	trust        *trust.Pool
	consoleDB    console.DB
	reputationDB reputation.DB

	statsCycle *sync2.Cycle
	spaceCycle *sync2.Cycle
}

// NewCache creates new caching service instance
func NewCache(log *zap.Logger, service *Service, trust *trust.Pool, consoleDB console.DB, reputationDB reputation.DB) *Cache {
	return &Cache{
		log:          log,
		service:      service,
		trust:        trust,
		consoleDB:    consoleDB,
		reputationDB: reputationDB,
	}
}

// Run runs loop
func (cache *Cache) Run(ctx context.Context) error {
	var group errgroup.Group

	cache.statsCycle.Start(ctx, &group, func(ctx context.Context) error {
		err := cache.CacheReputationStats(ctx)
		if err != nil {
			cache.log.Error("Get stats query failed", zap.Error(err))
		}

		return nil
	})
	cache.spaceCycle.Start(ctx, &group, func(ctx context.Context) error {
		err := cache.CacheSpaceUsage(ctx)
		if err != nil {
			cache.log.Error("Get disk space usage query failed", zap.Error(err))
		}

		return nil
	})

	return group.Wait()
}

// CacheReputationStats queries node stats from all the satellites
// known to the storagenode and stores information into db
func (cache *Cache) CacheReputationStats(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	var cacheStatsErr errs.Group
	for _, satellite := range cache.trust.GetSatellites(ctx) {
		stats, err := cache.service.GetReputationStats(ctx, satellite)
		if err != nil {
			cacheStatsErr.Add(NodeStatsCacheErr.Wrap(err))
			continue
		}

		if err = cache.reputationDB.Store(ctx, *stats); err != nil {
			cacheStatsErr.Add(NodeStatsCacheErr.Wrap(err))
			continue
		}
	}

	return cacheStatsErr.Err()
}

// CacheSpaceUsage queries disk space usage from all the satellites
// known to the storagenode and stores information into db
func (cache *Cache) CacheSpaceUsage(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	// get current month edges
	startDate, endDate := date.MonthBoundary(time.Now().UTC())

	var cacheSpaceErr errs.Group
	for _, satellite := range cache.trust.GetSatellites(ctx) {
		spaceUsages, err := cache.service.GetDailyStorageUsage(ctx, satellite, startDate, endDate)
		if err != nil {
			cacheSpaceErr.Add(NodeStatsCacheErr.Wrap(err))
			continue
		}

		err = cache.consoleDB.DiskSpaceUsages().Store(ctx, spaceUsages)
		if err != nil {
			cacheSpaceErr.Add(NodeStatsCacheErr.Wrap(err))
			continue
		}
	}

	return cacheSpaceErr.Err()
}

// fromSpaceUsageResponse get DiskSpaceUsage slice from pb.SpaceUsageResponse
func fromSpaceUsageResponse(resp *pb.DailyStorageUsageResponse, satelliteID storj.NodeID) []console.DiskSpaceUsage {
	var stamps []console.DiskSpaceUsage

	for _, pbUsage := range resp.GetDailyStorageUsage() {
		stamps = append(stamps, console.DiskSpaceUsage{
			SatelliteID: satelliteID,
			AtRestTotal: pbUsage.AtRestTotal,
			Timestamp:   pbUsage.Timestamp,
		})
	}

	return stamps
}

// Close closes underlying cycles
func (cache *Cache) Close() error {
	defer mon.Task()(nil)(nil)
	cache.statsCycle.Close()
	cache.spaceCycle.Close()
	return nil
}
