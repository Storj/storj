// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package contact

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"

	"storj.io/storj/pkg/identity"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/satellite/overlay"
)

// Endpoint implements the contact service Endpoints.
type Endpoint struct {
	log     *zap.Logger
	service *Service
}

// NewEndpoint returns a new contact service endpoint.
func NewEndpoint(log *zap.Logger, service *Service) *Endpoint {
	return &Endpoint{
		log:     log,
		service: service,
	}
}

// CheckIn is periodically called by storage nodes to keep the satellite informed of its existence,
// address, and operator information. In return, this satellite keeps the node informed of its
// reachability.
// When a node checks-in with the satellite, the satellite pings the node back to confirm they can
// successfully connect.
func (endpoint *Endpoint) CheckIn(ctx context.Context, req *pb.CheckInRequest) (_ *pb.CheckInResponse, err error) {
	defer mon.Task()(&ctx)(&err)

	peerID, err := peerIDFromContext(ctx)
	if err != nil {
		return nil, Error.Wrap(err)
	}
	nodeID := peerID.ID

	err = endpoint.service.peerIDs.Set(ctx, nodeID, peerID)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	lastIP, err := overlay.GetNetwork(ctx, req.Address)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	pingNodeSuccess, pingErrorMessage, err := endpoint.pingBack(ctx, req, nodeID)
	if err != nil {
		return nil, Error.Wrap(err)
	}
	nodeInfo := overlay.NodeCheckInInfo{
		NodeID: peerID.ID,
		Address: &pb.NodeAddress{
			Address:   req.Address,
			Transport: pb.NodeTransport_TCP_TLS_GRPC,
		},
		LastIP:   lastIP,
		IsUp:     pingNodeSuccess,
		Capacity: req.Capacity,
		Operator: req.Operator,
		Version:  req.Version,
	}
	err = endpoint.service.overlay.UpdateCheckIn(ctx, nodeInfo)

	endpoint.log.Debug("checking in", zap.String("node addr", req.Address), zap.Bool("ping node succes", pingNodeSuccess))
	return &pb.CheckInResponse{
		PingNodeSuccess:  pingNodeSuccess,
		PingErrorMessage: pingErrorMessage,
	}, nil
}

func (endpoint *Endpoint) pingBack(ctx context.Context, req *pb.CheckInRequest, peerID storj.NodeID) (bool, string, error) {
	client, err := newClient(ctx,
		endpoint.service.transport,
		req.Address,
		peerID,
	)
	if err != nil {
		// if this is a network error, then return the error otherwise just report internal error
		_, ok := err.(net.Error)
		if ok {
			return false, "", Error.New("failed to connect to %s: %v", req.Address, err)
		}
		endpoint.log.Info("pingBack internal error", zap.String("error", err.Error()))
		return false, "", Error.New("couldn't connect to client at addr: %s due to internal error.", req.Address)
	}

	pingNodeSuccess := true
	var pingErrorMessage string

	p := &peer.Peer{}
	_, err = client.pingNode(ctx, &pb.ContactPingRequest{}, grpc.Peer(p))
	if err != nil {
		pingNodeSuccess = false
		pingErrorMessage = "erroring while trying to pingNode due to internal error"
		_, ok := err.(net.Error)
		if ok {
			pingErrorMessage = fmt.Sprintf("network erroring while trying to pingNode: %v\n", err)
		}
	}

	return pingNodeSuccess, pingErrorMessage, nil
}

func peerIDFromContext(ctx context.Context) (*identity.PeerIdentity, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, Error.New("unable to get grpc peer from context")
	}
	peerIdentity, err := identity.PeerIdentityFromPeer(p)
	if err != nil {
		return nil, err
	}
	return peerIdentity, nil
}
