// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package client

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/mr-tron/base58/base58"

	"storj.io/storj/pkg/ranger"
	pb "storj.io/storj/protos/piecestore"
)

type PSClient interface {
	Meta(ctx context.Context, id PieceID) (*pb.PieceSummary, error)
	Put(ctx context.Context, id PieceID, data io.Reader, ttl time.Time) error
	Get(ctx context.Context, id PieceID, offset, length int64) (ranger.RangeCloser, error)
	Delete(ctx context.Context, pieceID PieceID) error
}

// PieceID - Id for piece
type PieceID string

// String -- Get String from PieceID
func (id PieceID) String() string {
	return string(id)
}

// Client -- Struct Info needed for protobuf api calls
type Client struct {
	route pb.PieceStoreRoutesClient
}

// NewPSClient -- Initilize Client
func NewPSClient(conn *grpc.ClientConn) PSClient {
	return &Client{route: pb.NewPieceStoreRoutesClient(conn)}
}

// NewCustomRoute creates new Client with custom route interface
func NewCustomRoute(route pb.PieceStoreRoutesClient) *Client {
	return &Client{route: route}
}

// Meta -- Request info about a piece by Id
func (client *Client) Meta(ctx context.Context, id PieceID) (*pb.PieceSummary, error) {
	return client.route.Piece(ctx, &pb.PieceId{Id: id.String()})
}

// Put -- Upload Piece to Server
func (client *Client) Put(ctx context.Context, id PieceID, data io.Reader, ttl time.Time) error {
	stream, err := client.route.Store(ctx)
	if err != nil {
		return err
	}

	// SSend preliminary data
	if err := stream.Send(&pb.PieceStore{Id: id.String(), Ttl: ttl.Unix()}); err != nil {
		stream.CloseAndRecv()
		return fmt.Errorf("%v.Send() = %v", stream, err)
	}

	writer := &StreamWriter{stream: stream}

	defer writer.Close()

	_, err = io.Copy(writer, data)

	return err
}

// Get -- Begin Download Piece from Server
func (client *Client) Get(ctx context.Context, id PieceID, offset, length int64) (ranger.RangeCloser, error) {
	stream, err := client.route.Retrieve(ctx, &pb.PieceRetrieval{Id: id.String(), Size: length, Offset: offset})
	if err != nil {
		return nil, err
	}

	return NewStreamReader(stream), nil
}

// Delete -- Delete Piece From Server
func (client *Client) Delete(ctx context.Context, id PieceID) error {
	reply, err := client.route.Delete(ctx, &pb.PieceDelete{Id: id.String()})
	if err != nil {
		return err
	}
	log.Printf("Route summary : %v", reply)
	return nil
}

func (id PieceID) IsValid() bool {
	return len(id) >= 20
}

// NewPieceID creates random id
func NewPieceID() PieceID {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return PieceID(base58.Encode(b))
}
