// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"github.com/zeebo/errs"

	"storj.io/storj/private/dbutil/pgutil"
	"storj.io/storj/private/migrate"
)

var (
	// ErrMigrate is for tracking migration errors.
	ErrMigrate = errs.Class("migrate")
	// ErrMigrateMinVersion is for migration min version errors
	ErrMigrateMinVersion = errs.Class("migrate min version")
)

// MigrateToLatest migrates the database to the latest version.
func (db *satelliteDB) MigrateToLatest(ctx context.Context) error {
	// First handle the idiosyncrasies of postgres and cockroach migrations. Postgres
	// will need to create any schemas specified in the search path, and cockroach
	// will need to create the database it was told to connect to. These things should
	// not really be here, and instead should be assumed to exist.
	// This is tracked in jira ticket SM-200
	switch db.implementation {
	case dbutil.Postgres:
		schema, err := pgutil.ParseSchemaFromConnstr(db.source)
		if err != nil {
			return errs.New("error parsing schema: %+v", err)
		}

		if schema != "" {
			err = pgutil.CreateSchema(ctx, db, schema)
			if err != nil {
				return errs.New("error creating schema: %+v", err)
			}
		}

	case dbutil.Cockroach:
		var dbName string
		if err := db.QueryRow(ctx, `SELECT current_database();`).Scan(&dbName); err != nil {
			return errs.New("error querying current database: %+v", err)
		}

		_, err := db.Exec(ctx, fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %s;`,
			pgutil.QuoteIdentifier(dbName)))
		if err != nil {
			return errs.Wrap(err)
		}
	}

	switch db.implementation {
	case dbutil.Postgres, dbutil.Cockroach:
		migration := db.PostgresMigration()
		// since we merged migration steps 0-64, the current db version should never be
		// less than 65 unless the migration hasn't run yet
		const minDBVersion = 65
		dbVersion, err := migration.CurrentVersion(db.log, db.db)
		if err != nil {
			return errs.New("error current version: %+v", err)
		}
		if dbVersion > -1 && dbVersion < minDBVersion {
			return ErrMigrateMinVersion.New("current database version is %d, it shouldn't be less than the min version %d",
				dbVersion, minDBVersion,
			)
		}

		return migration.Run(db.log.Named("migrate"))
	default:
		return migrate.Create(ctx, "database", db.DB)
	}
}

// CheckVersion confirms the database is at the desired version
func (db *DB) CheckVersion() error {
	switch db.driver {
	case "postgres":
		migration := db.PostgresMigration()

		dbVersion, err := migration.CurrentVersion(ctx, db.log, db.DB)
		if err != nil {
			return ErrMigrateMinVersion.Wrap(err)
		}
		if dbVersion > -1 {
			return ErrMigrateMinVersion.New("the database must be empty, got version %d", dbVersion)
		}

		flattened, err := flattenMigration(migration)
		if err != nil {
			return ErrMigrateMinVersion.Wrap(err)
		}

		return flattened.Run(ctx, db.log.Named("migrate"))
	default:
		return migrate.Create(ctx, "database", db.DB)
	}
}

// CheckVersion confirms the database is at the desired version.
func (db *satelliteDB) CheckVersion(ctx context.Context) error {
	switch db.implementation {
	case dbutil.Postgres, dbutil.Cockroach:
		migration := db.PostgresMigration()
		return migration.ValidateVersions(ctx, db.log)

	default:
		return nil
	}
}

// flattenMigration joins the migration sql queries from
// each migration step to speed up the database setup.
//
// Steps with "SeparateTx" end up as separate migration transactions.
// Cockroach requires schema changes and updates to the values of that
// schema change to be in a different transaction.
func flattenMigration(m *migrate.Migration) (*migrate.Migration, error) {
	var db tagsql.DB
	var version int
	var statements migrate.SQL
	var steps []*migrate.Step

	pushMerged := func() {
		if len(statements) == 0 {
			return
		}

		steps = append(steps, &migrate.Step{
			DB:          &db,
			Description: "Setup",
			Version:     version,
			Action:      migrate.SQL{strings.Join(statements, ";\n")},
		})

		statements = nil
	}

	for _, step := range m.Steps {
		if db == nil {
			db = *step.DB
		} else if db != *step.DB {
			return nil, errs.New("multiple databases not supported")
		}

		if sql, ok := step.Action.(migrate.SQL); ok {
			if step.SeparateTx {
				pushMerged()
			}

			version = step.Version
			statements = append(statements, sql...)
		} else {
			pushMerged()
			steps = append(steps, step)
		}
	}

	pushMerged()

	return &migrate.Migration{
		Table: "versions",
		Steps: steps,
	}, nil
}

// PostgresMigration returns steps needed for migrating postgres database.
func (db *satelliteDB) PostgresMigration() *migrate.Migration {
	return &migrate.Migration{
		Table: "versions",
		Steps: []*migrate.Step{

			{
				DB:          db.db,
				Description: "Initial setup",
				Version:     65,
				Action: migrate.SQL{
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
					);`,

					`CREATE TABLE accounting_timestamps (
						name text NOT NULL,
						value timestamp with time zone NOT NULL,
						PRIMARY KEY ( name )
					);`,

					`CREATE TABLE bucket_bandwidth_rollups (
						bucket_name bytea NOT NULL,
						interval_start timestamp NOT NULL,
						interval_seconds integer NOT NULL,
						action integer NOT NULL,
						inline bigint NOT NULL,
						allocated bigint NOT NULL,
						settled bigint NOT NULL,
						project_id bytea NOT NULL ,
						CONSTRAINT bucket_bandwidth_rollups_pk PRIMARY KEY (bucket_name, project_id, interval_start, action)
					);`,
					`CREATE INDEX bucket_name_project_id_interval_start_interval_seconds ON bucket_bandwidth_rollups ( bucket_name, project_id, interval_start, interval_seconds );`,

					`CREATE TABLE bucket_storage_tallies (
						bucket_name bytea NOT NULL,
						interval_start timestamp NOT NULL,
						inline bigint NOT NULL,
						remote bigint NOT NULL,
						remote_segments_count integer NOT NULL,
						inline_segments_count integer NOT NULL,
						object_count integer NOT NULL,
						metadata_size bigint NOT NULL,
						project_id bytea NOT NULL,
						CONSTRAINT bucket_storage_tallies_pk PRIMARY KEY (bucket_name, project_id, interval_start)
					);`,

					`CREATE TABLE injuredsegments (
						data bytea NOT NULL,
						attempted timestamp,
						path bytea NOT NULL,
						CONSTRAINT injuredsegments_pk PRIMARY KEY (path)
					);`,
					`CREATE INDEX injuredsegments_attempted_index ON injuredsegments ( attempted );`,

					`CREATE TABLE irreparabledbs (
						segmentpath bytea NOT NULL,
						segmentdetail bytea NOT NULL,
						pieces_lost_count bigint NOT NULL,
						seg_damaged_unix_sec bigint NOT NULL,
						repair_attempt_count bigint NOT NULL,
						PRIMARY KEY ( segmentpath )
					);`,

					`CREATE TABLE nodes (
						id bytea NOT NULL,
						audit_success_count bigint NOT NULL DEFAULT 0,
						total_audit_count bigint NOT NULL DEFAULT 0,
						uptime_success_count bigint NOT NULL,
						total_uptime_count bigint NOT NULL,
						created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
						updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
						wallet text NOT NULL,
						email text NOT NULL,
						address text NOT NULL DEFAULT '',
						protocol INTEGER NOT NULL DEFAULT 0,
						type INTEGER NOT NULL DEFAULT 0,
						free_bandwidth BIGINT NOT NULL DEFAULT -1,
						free_disk BIGINT NOT NULL DEFAULT -1,
						latency_90 BIGINT NOT NULL DEFAULT 0,
						last_contact_success TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT 'epoch',
						last_contact_failure TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT 'epoch',
						major bigint NOT NULL DEFAULT 0,
						minor bigint NOT NULL DEFAULT 0,
						patch bigint NOT NULL DEFAULT 0,
						hash TEXT NOT NULL DEFAULT '',
						timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT '0001-01-01 00:00:00+00',
						release bool NOT NULL DEFAULT FALSE,
						contained bool NOT NULL DEFAULT FALSE,
						last_net text NOT NULL,
						disqualified timestamp with time zone,
						audit_reputation_alpha double precision NOT NULL DEFAULT 1,
						audit_reputation_beta double precision NOT NULL DEFAULT 0,
						uptime_reputation_alpha double precision NOT NULL DEFAULT 1,
						uptime_reputation_beta double precision NOT NULL DEFAULT 0,
						piece_count bigint NOT NULL DEFAULT 0,
						exit_loop_completed_at TIMESTAMP,
						exit_initiated_at TIMESTAMP,
						exit_finished_at TIMESTAMP,
						exit_success boolean NOT NULL DEFAULT FALSE,
						PRIMARY KEY ( id )
					);`,
					`CREATE INDEX node_last_ip ON nodes ( last_net );`,

					`CREATE TABLE offers (
						id serial NOT NULL,
						name text NOT NULL,
						description text NOT NULL,
						type integer NOT NULL,
						award_credit_duration_days integer,
						invitee_credit_duration_days integer,
						redeemable_cap integer,
						expires_at timestamp with time zone NOT NULL,
						created_at timestamp with time zone NOT NULL,
						status integer NOT NULL,
						award_credit_in_cents integer NOT NULL DEFAULT 0,
						invitee_credit_in_cents integer NOT NULL DEFAULT 0,
						PRIMARY KEY ( id )
					);`,

					`CREATE TABLE peer_identities (
						node_id bytea NOT NULL,
						leaf_serial_number bytea NOT NULL,
						chain bytea NOT NULL,
						updated_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( node_id )
					);`,

					`CREATE TABLE pending_audits (
						node_id bytea NOT NULL,
						piece_id bytea NOT NULL,
						stripe_index bigint NOT NULL,
						share_size bigint NOT NULL,
						expected_share_hash bytea NOT NULL,
						reverify_count bigint NOT NULL,
						path bytea NOT NULL,
						PRIMARY KEY ( node_id )
					);`,

					`CREATE TABLE projects (
						id bytea NOT NULL,
						name text NOT NULL,
						description text NOT NULL,
						created_at timestamp with time zone NOT NULL,
						usage_limit bigint NOT NULL DEFAULT 0,
						partner_id bytea,
						owner_id bytea NOT NULL,
						PRIMARY KEY ( id )
					);`,

					`CREATE TABLE registration_tokens (
						secret bytea NOT NULL,
						owner_id bytea UNIQUE,
						project_limit integer NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( secret )
					);`,

					`CREATE TABLE reset_password_tokens (
						secret bytea NOT NULL,
						owner_id bytea NOT NULL UNIQUE,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( secret )
					);`,

					`CREATE TABLE serial_numbers (
						id serial NOT NULL,
						serial_number bytea NOT NULL,
						bucket_id bytea NOT NULL,
						expires_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id )
					);`,
					`CREATE INDEX serial_numbers_expires_at_index ON serial_numbers ( expires_at );`,
					`CREATE UNIQUE INDEX serial_number_index ON serial_numbers ( serial_number )`,

					`CREATE TABLE storagenode_bandwidth_rollups (
						storagenode_id bytea NOT NULL,
						interval_start timestamp with time zone NOT NULL,
						interval_seconds integer NOT NULL,
						action integer NOT NULL,
						allocated bigint DEFAULT 0,
						settled bigint NOT NULL,
						PRIMARY KEY ( storagenode_id, interval_start, action )
					);`,
					`CREATE INDEX storagenode_id_interval_start_interval_seconds_index ON storagenode_bandwidth_rollups ( storagenode_id, interval_start, interval_seconds );`,

					`CREATE TABLE accounting_raws (
						id bigserial NOT NULL,
						node_id bytea NOT NULL,
						interval_end_time timestamp with time zone NOT NULL,
						data_total double precision NOT NULL,
						PRIMARY KEY ( id )
					)`,
					`ALTER TABLE accounting_raws RENAME TO storagenode_storage_tallies`,

					`CREATE TABLE users (
						id bytea NOT NULL,
						full_name text NOT NULL,
						short_name text,
						email text NOT NULL,
						password_hash bytea NOT NULL,
						status integer NOT NULL,
						created_at timestamp with time zone NOT NULL,
						partner_id bytea,
						normalized_email text NOT NULL,
						PRIMARY KEY ( id )
					);`,

					`CREATE TABLE value_attributions (
						bucket_name bytea NOT NULL,
						partner_id bytea NOT NULL,
						last_updated timestamp NOT NULL,
						project_id bytea NOT NULL,
						PRIMARY KEY (project_id, bucket_name)
					);`,

					`CREATE TABLE api_keys (
						id bytea NOT NULL,
						project_id bytea NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
						head bytea NOT NULL UNIQUE,
						name text NOT NULL,
						secret bytea NOT NULL,
						created_at timestamp with time zone NOT NULL,
						partner_id bytea,
						PRIMARY KEY ( id ),
						UNIQUE ( name, project_id )
					);`,

					`CREATE TABLE bucket_metainfos (
						id bytea NOT NULL,
						project_id bytea NOT NULL REFERENCES projects( id ),
						name bytea NOT NULL,
						path_cipher integer NOT NULL,
						created_at timestamp with time zone NOT NULL,
						default_segment_size integer NOT NULL,
						default_encryption_cipher_suite integer NOT NULL,
						default_encryption_block_size integer NOT NULL,
						default_redundancy_algorithm integer NOT NULL,
						default_redundancy_share_size integer NOT NULL,
						default_redundancy_required_shares integer NOT NULL,
						default_redundancy_repair_shares integer NOT NULL,
						default_redundancy_optimal_shares integer NOT NULL,
						default_redundancy_total_shares integer NOT NULL,
						partner_id bytea,
						PRIMARY KEY ( id ),
						UNIQUE ( name, project_id )
					);`,

					`CREATE TABLE project_invoice_stamps (
						project_id bytea NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
						invoice_id bytea NOT NULL UNIQUE,
						start_date timestamp with time zone NOT NULL,
						end_date timestamp with time zone NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( project_id, start_date, end_date )
					);`,

					`CREATE TABLE project_members (
						member_id bytea NOT NULL REFERENCES users( id ) ON DELETE CASCADE,
						project_id bytea NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( member_id, project_id )
					);`,

					`CREATE TABLE used_serials (
						serial_number_id integer NOT NULL REFERENCES serial_numbers( id ) ON DELETE CASCADE,
						storage_node_id bytea NOT NULL,
						PRIMARY KEY ( serial_number_id, storage_node_id )
					);`,

					`CREATE TABLE user_credits (
						id serial NOT NULL,
						user_id bytea NOT NULL REFERENCES users( id ) ON DELETE CASCADE,
						offer_id integer NOT NULL REFERENCES offers( id ),
						referred_by bytea REFERENCES users( id ) ON DELETE SET NULL,
						credits_earned_in_cents integer NOT NULL,
						credits_used_in_cents integer NOT NULL,
						expires_at timestamp with time zone NOT NULL,
						created_at timestamp with time zone NOT NULL,
						type text NOT NULL,
						PRIMARY KEY ( id )
					);`,
					`CREATE UNIQUE INDEX credits_earned_user_id_offer_id ON user_credits (id, offer_id);`,

					`INSERT INTO offers (
						id,
						name,
						description,
						award_credit_in_cents,
						invitee_credit_in_cents,
						expires_at,
						created_at,
						status,
						type,
						award_credit_duration_days,
						invitee_credit_duration_days
					)
					VALUES (
						1,
						'Default referral offer',
						'Is active when no other active referral offer',
						300,
						600,
						'2119-03-14 08:28:24.636949+00',
						'2019-07-14 08:28:24.636949+00',
						1,
						2,
						365,
						14
					),
					(
						2,
						'Default free credit offer',
						'Is active when no active free credit offer',
						0,
						300,
						'2119-03-14 08:28:24.636949+00',
						'2019-07-14 08:28:24.636949+00',
						1,
						1,
						NULL,
						14
					) ON CONFLICT DO NOTHING;`,

					`CREATE TABLE graceful_exit_progress (
						node_id bytea NOT NULL,
						bytes_transferred bigint NOT NULL,
						updated_at timestamp NOT NULL,
						pieces_transferred bigint NOT NULL DEFAULT 0,
						pieces_failed bigint NOT NULL DEFAULT 0,
						PRIMARY KEY ( node_id )
					);`,

					`CREATE TABLE graceful_exit_transfer_queue (
						node_id bytea NOT NULL,
						path bytea NOT NULL,
						piece_num integer NOT NULL,
						durability_ratio double precision NOT NULL,
						queued_at timestamp NOT NULL,
						requested_at timestamp,
						last_failed_at timestamp,
						last_failed_code integer,
						failed_count integer,
						finished_at timestamp,
						PRIMARY KEY ( node_id, path, piece_num )
					);`,

					`CREATE TABLE stripe_customers (
						user_id bytea NOT NULL,
						customer_id text NOT NULL UNIQUE,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( user_id )
					);`,

					`CREATE TABLE coinpayments_transactions (
						id text NOT NULL,
						user_id bytea NOT NULL,
						address text NOT NULL,
						amount bytea NOT NULL,
						received bytea NOT NULL,
						status integer NOT NULL,
						key text NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id )
					);`,

					`CREATE TABLE stripecoinpayments_apply_balance_intents (
						tx_id text NOT NULL REFERENCES coinpayments_transactions( id ) ON DELETE CASCADE,
						state integer NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( tx_id )
					);`,

					`CREATE TABLE stripecoinpayments_invoice_project_records (
						id bytea NOT NULL,
						project_id bytea NOT NULL,
						storage double precision NOT NULL,
						egress bigint NOT NULL,
						objects bigint NOT NULL,
						period_start timestamp with time zone NOT NULL,
						period_end timestamp with time zone NOT NULL,
						state integer NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id ),
						UNIQUE ( project_id, period_start, period_end )
					);`,
				},
			},
			{
				DB:          &db.migrationDB,
				Description: "reset everyone with default rate limits to NULL",
				Version:     129,
				SeparateTx:  true,
				Action: migrate.SQL{
					`UPDATE projects SET max_buckets = NULL WHERE max_buckets <= 100;`,
					`UPDATE projects SET usage_limit = NULL WHERE usage_limit <= 50000000000;`,
					`UPDATE projects SET bandwidth_limit = NULL WHERE bandwidth_limit <= 50000000000;`,
				},
			},
			{
				DB:          &db.migrationDB,
				Description: "add index for the gracefule exit transfer queue",
				Version:     130,
				SeparateTx:  true,
				Action: migrate.SQL{
					`CREATE INDEX IF NOT EXISTS graceful_exit_transfer_queue_nid_dr_qa_fa_lfa_index ON graceful_exit_transfer_queue ( node_id, durability_ratio, queued_at, finished_at, last_failed_at );`,
				},
			},
			{
				DB:          db.db,
				Description: "Add stripecoinpayments_tx_conversion_rates",
				Version:     68,
				Action: migrate.SQL{
					`CREATE TABLE stripecoinpayments_tx_conversion_rates (
						tx_id text NOT NULL,
						rate bytea NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( tx_id )
					);`,
				},
			},
			{
				DB:          db.db,
				Description: "Add timeout field to coinpayments_transaction",
				Version:     69,
				Action: migrate.SQL{
					`DROP TABLE coinpayments_transactions CASCADE;`,
					`DELETE FROM stripecoinpayments_apply_balance_intents`,
					`CREATE TABLE coinpayments_transactions (
						id text NOT NULL,
						user_id bytea NOT NULL,
						address text NOT NULL,
						amount bytea NOT NULL,
						received bytea NOT NULL,
						status integer NOT NULL,
						key text NOT NULL,
						timeout integer NOT NULL,
						created_at timestamp with time zone NOT NULL,
						PRIMARY KEY ( id )
					);`,
					`ALTER TABLE stripecoinpayments_apply_balance_intents
						ADD CONSTRAINT fk_transactions FOREIGN KEY(tx_id) REFERENCES coinpayments_transactions(id) 
						ON DELETE CASCADE;`,
				},
			},
		},
	}
}
