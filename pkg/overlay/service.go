// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package overlay

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/kademlia"
	"storj.io/storj/pkg/peertls"
	"storj.io/storj/pkg/process"
	proto "storj.io/storj/protos/overlay"
)

var (
	redisAddress, redisPassword, httpPort, bootstrapIP, bootstrapPort, localPort, boltdbPath string
	db                                                                                       int
	srvPort                                                                                  uint
	options                                                                                  peertls.TLSFileOptions
)

func init() {
	flag.StringVar(&httpPort, "httpPort", "", "The port for the health endpoint")
	flag.StringVar(&redisAddress, "redisAddress", "", "The <IP:PORT> string to use for connection to a redis cache")
	flag.StringVar(&redisPassword, "redisPassword", "", "The password used for authentication to a secured redis instance")
	flag.StringVar(&boltdbPath, "boltdbPath", defaultBoltDBPath(), "The path to the boltdb file that should be loaded or created")
	flag.IntVar(&db, "db", 0, "The network cache database")
	flag.UintVar(&srvPort, "srvPort", 8082, "Port to listen on")
	flag.StringVar(&bootstrapIP, "bootstrapIP", "", "Optional IP to bootstrap node against")
	flag.StringVar(&bootstrapPort, "bootstrapPort", "", "Optional port of node to bootstrap against")
	flag.StringVar(&localPort, "localPort", "8081", "Specify a different port to listen on locally")
	flag.StringVar(&options.RootCertRelPath, "tlsCertBasePath", "", "The base path for TLS certificates")
	flag.StringVar(&options.RootKeyRelPath, "tlsKeyBasePath", "", "The base path for TLS keys")
	flag.BoolVar(&options.Create, "tlsCreate", false, "If true, generate a new TLS cert/key files")
	flag.BoolVar(&options.Overwrite, "tlsOverwrite", false, "If true, overwrite existing TLS cert/key files")
}

func defaultBoltDBPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".storj", "overlaydb.db")
}

// NewServer creates a new Overlay Service Server
func NewServer(k *kademlia.Kademlia, cache *Cache, l *zap.Logger, m *monkit.Registry) *grpc.Server {
	grpcServer := grpc.NewServer()
	proto.RegisterOverlayServer(grpcServer, &Server{
		dht:     k,
		cache:   cache,
		logger:  l,
		metrics: m,
	})

	return grpcServer
}

// NewClient connects to grpc server at the provided address with the provided options
// returns a new instance of an overlay Client
func NewClient(serverAddr *string, opts ...grpc.DialOption) (proto.OverlayClient, error) {
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		return nil, err
	}

	return proto.NewOverlayClient(conn), nil
}

// NewTLSServer returns a newly initialized gRPC overlay server, configured with TLS
func NewTLSServer(k *kademlia.Kademlia, cache *Cache, l *zap.Logger, m *monkit.Registry, fopts peertls.TLSFileOptions) (_ *grpc.Server, _ error) {
	t, err := peertls.NewTLSFileOptions(
		fopts.RootCertRelPath,
		fopts.RootKeyRelPath,
		fopts.Create,
		fopts.Overwrite,
	)
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer(t.ServerOption())
	proto.RegisterOverlayServer(grpcServer, &Server{
		dht:     k,
		cache:   cache,
		logger:  l,
		metrics: m,
	})

	return grpcServer, nil
}

// NewTLSClient connects to grpc server at the provided address with the provided options plus TLS option(s)
// returns a new instance of an overlay Client
func NewTLSClient(serverAddr *string, fopts peertls.TLSFileOptions, opts ...grpc.DialOption) (proto.OverlayClient, error) {
	t, err := peertls.NewTLSFileOptions(
		fopts.RootCertRelPath,
		fopts.RootCertRelPath,
		fopts.Create,
		fopts.Overwrite,
	)
	if err != nil {
		return nil, err
	}

	opts = append(opts, t.DialOption())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		return nil, err
	}

	return proto.NewOverlayClient(conn), nil
}

// Service contains all methods needed to implement the process.Service interface
type Service struct {
	logger  *zap.Logger
	metrics *monkit.Registry
}

// Process is the main function that executes the service
func (s *Service) Process(ctx context.Context, _ *cobra.Command, _ []string) (
	err error) {
	// TODO
	// 1. Boostrap a node on the network
	// 2. Start up the overlay gRPC service
	// 3. Connect to Redis
	// 4. Boostrap Redis Cache

	// TODO(coyle): Should add the ability to pass a configuration to change the bootstrap node
	in, err := kademlia.GetIntroNode("", bootstrapIP, bootstrapPort)
	if err != nil {
		return err
	}

	id, err := kademlia.NewID()
	if err != nil {
		return err
	}

	kad, err := kademlia.NewKademlia(id, []proto.Node{*in}, "0.0.0.0", localPort)
	if err != nil {
		s.logger.Error("Failed to instantiate new Kademlia", zap.Error(err))
		return err
	}

	if err := kad.ListenAndServe(); err != nil {
		s.logger.Error("Failed to ListenAndServe on new Kademlia", zap.Error(err))
		return err
	}

	if err := kad.Bootstrap(ctx); err != nil {
		s.logger.Error("Failed to Bootstrap on new Kademlia", zap.Error(err))
		return err
	}

	// bootstrap cache
	var cache *Cache
	if viper.GetString("redisaddress") != "" {
		cache, err = NewRedisOverlayCache(viper.GetString("redisaddress"), redisPassword, db, kad)
		if err != nil {
			s.logger.Error("Failed to create a new redis overlay client", zap.Error(err))
			return err
		}
		s.logger.Info("starting overlay cache with redis")
	} else if viper.GetString("boltdbpath") != "" {
		cache, err = NewBoltOverlayCache(viper.GetString("boltdbpath"), kad)
		if err != nil {
			s.logger.Error("Failed to create a new boltdb overlay client", zap.Error(err))
			return err
		}
		s.logger.Info("starting overlay cache with boltDB")
	} else {
		return process.ErrUsage.New("You must specify one of `--boltdbPath` or `--redisAddress`")
	}

	if err := cache.Bootstrap(ctx); err != nil {
		s.logger.Error("Failed to boostrap cache", zap.Error(err))
		return err
	}

	// send off cache refreshes concurrently
	go cache.Refresh(ctx)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("srvport")))
	if err != nil {
		s.logger.Error("Failed to initialize TCP connection", zap.Error(err))
		return err
	}

	grpcServer := NewServer(kad, cache, s.logger, s.metrics)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "OK") })
	go func() { http.ListenAndServe(fmt.Sprintf(":%s", httpPort), mux) }()
	go cache.Walk(ctx)

	// If the passed context times out or is cancelled, shutdown the gRPC server
	go func() {
		if _, ok := <-ctx.Done(); !ok {
			grpcServer.GracefulStop()
		}
	}()

	// If `grpcServer.Serve(...)` returns an error, shutdown/cleanup the gRPC server
	defer grpcServer.GracefulStop()
	return grpcServer.Serve(lis)
}

// SetLogger adds the initialized logger to the Service
func (s *Service) SetLogger(l *zap.Logger) error {
	s.logger = l
	return nil
}

// SetMetricHandler adds the initialized metric handler to the Service
func (s *Service) SetMetricHandler(m *monkit.Registry) error {
	s.metrics = m
	return nil
}

// InstanceID implements Service.InstanceID
func (s *Service) InstanceID() string { return "" }
