// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package redis

import (
	"fmt"
	"testing"

	"storj.io/storj/storage"
)

type RedisClientTest struct {
	*testing.T
	c storage.KeyValueStore
}

func NewRedisClientTest(t *testing.T) *RedisClientTest {
	c, err := NewClient("127.0.0.1:6379", "", 1)
	if err != nil {
		panic(err)
	}
	return &RedisClientTest{
		T: t,
		c: c,
	}
}

func (rt *RedisClientTest) Close() {
	rt.c.Close()
}

func (rt *RedisClientTest) HandleErr(err error, msg string) {
	rt.Error(msg)
	if err != nil {
		panic(err)
	}
	panic(msg)
}

func TestListWithoutStartKey(t *testing.T) {
	rt := NewRedisClientTest(t)
	defer rt.Close()

	if err := rt.c.Put(storage.Key([]byte("path/1")), []byte("pointer1")); err != nil {
		rt.HandleErr(err, "Failed to put")
	}
	if err := rt.c.Put(storage.Key([]byte("path/2")), []byte("pointer2")); err != nil {
		rt.HandleErr(err, "Failed to put")
	}
	if err := rt.c.Put(storage.Key([]byte("path/3")), []byte("pointer3")); err != nil {
		rt.HandleErr(err, "Failed to put")
	}

	keys, err := rt.c.List(nil, storage.Limit(3))
	if err != nil {
		rt.HandleErr(err, "Failed to list")
	}
	if fmt.Sprintf("%s", keys) != "[path/1 path/2 path/3]" {
		rt.HandleErr(nil, "Failed to list correct values")
	}
}

func TestListWithStartKey(t *testing.T) {
	rt := NewRedisClientTest(t)
	defer rt.Close()

	if err := rt.c.Put(storage.Key([]byte("path/1")), []byte("pointer1")); err != nil {
		rt.HandleErr(err, "Failed to put")
	}
	if err := rt.c.Put(storage.Key([]byte("path/2")), []byte("pointer2")); err != nil {
		rt.HandleErr(err, "Failed to put")
	}
	if err := rt.c.Put(storage.Key([]byte("path/3")), []byte("pointer3")); err != nil {
		rt.HandleErr(err, "Failed to put")
	}
	if err := rt.c.Put(storage.Key([]byte("path/4")), []byte("pointer4")); err != nil {
		rt.HandleErr(err, "Failed to put")
	}
	if err := rt.c.Put(storage.Key([]byte("path/5")), []byte("pointer5")); err != nil {
		rt.HandleErr(err, "Failed to put")
	}

	_, err := rt.c.List([]byte("path/2"), storage.Limit(2))
	if err != nil {
		rt.HandleErr(err, "Failed to list")
	}
}
