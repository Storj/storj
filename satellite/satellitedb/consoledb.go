// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"

	"github.com/zeebo/errs"

	"storj.io/storj/satellite/console"
	dbx "storj.io/storj/satellite/satellitedb/dbx"
)

// ConsoleDB contains access to different satellite databases
type ConsoleDB struct {
	db *dbx.DB
	tx *dbx.Tx

	methods dbx.Methods
}

// Users is getter a for Users repository
func (db *ConsoleDB) Users() console.Users {
	return &users{db.methods}
}

// Projects is a getter for Projects repository
func (db *ConsoleDB) Projects() console.Projects {
	return &projects{db.methods}
}

// ProjectMembers is a getter for ProjectMembers repository
func (db *ConsoleDB) ProjectMembers() console.ProjectMembers {
	return &projectMembers{db.methods, db.db}
}

// APIKeys is a getter for APIKeys repository
func (db *ConsoleDB) APIKeys() console.APIKeys {
	return &apikeys{db.methods}
}

// BucketUsages is a getter for APIKeys repository
func (db *ConsoleDB) BucketUsages() console.BucketUsages {
	return &bucketusages{db.methods}
}

// BeginTx is a method for opening transaction
func (db *ConsoleDB) BeginTx(ctx context.Context) (console.DBTx, error) {
	if db.db == nil {
		return nil, errs.New("DB is not initialized!")
	}

	tx, err := db.db.Open(ctx)
	if err != nil {
		return nil, err
	}

	return &DBTx{
		ConsoleDB: &ConsoleDB{
			tx:      tx,
			methods: tx,
		},
	}, nil
}

// DBTx extends Database with transaction scope
type DBTx struct {
	*ConsoleDB
}

// Commit is a method for committing and closing transaction
func (db *DBTx) Commit() error {
	if db.tx == nil {
		return errs.New("begin transaction before commit it!")
	}

	return db.tx.Commit()
}

// Rollback is a method for rollback and closing transaction
func (db *DBTx) Rollback() error {
	if db.tx == nil {
		return errs.New("begin transaction before rollback it!")
	}

	return db.tx.Rollback()
}
