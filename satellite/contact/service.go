// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package contact

import (
	"context"
	"sync"

	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/transport"
	"storj.io/storj/satellite/overlay"
)

// Error is the default error class for contact package
var Error = errs.Class("contact")

var mon = monkit.Package()

// Config contains configurable values for contact service
type Config struct {
	ExternalAddress string `user:"true" help:"the public address of the node, useful for nodes behind NAT" default:""`
}

// Conn represents a connection
type Conn struct {
	conn   *grpc.ClientConn
	client pb.ContactClient
}

// Service is the contact service between storage nodes and satellites
//
// architecture: Service
type Service struct {
	log *zap.Logger

	mutex sync.Mutex
	self  *overlay.NodeDossier

	overlay   *overlay.Service
	transport transport.Client
}

// NewService creates a new contact service
func NewService(log *zap.Logger, self *overlay.NodeDossier, overlay *overlay.Service, transport transport.Client) *Service {
	return &Service{
		log:       log,
		mutex:     sync.Mutex{},
		self:      self,
		overlay:   overlay,
		transport: transport,
	}
}

// Local returns the satellite node dossier
func (service *Service) Local() overlay.NodeDossier {
	service.mutex.Lock()
	defer service.mutex.Unlock()
	return *service.self
}

// FetchInfo connects to a node and returns its node info.
func (service *Service) FetchInfo(ctx context.Context, target pb.Node) (_ *pb.InfoResponse, err error) {
	conn, err := service.dialNode(ctx, target)
	if err != nil {
		return nil, err
	}

	resp, err := conn.client.RequestInf(ctx, &pb.InfoRequest{})

	return resp, errs.Combine(err, conn.close())
}

// dialNode dials the specified node.
func (service *Service) dialNode(ctx context.Context, target pb.Node) (_ *Conn, err error) {
	defer mon.Task()(&ctx)(&err)
	grpcconn, err := service.transport.DialNode(ctx, &target)
	return &Conn{
		conn:   grpcconn,
		client: pb.NewContactClient(grpcconn),
	}, err
}

// close disconnects this connection.
func (conn *Conn) close() error {
	return conn.conn.Close()
}
