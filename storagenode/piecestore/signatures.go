// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package piecestore

import (
	"context"
	"time"

	"github.com/zeebo/errs"

	"storj.io/storj/pkg/identity"
	"storj.io/storj/pkg/pb"
)

var (
	ErrVerifyBadRequest       = errs.New("bad request")
	ErrVerifyNotAuthorized    = errs.New("not authorized")
	ErrVerifyUntrusted        = errs.New("untrusted")
	ErrVerifyDuplicateRequest = errs.New("duplicate request")
)

func (endpoint *Endpoint) VerifyOrderLimit(ctx context.Context, limit *pb.OrderLimit2) error {
	// sanity checks
	switch {
	case limit.Limit < 0:
		return ErrVerifyBadRequest.New("order limit is negative")
	case endpoint.Signer.ID() != limit.StorageNodeId:
		return ErrVerifyBadRequest.New("order intended for other storagenode: %v", limit.StorageNodeId)
	case endpoint.IsExpired(limit.PieceExpiration):
		return ErrVerifyBadRequest.New("piece expired: %v", limit.PieceExpiration)
	case endpoint.IsExpired(limit.OrderExpiration):
		return ErrVerifyBadRequest.New("order expired: %v", limit.OrderExpiration)

	case limit.SatelliteId.IsZero():
		return ErrVerifyBadRequest.New("missing satellite id")
	case limit.UplinkId.IsZero():
		return ErrVerifyBadRequest.New("missing uplink id")
	case len(limit.SatelliteSignature) == 0:
		return ErrVerifyBadRequest.New("satellite signature missing")
	}

	// either uplink or satellite can only make the request
	// TODO: should this check be based on the action?
	//       with macaroons we might not have either of them doing the action
	peer, err := identity.PeerIdentityFromContext(ctx)
	if err != nil || limit.UplinkId != peer.ID && limit.SatelliteId != peer.ID {
		return ErrVerifyNotAuthorized.New("uplink:%s satellite:%s sender %s", limit.UplinkId, limit.SatelliteId, peer.ID)
	}

	if err := endpoint.trust.VerifySatellite(ctx, limit.SatelliteID); err != nil {
		return ErrVerifyUntrusted.Wrap(err)
	}
	if err := endpoint.trust.VerifyUplink(ctx, limit.UplinkID); err != nil {
		return ErrVerifyUntrusted.Wrap(err)
	}
	if err := endpoint.VerifyOrderLimitSignature(ctx, limit); err != nil {
		return ErrVerifyUntrusted.Wrap(err)
	}

	if ok := endpoint.ActiveSerials.Add(limit.SatelliteID, limit.SerialNumber, limit.OrderExpiration); !ok {
		return ErrVerifyDuplicateRequest.Wrap(err)
	}

	return nil
}

func (endpoint *Endpoint) VerifyOrderLimitSignature(ctx context.Context, limit *pb.OrderLimit2) error {
	// TODO: remove signature before encoding and verifying
	bytes := encodeLimitAsBytes(limit)
	err := endpoint.Trust.VerifySatelliteSignature(ctx, bytes, limit.SatelliteId)
	return Error.Wrap(err)
}

func (endpoint *Endpoint) VerifyOrder(ctx context.Context, peer *identity.PeerIdentity, limit *pb.OrderLimit2, order *pb.Order2, largestOrderAmount int64) error {
	if order.SerialNumber != limit.SerialNumber {
		return ErrProtocol.New("order serial number changed during upload") // TODO: report grpc status bad message
	}
	if order.Amount < largestOrderAmount {
		return ErrProtocol.New("order contained smaller amount") // TODO: report grpc status bad message
	}
	if order.Amount > limit.Limit {
		return ErrProtocol.New("order exceeded allowed amount") // TODO: report grpc status bad message
	}
	if err := endpoint.VerifyOrderSignature(ctx, order, peer); err != nil {
		return ErrProtocol.New("order invalid signature") // TODO: report grpc status bad message
	}
	return nil
}

func (endpoint *Endpoint) IsExpired(date time.Time) bool {
	return date.Expired.Before(time.Now().Sub(endpoint.Config.ExpirationGracePeriod))
}
