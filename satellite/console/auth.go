// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package console

import (
	"context"
	"encoding/base64"

	"github.com/zeebo/errs"

	"storj.io/storj/satellite/console/consoleauth"
)

//TODO: change to JWT or Macaroon based auth

// Signer creates signature for provided data
type Signer interface {
	Sign(data []byte) ([]byte, error)
}

// signToken signs token with given signer
func signToken(token *consoleauth.Token, signer Signer) error {
	encoded := base64.URLEncoding.EncodeToString(token.Payload)

	signature, err := signer.Sign([]byte(encoded))
	if err != nil {
		return err
	}

	token.Signature = signature
	return nil
}

// key is a context value key type
type key int

// connectorID is a context value key type
type connectorID string

// authKey is context key for Authorization
const authKey key = 0

// connectorKey is context key for connector ID
var connectorKey connectorID = "CONNECTORID"

// ErrUnauthorized is error class for authorization related errors
var ErrUnauthorized = errs.Class("unauthorized error")

// Authorization contains auth info of authorized User
type Authorization struct {
	User   User
	Claims consoleauth.Claims
}

// WithAuth creates new context with Authorization
func WithAuth(ctx context.Context, auth Authorization) context.Context {
	return context.WithValue(ctx, authKey, auth)
}

// WithAuthFailure creates new context with authorization failure
func WithAuthFailure(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, authKey, err)
}

// WithConnectorID creates new context with partner connector ID
func WithConnectorID(ctx context.Context, auth Authorization) context.Context {
	return context.WithValue(ctx, connectorKey, auth)
}

// GetAuth gets Authorization from context
func GetAuth(ctx context.Context) (Authorization, error) {
	value := ctx.Value(authKey)

	if auth, ok := value.(Authorization); ok {
		return auth, nil
	}

	if _, ok := value.(error); ok {
		return Authorization{}, errs.New(internalErrMsg)
	}

	return Authorization{}, errs.New(unauthorizedErrMsg)
}

// GetConnectorIDInfo gets partner's connector ID
func GetConnectorIDInfo(ctx context.Context) (Authorization, error) {
	value := ctx.Value(connectorKey)

	if auth, ok := value.(Authorization); ok {
		return auth, nil
	}

	if _, ok := value.(error); ok {
		return Authorization{}, errs.New(internalErrMsg)
	}

	return Authorization{}, errs.New(NoConnectorIDSetErrMsg)
}
