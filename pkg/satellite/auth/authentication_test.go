// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type mockGenerator struct {
}

func (g *mockGenerator) Generate() (string, error) {
	return "", nil
}

type mockServerTransportStream struct {
	grpc.ServerTransportStream
}

func (s *mockServerTransportStream) SetHeader(md metadata.MD) error {
	return nil
}

func TestSatelliteAuthenticator(t *testing.T) {
	for _, tt := range []struct {
		APIKey string
		method string
		err    error
	}{
		// currently default apikey is empty
		{"", "/pointerdb", nil},
		{"wrong key", "/pointerdb", status.Errorf(codes.Unauthenticated, "Invalid API credential")},
		{"", "/otherservice", nil},
	} {
		authenticator := NewSatelliteAuthenticator(&mockGenerator{})

		// mock for method handler
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return nil, nil
		}

		ctx := context.Background()
		ctx = grpc.NewContextWithServerTransportStream(ctx, &mockServerTransportStream{})
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("apikey", tt.APIKey))
		info := &grpc.UnaryServerInfo{FullMethod: tt.method}

		_, err := authenticator(ctx, nil, info, handler)

		assert.Equal(t, err, tt.err)
	}

}
