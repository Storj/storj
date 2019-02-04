// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information

package testplanet

import (
	"encoding/hex"
	"math/rand"
	"strconv"
	"testing"

	"github.com/zeebo/errs"
	"go.uber.org/zap/zaptest"

	"storj.io/storj/satellite"
	"storj.io/storj/satellite/satellitedb"
	"storj.io/storj/satellite/satellitedb/satellitedbtest"

	"storj.io/storj/internal/testcontext"
)

// Run runs testplanet in multiple configurations.
func Run(t *testing.T, config Config, test func(t *testing.T, ctx *testcontext.Context, planet *Planet)) {
	schemaSuffix := randomSchemaSuffix()
	t.Log("schema-suffix ", schemaSuffix)

	for _, satelliteDB := range satellitedbtest.Databases() {
		t.Run(satelliteDB.Name, func(t *testing.T) {
			ctx := testcontext.New(t)
			defer ctx.Cleanup()

			if satelliteDB.URL == "" {
				t.Skipf("Database %s connection string not provided. %s", satelliteDB.Name, satelliteDB.Message)
			}

			planetConfig := config
			planetConfig.Reconfigure.NewBootstrapDB = nil
			planetConfig.Reconfigure.NewSatelliteDB = func(index int) (satellite.DB, error) {
				db, err := satellitedb.New(satelliteDB.URL)
				if err != nil {
					t.Fatal(err)
				}

				schema := t.Name() + "-satellite/" + strconv.Itoa(index) + "-" + schemaSuffix

				err = db.SetSchema(schema)
				if err != nil {
					t.Fatal(err)
				}

				return &satelliteSchema{
					DB:     db,
					schema: schema,
				}, db.CreateTables()
			}
			planetConfig.Reconfigure.NewStorageNodeDB = nil

			planet, err := NewCustom(zaptest.NewLogger(t), planetConfig)
			if err != nil {
				t.Fatal(err)
			}
			defer ctx.Check(planet.Shutdown)

			planet.Start(ctx)
			test(t, ctx, planet)
		})
	}
}

// satelliteSchema closes database and drops the associated schema
type satelliteSchema struct {
	satellite.DB
	schema string
}

func (db *satelliteSchema) Close() error {
	return errs.Combine(
		db.DB.DropSchema(db.schema),
		db.DB.Close(),
	)
}

func randomSchemaSuffix() string {
	var data [8]byte
	_, _ = rand.Read(data[:])
	return hex.EncodeToString(data[:])
}
