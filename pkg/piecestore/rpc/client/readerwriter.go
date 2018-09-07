// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package client

import (
	"fmt"
	"log"

	"storj.io/storj/internal/sync2"

	"github.com/gogo/protobuf/proto"
	"storj.io/storj/pkg/utils"
	pb "storj.io/storj/protos/piecestore"
)

// StreamWriter creates a StreamWriter for writing data to the piece store server
type StreamWriter struct {
	stream       pb.PieceStoreRoutes_StoreClient
	signer       *Client // We need this for signing
	totalWritten int64
	pba          *pb.PayerBandwidthAllocation
}

// Write Piece data to a piece store server upload stream
func (s *StreamWriter) Write(b []byte) (int, error) {
	updatedAllocation := s.totalWritten + int64(len(b))
	allocationData := &pb.RenterBandwidthAllocation_Data{
		PayerAllocation: s.pba,
		Total:           updatedAllocation,
	}

	serializedAllocation, err := proto.Marshal(allocationData)
	if err != nil {
		return 0, err
	}

	sig, err := s.signer.sign(serializedAllocation)
	if err != nil {
		return 0, err
	}

	msg := &pb.PieceStore{
		Piecedata: &pb.PieceStore_PieceData{Content: b},
		Bandwidthallocation: &pb.RenterBandwidthAllocation{
			Data: serializedAllocation, Signature: sig,
		},
	}

	s.totalWritten = updatedAllocation

	// Second we send the actual content
	if err := s.stream.Send(msg); err != nil {
		return 0, fmt.Errorf("%v.Send() = %v", s.stream, err)
	}

	return len(b), nil
}

// Close the piece store Write Stream
func (s *StreamWriter) Close() error {
	reply, err := s.stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("Route summary: %v", reply)

	return nil
}

// StreamReader is a struct for reading piece download stream from server
type StreamReader struct {
	throttle  *sync2.Throttle
	stream    pb.PieceStoreRoutes_RetrieveClient
	src       *utils.ReaderSource
	totalRead int64
	max       int64
	allocated int64
	err       error
}

// NewStreamReader creates a StreamReader for reading data from the piece store server
func NewStreamReader(signer *Client, stream pb.PieceStoreRoutes_RetrieveClient, pba *pb.PayerBandwidthAllocation, max int64) *StreamReader {
	sr := &StreamReader{
		throttle: sync2.NewThrottle(),
		stream:   stream,
		max:      max,
	}

	// Send signed allocations to the piece store server
	go func() {
		for {
			nextAllocSize := int64(signer.bandwidthMsgSize)
			if sr.allocated+int64(signer.bandwidthMsgSize) > max {
				nextAllocSize = max - sr.allocated
			}

			allocationData := &pb.RenterBandwidthAllocation_Data{
				PayerAllocation: pba,
				Total:           sr.allocated + nextAllocSize,
			}

			serializedAllocation, err := proto.Marshal(allocationData)
			if err != nil {
				sr.err = err
				break
			}

			sig, err := signer.sign(serializedAllocation)
			if err != nil {
				sr.err = err
				break
			}

			msg := &pb.PieceRetrieval{
				Bandwidthallocation: &pb.RenterBandwidthAllocation{
					Signature: sig,
					Data:      serializedAllocation,
				},
			}

			if err = stream.Send(msg); err != nil {
				sr.err = err
				break
			}

			sr.allocated += nextAllocSize

			if err = sr.throttle.ProduceAndWaitUntilBelow(nextAllocSize, max); err != nil {
				sr.err = err
				break
			}

		}
	}()

	sr.src = utils.NewReaderSource(func() ([]byte, error) {
		if sr.err != nil {
			return nil, sr.err
		}

		resp, err := stream.Recv()
		if err != nil {
			return nil, err
		}

		sr.totalRead += int64(len(resp.GetContent()))

		err = sr.throttle.Consume(int64(len(resp.GetContent())))
		if err != nil {
			return resp.GetContent(), err
		}

		return resp.GetContent(), nil
	})

	return sr
}

// Read Piece data from piece store server download stream
func (s *StreamReader) Read(b []byte) (int, error) {
	return s.src.Read(b)
}

// Close the piece store server Read Stream
func (s *StreamReader) Close() error {
	return s.stream.CloseSend()
}
