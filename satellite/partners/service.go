// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

// Package partners implements partners management for attributions.
package partners

import (
	"context"

	"github.com/zeebo/errs"
	"go.uber.org/zap"
)

var (
	// Error is the default error class for partners package.
	Error = errs.Class("partners error class")

	// ErrNotExist is returned when a particular partner does not exist.
	ErrNotExist = errs.Class("partner does not exist")
)

// DB allows access to partners database.
//
// architecture: Database
type DB interface {
	// All returns all partners.
	All(ctx context.Context) ([]Partner, error)
	// ByName returns partner definitions for a given name.
	ByName(ctx context.Context, name string) (Partner, error)
	// ByID returns partner definition corresponding to an id.
	ByID(ctx context.Context, id string) (Partner, error)
	// ByUserAgent returns partner definition corresponding to an user agent string.
	ByUserAgent(ctx context.Context, agent string) (Partner, error)
}

// Service allows manipulating and accessing partner information.
//
// architecture: Service
type Service struct {
	log *zap.Logger
	db  DB
}

// NewService returns a service for handling partner information.
func NewService(log *zap.Logger, db DB) *Service {
	return &Service{
		log: log,
		db:  db,
	}
}
