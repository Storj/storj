// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"database/sql"

	"github.com/zeebo/errs"
	"go.uber.org/zap"

	"storj.io/storj/internal/migrate"
)

var ErrMigrate = errs.Class("migrate")

// CreateTables is a method for creating all tables for database
func (db *DB) CreateTables() error {
	switch db.driver {
	case "postgres":
		migration := db.PostgresMigration()
		return migration.Run(db.log.Named("migrate"), db.db)
	default:
		return migrate.Create("database", db.db)
	}
}

func (db *DB) PostgresMigration() *migrate.Migration {
	return &migrate.Migration{
		Table: "versions",
		Steps: []*migrate.Step{
			{
				// some databases may have already this done, although the version may not match
				Description: "Initial setup",
				Version:     0,
				Action: migrate.SQL{
					`CREATE TABLE IF NOT EXISTS accounting_raws (
						id bigserial NOT NULL,
						node_id bytea NOT NULL,
						interval_end_time timestamp with time zone NOT NULL,
						data_total double precision NOT NULL,
						data_type integer NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id )
					)`,
					`CREATE TABLE IF NOT EXISTS accounting_rollups (
						id bigserial NOT NULL,
						node_id bytea NOT NULL,
						start_time timestamp with time zone NOT NULL,
						put_total bigint NOT NULL,
						get_total bigint NOT NULL,
						get_audit_total bigint NOT NULL,
						get_repair_total bigint NOT NULL,
						put_repair_total bigint NOT NULL,
						at_rest_total double precision NOT NULL,
						PRIMARY KEY ( id )
					)`,
					`CREATE TABLE IF NOT EXISTS accounting_timestamps (
						name text NOT NULL,
						value timestamp with time zone NOT NULL,
						PRIMARY KEY ( name )
					)`,
					`CREATE TABLE IF NOT EXISTS bwagreements (
						serialnum text NOT NULL,
						data bytea NOT NULL,
						storage_node bytea NOT NULL,
						action bigint NOT NULL,
						total bigint NOT NULL,
						created_at timestamp with time zone NOT NULL,
						expires_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( serialnum )
					)`,
					`CREATE TABLE IF NOT EXISTS injuredsegments (
						id bigserial NOT NULL,
						info bytea NOT NULL,
						PRIMARY KEY ( id )
					)`,
					`CREATE TABLE IF NOT EXISTS irreparabledbs (
						segmentpath bytea NOT NULL,
						segmentdetail bytea NOT NULL,
						pieces_lost_count bigint NOT NULL,
						seg_damaged_unix_sec bigint NOT NULL,
						repair_attempt_count bigint NOT NULL,
						PRIMARY KEY ( segmentpath )
					)`,
					`CREATE TABLE IF NOT EXISTS nodes (
						id bytea NOT NULL,
						audit_success_count bigint NOT NULL,
						total_audit_count bigint NOT NULL,
						audit_success_ratio double precision NOT NULL,
						uptime_success_count bigint NOT NULL,
						total_uptime_count bigint NOT NULL,
						uptime_ratio double precision NOT NULL,
						created_at timestamp with time zone NOT NULL,
						updated_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id )
					)`,
					`CREATE TABLE IF NOT EXISTS overlay_cache_nodes (
						node_id bytea NOT NULL,
						node_type integer NOT NULL,
						address text NOT NULL,
						protocol integer NOT NULL,
						operator_email text NOT NULL,
						operator_wallet text NOT NULL,
						free_bandwidth bigint NOT NULL,
						free_disk bigint NOT NULL,
						latency_90 bigint NOT NULL,
						audit_success_ratio double precision NOT NULL,
						audit_uptime_ratio double precision NOT NULL,
						audit_count bigint NOT NULL,
						audit_success_count bigint NOT NULL,
						uptime_count bigint NOT NULL,
						uptime_success_count bigint NOT NULL,
						PRIMARY KEY ( node_id ),
						UNIQUE ( node_id )
					)`,
					`CREATE TABLE IF NOT EXISTS projects (
						id bytea NOT NULL,
						name text NOT NULL,
						description text NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id )
					)`,
					`CREATE TABLE IF NOT EXISTS users (
						id bytea NOT NULL,
						first_name text NOT NULL,
						last_name text NOT NULL,
						email text NOT NULL,
						password_hash bytea NOT NULL,
						status integer NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id )
					)`,
					`CREATE TABLE IF NOT EXISTS api_keys (
						id bytea NOT NULL,
						project_id bytea NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
						key bytea NOT NULL,
						name text NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id ),
						UNIQUE ( key ),
						UNIQUE ( name, project_id )
					)`,
					`CREATE TABLE IF NOT EXISTS project_members (
						member_id bytea NOT NULL REFERENCES users( id ) ON DELETE CASCADE,
						project_id bytea NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( member_id, project_id )
					)`,
				},
			},
			{
				// some databases may have already this done, although the version may not match
				Description: "Adjust table naming",
				Version:     1,
				Action: migrate.Func(func(log *zap.Logger, db migrate.DB, tx *sql.Tx) error {
					has_storage_node_id, err := postgresHasColumn(tx, "bwagreements", "storage_node_id")
					if err != nil {
						return ErrMigrate.Wrap(err)
					}
					if !has_storage_node_id {
						// - storage_node bytea NOT NULL,
						// + storage_node_id bytea NOT NULL,
						_, err := tx.Exec(`ALTER TABLE bwagreements
							RENAME COLUMN storage_node TO storage_node_id;`)
						if err != nil {
							return ErrMigrate.Wrap(err)
						}
					}

					has_uplink_id, err := postgresHasColumn(tx, "bwagreements", "uplink_id")
					if err != nil {
						return ErrMigrate.Wrap(err)
					}
					if !has_uplink_id {
						// + uplink_id bytea NOT NULL,
						_, err := tx.Exec(`
							ALTER TABLE bwagreements
								ADD COLUMN uplink_id BYTEA NOT NULL;
						`)
						if err != nil {
							return ErrMigrate.Wrap(err)
						}

						// TODO:
						// walk all data rows
						// unmarshal using the specific protobuf version
						//    rba *pb.RenterBandwidthAllocation
						//    dbx.Bwagreement_Data(rbaBytes),
						//    rbaBytes, err := proto.Marshal(rba)
						//    dbx.Bwagreement_UplinkId(rba.PayerAllocation.UplinkId.Bytes()),

						_, err = tx.Exec(`
							ALTER TABLE bwagreements
								DROP COLUMN data;
						`)
						if err != nil {
							return ErrMigrate.Wrap(err)
						}
					}
					return nil
				}),
			},
			{
				// some databases may have already this done, although the version may not match
				Description: "Remove bucket infos",
				Version:     2,
				Action: migrate.SQL{
					`DROP TABLE IF EXISTS bucket_infos CASCADE`,
				},
			},
			{
				// some databases may have already this done, although the version may not match
				Description: "Add certificates table",
				Version:     3,
				Action: migrate.SQL{
					`CREATE TABLE IF NOT EXISTS certRecords (
						publickey bytea NOT NULL,
						id bytea NOT NULL,
						update_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id )
					)`,
				},
			},
			{
				// some databases may have already this done, although the version may not match
				Description: "Adjust users table",
				Version:     4,
				Action: migrate.Func(func(log *zap.Logger, db migrate.DB, tx *sql.Tx) error {
					// - email text,
					// + email text NOT NULL,
					email_nullable, err := postgresColumnNullability(tx, "users", "email")
					if err != nil {
						return ErrMigrate.Wrap(err)
					}
					if email_nullable {
						_, err := tx.Exec(`
							ALTER TABLE users
							ALTER COLUMN email SET NOT NULL;
						`)
						if err != nil {
							return ErrMigrate.Wrap(err)
						}
					}

					// + status integer NOT NULL,
					has_status, err := postgresHasColumn(tx, "users", "status")
					if err != nil {
						return ErrMigrate.Wrap(err)
					}
					if !has_status {
						_, err := tx.Exec(`
							ALTER TABLE users
								ADD COLUMN status INTEGER NOT NULL;
						`)
						// TODO: what should be the default value?
						if err != nil {
							return ErrMigrate.Wrap(err)
						}
					}

					// - UNIQUE ( email )
					_, err = tx.Exec(`
						ALTER TABLE users
						DROP CONSTRAINT IF EXISTS users_email_key;
					`)
					if err != nil {
						return ErrMigrate.Wrap(err)
					}

					return nil
				}),
			},
		},
	}
}

func postgresHasColumn(tx *sql.Tx, table, column string) (bool, error) {
	var columnName string
	err := tx.QueryRow(`
		SELECT column_name FROM information_schema.COLUMNS 
			WHERE table_schema = CURRENT_SCHEMA
				AND table_name = $1
				AND column_name = $2
		`, table, column).Scan(&columnName)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, ErrMigrate.Wrap(err)
	}

	return columnName == column, nil
}

func postgresColumnNullability(tx *sql.Tx, table, column string) (bool, error) {
	var nullability string
	err := tx.QueryRow(`
		SELECT is_nullable FROM information_schema.COLUMNS 
			WHERE table_schema = CURRENT_SCHEMA
				AND table_name = $1
				AND column_name = $2
		`, table, column).Scan(&nullability)
	if err != nil {
		return false, ErrMigrate.Wrap(err)
	}
	return nullability == "YES", nil
}
