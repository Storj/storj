// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package accounting_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testrand"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/satellite"
	"storj.io/storj/satellite/accounting"
	"storj.io/storj/satellite/satellitedb/satellitedbtest"
)

func TestSaveBucketTallies(t *testing.T) {
	satellitedbtest.Run(t, func(t *testing.T, db satellite.DB) {
		ctx := testcontext.New(t)
		defer ctx.Cleanup()

		// Setup: create bucket storage tallies
		projectID := testrand.UUID()

		bucketTallies, expectedTallies, err := createBucketStorageTallies(projectID)
		require.NoError(t, err)

		// Execute test:  retrieve the save tallies and confirm they contains the expected data
		intervalStart := time.Now()
		pdb := db.ProjectAccounting()
		actualTallies, err := pdb.SaveTallies(ctx, intervalStart, bucketTallies)
		require.NoError(t, err)
		for _, tally := range actualTallies {
			require.Contains(t, expectedTallies, tally)
		}
	})
}

func TestStorageNodeUsage(t *testing.T) {
	satellitedbtest.Run(t, func(t *testing.T, db satellite.DB) {
		ctx := testcontext.New(t)
		defer ctx.Cleanup()

		const (
			days = 30
		)

		nodeID := testrand.NodeID()
		startDate := time.Now().UTC().Add(time.Hour * 24 * -days)

		var nodes storj.NodeIDList
		nodes = append(nodes, nodeID)
		nodes = append(nodes, testrand.NodeID())
		nodes = append(nodes, testrand.NodeID())
		nodes = append(nodes, testrand.NodeID())

		rollups, tallies, lastDate := makeRollupsAndStorageNodeStorageTallies(nodes, startDate, days)

		lastRollup := rollups[lastDate]
		delete(rollups, lastDate)

		accountingDB := db.StoragenodeAccounting()

		// create last rollup timestamp
		_, err := accountingDB.LastTimestamp(ctx, accounting.LastRollup)
		require.NoError(t, err)

		// save tallies
		for latestTally, tallies := range tallies {
			err = accountingDB.SaveTallies(ctx, latestTally, tallies)
			require.NoError(t, err)
		}

		// save rollup
		err = accountingDB.SaveRollup(ctx, lastDate.Add(time.Hour*-24), rollups)
		require.NoError(t, err)

		nodeStorageUsages, err := accountingDB.QueryStorageNodeUsage(ctx, nodeID, time.Time{}, time.Now().UTC())
		require.NoError(t, err)
		assert.NotNil(t, nodeStorageUsages)
		assert.Equal(t, days, len(nodeStorageUsages))

		// check usage from rollups
		for _, usage := range nodeStorageUsages[:len(nodeStorageUsages)-1] {
			assert.Equal(t, nodeID, usage.NodeID)
			assert.Equal(t, rollups[usage.Timestamp.UTC()][nodeID].AtRestTotal, usage.StorageUsed)
		}

		// check last usage that calculated from tallies
		lastUsage := nodeStorageUsages[len(nodeStorageUsages)-1]

		assert.Equal(t,
			nodeID,
			lastUsage.NodeID)
		assert.Equal(t,
			lastRollup[nodeID].StartTime,
			lastUsage.Timestamp.UTC())
		assert.Equal(t,
			lastRollup[nodeID].AtRestTotal,
			lastUsage.StorageUsed)
	})
}

func createBucketStorageTallies(projectID uuid.UUID) (map[string]*accounting.BucketTally, []accounting.BucketTally, error) {
	bucketTallies := make(map[string]*accounting.BucketTally)
	var expectedTallies []accounting.BucketTally

	for i := 0; i < 4; i++ {
		bucketName := fmt.Sprintf("%s%d", "testbucket", i)
		bucketID := storj.JoinPaths(projectID.String(), bucketName)

		// Setup: The data in this tally should match the pointer that the uplink.upload created
		tally := accounting.BucketTally{
			BucketName:     []byte(bucketName),
			ProjectID:      projectID[:],
			InlineSegments: int64(1),
			RemoteSegments: int64(1),
			Files:          int64(1),
			InlineBytes:    int64(1),
			RemoteBytes:    int64(1),
			MetadataSize:   int64(1),
		}
		bucketTallies[bucketID] = &tally
		expectedTallies = append(expectedTallies, tally)

	}
	return bucketTallies, expectedTallies, nil
}

// make rollups and tallies for specified nodes and date range
func makeRollupsAndStorageNodeStorageTallies(nodes []storj.NodeID, start time.Time, days int) (accounting.RollupStats, map[time.Time]map[storj.NodeID]float64, time.Time) {
	rollups := make(accounting.RollupStats)
	tallies := make(map[time.Time]map[storj.NodeID]float64)

	const (
		hours = 12
	)

	for i := 0; i < days; i++ {
		startDay := time.Date(start.Year(), start.Month(), start.Day()+i, 0, 0, 0, 0, start.Location())
		if rollups[startDay] == nil {
			rollups[startDay] = make(map[storj.NodeID]*accounting.Rollup)
		}

		for _, node := range nodes {
			rollup := &accounting.Rollup{
				NodeID:    node,
				StartTime: startDay,
			}

			for h := 0; h < hours; h++ {
				startTime := startDay.Add(time.Hour * time.Duration(h))
				//fmt.Println(startTime)
				if tallies[startTime] == nil {
					tallies[startTime] = make(map[storj.NodeID]float64)
				}

				tallieAtRest := math.Round(testrand.Float64n(1000))
				tallies[startTime][node] = tallieAtRest
				rollup.AtRestTotal += tallieAtRest
			}

			rollups[startDay][node] = rollup
		}
	}

	return rollups, tallies,
		time.Date(start.Year(), start.Month(), start.Day()+days-1, 0, 0, 0, 0, start.Location())
}
