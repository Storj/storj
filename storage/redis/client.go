// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package redis

import (
	"fmt"
	"sort"
	"time"

	"github.com/go-redis/redis"
	"github.com/zeebo/errs"
	"storj.io/storj/storage"
)

var (
	// Error is a redis error
	Error = errs.Class("redis error")
)

const (
	defaultNodeExpiration = 61 * time.Minute
	maxKeyLookup          = 100
)

// Client is the entrypoint into Redis
type Client struct {
	db  *redis.Client
	TTL time.Duration
}

// NewClient returns a configured Client instance, verifying a successful connection to redis
func NewClient(address, password string, db int) (*Client, error) {
	c := &Client{
		db: redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       db,
		}),
		TTL: defaultNodeExpiration,
	}

	// ping here to verify we are able to connect to redis with the initialized client.
	if err := c.db.Ping().Err(); err != nil {
		return nil, Error.New("ping failed: %v", err)
	}

	return c, nil
}

// Get looks up the provided key from redis returning either an error or the result.
func (c *Client) Get(key storage.Key) (storage.Value, error) {
	b, err := c.db.Get(string(key)).Bytes()

	if len(b) == 0 {
		return nil, storage.ErrKeyNotFound.New(key.String())
	}

	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}

		// TODO: log
		return nil, Error.New("get error: %v", err)
	}

	return b, nil
}

// Put adds a value to the provided key in redis, returning an error on failure.
func (c *Client) Put(key storage.Key, value storage.Value) error {
	if key == nil {
		return Error.New("invalid key")
	}

	v, err := value.MarshalBinary()

	if err != nil {
		return Error.New("put error: %v", err)
	}

	err = c.db.Set(key.String(), v, c.TTL).Err()
	if err != nil {
		return Error.New("put error: %v", err)
	}

	return nil
}

// List returns either a list of keys for which boltdb has values or an error.
func (c *Client) List(first storage.Key, limit storage.Limit) (storage.Keys, error) {
	return storage.ListKeys(c, first, limit)
}

// ReverseList returns either a list of keys for which redis has values or an error.
// Starts from startingKey and iterates backwards
func (c *Client) ReverseList(startingKey storage.Key, limit storage.Limit) (storage.Keys, error) {
	//TODO
	return storage.Keys{}, nil
}

// Delete deletes a key/value pair from redis, for a given the key
func (c *Client) Delete(key storage.Key) error {
	err := c.db.Del(key.String()).Err()
	if err != nil {
		return Error.New("delete error: %v", err)
	}

	return err
}

// Close closes a redis client
func (c *Client) Close() error {
	return c.db.Close()
}

// GetAll is the bulk method for gets from the redis data store
// The maximum keys returned will be 100. If more than that is requested an
// error will be returned
func (c *Client) GetAll(keys storage.Keys) (storage.Values, error) {
	lk := len(keys)
	if lk > maxKeyLookup {
		return nil, Error.New(fmt.Sprintf("requested %d keys, maximum is %d", lk, maxKeyLookup))
	}

	ks := make([]string, lk)
	for i, v := range keys {
		ks[i] = v.String()
	}

	vs, err := c.db.MGet(ks...).Result()
	if err != nil {
		return []storage.Value{}, err
	}

	values := []storage.Value{}
	for _, v := range vs {
		values = append(values, storage.Value([]byte(v.(string))))
	}
	return values, nil
}

// Iterate iterates over collapsed items with prefix starting from first or the next key
func (store *Client) Iterate(prefix, first storage.Key, delimiter byte, fn func(it storage.Iterator) error) error {
	var all storage.Items

	match := string(escapeMatch([]byte(prefix))) + "*"
	it := store.db.Scan(0, match, 0).Iterator()
	for it.Next() {
		key := it.Val()
		if storage.Key(key).Less(first) {
			continue
		}

		value, err := store.db.Get(key).Bytes()
		if err != nil {
			return err
		}

		all = append(all, storage.ListItem{
			Key:      storage.Key(key),
			Value:    storage.Value(value),
			IsPrefix: false,
		})
	}

	return fn(&storage.StaticIterator{
		Items: storage.SortAndCollapse(all, prefix, delimiter),
	})
}

// IterateAll iterates over all items with prefix starting from first or the next key
func (store *Client) IterateAll(prefix, first storage.Key, fn func(it storage.Iterator) error) error {
	var all storage.Items
	match := string(escapeMatch([]byte(prefix))) + "*"
	it := store.db.Scan(0, match, 0).Iterator()
	for it.Next() {
		key := it.Val()
		if storage.Key(key).Less(first) {
			continue
		}

		value, err := store.db.Get(key).Bytes()
		if err != nil {
			return err
		}

		all = append(all, storage.ListItem{
			Key:      storage.Key(key),
			Value:    storage.Value(value),
			IsPrefix: false,
		})
	}

	sort.Sort(all)

	return fn(&storage.StaticIterator{
		Items: all,
	})
}
