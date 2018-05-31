// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package netstate

import (
	"context"

	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc"

	pb "storj.io/storj/protos/netstate"
	"storj.io/storj/storage/boltdb"
	"storj.io/storj/netstate/auth"
)

// Server implements the network state RPC service
type Server struct {
	DB     DB
	logger *zap.Logger
}

// NewServer creates instance of Server
func NewServer(db DB, logger *zap.Logger) *Server {
	return &Server{
		DB:     db,
		logger: logger,
	}
}

// DB interface allows more modular unit testing
// and makes it easier in the future to substitute
// db clients other than bolt
type DB interface {
	Put(boltdb.PointerEntry) error
	Get([]byte) ([]byte, error)
	List() ([][]byte, error)
	Delete([]byte) error
}

func validateAuth(xAPIKeyBytes []byte) error {
	if !auth.ValidateAPIKey(string(xAPIKeyBytes)) {
		return grpc.Errorf(codes.Unauthenticated, "Invalid API credential")
	}
	return nil
}

// Put formats and hands off a file path to be saved to boltdb
func (s *Server) Put(ctx context.Context, putReq *pb.PutRequest) (*pb.PutResponse, error) {
	s.logger.Debug("entering netstate put")

	xAPIKeyBytes := []byte(putReq.XApiKey)
	if err := validateAuth(xAPIKeyBytes); err != nil {
		s.logger.Error("unauthorized request")
		return &pb.PutResponse{}, nil
	}

	pointerBytes, err := proto.Marshal(putReq.Pointer)
	if err != nil {
		s.logger.Error("err marshaling pointer", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pe := boltdb.PointerEntry{
		Path:    putReq.Path,
		Pointer: pointerBytes,
	}

	if err := s.DB.Put(pe); err != nil {
		s.logger.Error("err putting pointer", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	s.logger.Debug("put to the db: " + string(pe.Path))

	return &pb.PutResponse{}, nil
}

// Get formats and hands off a file path to get from boltdb
func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	s.logger.Debug("entering netstate get")
	
	xAPIKeyBytes := []byte(req.XApiKey)
	if err := validateAuth(xAPIKeyBytes); err != nil {
		s.logger.Error("unauthorized request")
		return &pb.GetResponse{
			Pointer: []byte("Unauthorized Request"),
		}, nil
	}
	
	pointerBytes, err := s.DB.Get(req.Path)

	if err != nil {
		s.logger.Error("err getting file", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.GetResponse{
		Pointer: pointerBytes,
	}, nil
}

// List calls the bolt client's List function and returns all Path keys in the Pointers bucket
func (s *Server) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	s.logger.Debug("entering netstate list")

	pathKeys, err := s.DB.List()
	
	xAPIKeyBytes := []byte(req.XApiKey)
	if err := validateAuth(xAPIKeyBytes); err != nil {
		c := []byte("Unauthorized Request")
		s.logger.Error("unauthorized request")
		return &pb.ListResponse{
			Paths: [][]byte{c},
		}, nil
	}

	if err != nil {
		s.logger.Error("err listing path keys", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	s.logger.Debug("path keys retrieved")
	return &pb.ListResponse{
		// pathKeys is an array of byte arrays
		Paths: pathKeys,
	}, nil
}

// Delete formats and hands off a file path to delete from boltdb
func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	s.logger.Debug("entering netstate delete")

	xAPIKeyBytes := []byte(req.XApiKey)
	if err := validateAuth(xAPIKeyBytes); err != nil {
		s.logger.Error("unauthorized request")

		return &pb.DeleteResponse{}, nil
	}

	err := s.DB.Delete(req.Path)
	if err != nil {
		s.logger.Error("err deleting pointer entry", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	s.logger.Debug("deleted pointer at path: " + string(req.Path))
	return &pb.DeleteResponse{}, nil
}
