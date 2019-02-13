// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"storj.io/storj/internal/migrate"
)

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
				Description: "Initial setup",
				Version:     0,
				Action: migrate.SQL{
					`CREATE TABLE accounting_raws (
						id bigserial NOT NULL,
						node_id bytea NOT NULL,
						interval_end_time timestamp with time zone NOT NULL,
						data_total double precision NOT NULL,
						data_type integer NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id )
					)`,
					`CREATE TABLE accounting_rollups (
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
					`CREATE TABLE accounting_timestamps (
						name text NOT NULL,
						value timestamp with time zone NOT NULL,
						PRIMARY KEY ( name )
					)`,
					`CREATE TABLE bwagreements (
						serialnum text NOT NULL,
						storage_node_id bytea NOT NULL,
						uplink_id bytea NOT NULL,
						action bigint NOT NULL,
						total bigint NOT NULL,
						created_at timestamp with time zone NOT NULL,
						expires_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( serialnum )
					)`,
					`CREATE TABLE certRecords (
						publickey bytea NOT NULL,
						id bytea NOT NULL,
						update_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id )
					)`,
					`CREATE TABLE injuredsegments (
						id bigserial NOT NULL,
						info bytea NOT NULL,
						PRIMARY KEY ( id )
					)`,
					`CREATE TABLE irreparabledbs (
						segmentpath bytea NOT NULL,
						segmentdetail bytea NOT NULL,
						pieces_lost_count bigint NOT NULL,
						seg_damaged_unix_sec bigint NOT NULL,
						repair_attempt_count bigint NOT NULL,
						PRIMARY KEY ( segmentpath )
					)`,
					`CREATE TABLE nodes (
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
					`CREATE TABLE overlay_cache_nodes (
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
					`CREATE TABLE projects (
						id bytea NOT NULL,
						name text NOT NULL,
						description text NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id )
					)`,
					`CREATE TABLE users (
						id bytea NOT NULL,
						first_name text NOT NULL,
						last_name text NOT NULL,
						email text NOT NULL,
						password_hash bytea NOT NULL,
						status integer NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id )
					)`,
					`CREATE TABLE api_keys (
						id bytea NOT NULL,
						project_id bytea NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
						key bytea NOT NULL,
						name text NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id ),
						UNIQUE ( key ),
						UNIQUE ( name, project_id )
					)`,
					`CREATE TABLE project_members (
						member_id bytea NOT NULL REFERENCES users( id ) ON DELETE CASCADE,
						project_id bytea NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( member_id, project_id )
					)`,
				},
			},
		},
	}
}
