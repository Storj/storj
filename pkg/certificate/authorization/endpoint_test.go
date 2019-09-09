// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package authorization

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"storj.io/storj/internal/errs2"
	"storj.io/storj/internal/testcontext"
	"storj.io/storj/pkg/pb"
)

func TestEndpoint_Create(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()


	authorizationDB := newTestAuthDB(t, ctx)
	endpoint := NewEndpoint(zaptest.NewLogger(t), authorizationDB, nil)
	require.NotNil(t, endpoint)

	{ // new user, no existing authorization tokens
		userID := "new@mail.example"
		group, err := authorizationDB.Get(ctx, userID)
		require.NoError(t, err)
		require.Empty(t, group)

		res, err := endpoint.Create(ctx, &pb.AuthorizationRequest{UserId: userID})
		require.NoError(t, err)
		require.NotNil(t, res)

		token, err := ParseToken(res.Token)
		require.NoError(t, err)
		require.NotNil(t, token)

		require.Equal(t, userID, token.UserID)
	}

	{ // existing user with unclaimed authorization token
		userID := "old@mail.example"
		group, err := authorizationDB.Create(ctx, userID, 1)
		require.NoError(t, err)
		require.NotEmpty(t, group)
		require.Len(t, group, 1)

		existingAuth := group[0]

		res, err := endpoint.Create(ctx, &pb.AuthorizationRequest{UserId: userID})
		require.NoError(t, err)
		require.NotNil(t, res)

		token, err := ParseToken(res.Token)
		require.NoError(t, err)
		require.NotNil(t, token)

		require.Equal(t, userID, token.UserID)
		require.Equal(t, existingAuth.Token, *token)
	}
}

func TestEndpoint_Run_httpSuccess(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	require.NotNil(t, listener)

	endpoint := NewEndpoint(zaptest.NewLogger(t), newTestAuthDB(t, ctx), listener)
	require.NotNil(t, endpoint)

	ctx.Go(func() error {
		return errs2.IgnoreCanceled(endpoint.Run(ctx))
	})
	defer ctx.Check(endpoint.Close)

	userID := "user@mail.example"
	url := "http://" + listener.Addr().String() + "/v1/authorization/create"
	res, err := http.Post(url, "text/plain", bytes.NewBuffer([]byte(userID)))
	require.NoError(t, err)
	require.NotNil(t, res)

	require.Equal(t, http.StatusCreated, res.StatusCode)

	tokenBytes, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)
	require.NotEmpty(t, tokenBytes)
	require.NoError(t, res.Body.Close())

	token, err := ParseToken(string(tokenBytes))
	require.NoError(t, err)
	require.NotNil(t, token)

	require.Equal(t, userID, token.UserID)
}

func TestEndpoint_Run_httpErrors(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	require.NotNil(t, listener)

	endpoint := NewEndpoint(zaptest.NewLogger(t), newTestAuthDB(t, ctx), listener)
	require.NotNil(t, endpoint)

	ctx.Go(func() error {
		return errs2.IgnoreCanceled(endpoint.Run(ctx))
	})
	defer ctx.Check(endpoint.Close)

	baseURL := "http://" + listener.Addr().String()

	testCases := []struct {
		name       string
		userID     string
		urlPath    string
		httpMethod string
		statusCode int
	}{
		{
			"missing user ID",
			"",
			"/v1/authorization/create",
			http.MethodPost,
			400,
		},
		{
			"invalid http method (GET)",
			"user@mail.example",
			"/v1/authorization/create",
			http.MethodGet,
			400,
		},
		{
			"unsupported http method (PUT)",
			"user@mail.example",
			"/v1/authorization/create",
			http.MethodPut,
			400,
		},
		{
			"not found",
			"",
			"/",
			http.MethodPost,
			404,
		},
	}

	for _, testCase := range testCases {
		t.Log(testCase.name)
		url := baseURL + testCase.urlPath
		client := http.Client{}
		req, err := http.NewRequest(
			testCase.httpMethod, url,
			bytes.NewBuffer([]byte(testCase.userID)),
		)
		require.NoError(t, err)

		res, err := client.Do(req)
		require.NoError(t, err)
		require.NotNil(t, res)
		require.NoError(t, res.Body.Close())

		require.Equal(t, testCase.statusCode, res.StatusCode)
	}
}
