// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package storage

import (
	"bytes"
	"errors"

	"github.com/zeebo/errs"
)

//ErrKeyNotFound used When something doesn't exist
var ErrKeyNotFound = errs.Class("key not found")

// ErrEmptyKey is returned when an empty key is used in Put
var ErrEmptyKey = errors.New("empty key")

// Key is the type for the keys in a `KeyValueStore`
type Key []byte

// Value is the type for the values in a `ValueValueStore`
type Value []byte

// Keys is the type for a slice of keys in a `KeyValueStore`
type Keys []Key

// Values is the type for a slice of Values in a `KeyValueStore`
type Values []Value

// Limit indicates how many keys to return when calling List
type Limit int

// Items keeps all ListItem
type Items []ListItem

// ListItem returns Key, Value, IsPrefix
type ListItem struct {
	Key      Key
	Value    Value
	IsPrefix bool
}

// KeyValueStore is an interface describing key/value stores like redis and boltdb
type KeyValueStore interface {
	// Put adds a value to store
	Put(Key, Value) error
	// Get gets a value to store
	Get(Key) (Value, error)
	// GetAll gets all values from the store
	GetAll(Keys) (Values, error)
	// Delete deletes key and the value
	Delete(Key) error
	// List lists all keys starting from start and upto limit items
	List(start Key, limit Limit) (Keys, error)
	// ReverseList lists all keys in revers order
	ReverseList(Key, Limit) (Keys, error)
	// Iterate iterates over collapsed items with prefix starting from first or the next key
	Iterate(prefix, first Key, delimiter byte, fn func(Iterator) error) error
	// IterateAll iterates over all items with prefix starting from first or the next key
	IterateAll(prefix, first Key, fn func(Iterator) error) error
	// Close closes the store
	Close() error
}

// Iterator iterates over a sequence of ListItems
type Iterator interface {
	// Next prepares the next list item
	// returns false when you reach final item
	Next(item *ListItem) bool
}

// IsZero returns true if the value struct is it's zero value
func (value Value) IsZero() bool {
	return len(value) == 0
}

// IsZero returns true if the key struct is it's zero value
func (key Key) IsZero() bool {
	return len(key) == 0
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for the Value type
func (value Value) MarshalBinary() ([]byte, error) {
	return value, nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for the Key type
func (key Key) MarshalBinary() ([]byte, error) {
	return key, nil
}

// ByteSlices converts a `Keys` struct to a slice of byte-slices (i.e. `[][]byte`)
func (keys Keys) ByteSlices() [][]byte {
	result := make([][]byte, len(keys))

	for key, val := range keys {
		result[key] = []byte(val)
	}

	return result
}

// String implements the Stringer interface
func (key Key) String() string { return string(key) }

// GetKeys gets all the Keys in []ListItem and converts them to Keys
func (items Items) GetKeys() Keys {
	if len(items) == 0 {
		return nil
	}
	var keys Keys
	for _, item := range items {
		keys = append(keys, item.Key)
	}
	return keys
}

// Len is the number of elements in the collection.
func (items Items) Len() int { return len(items) }

// Less reports whether the element with
// index i should sort before the element with index j.
func (items Items) Less(i, k int) bool { return items[i].Less(items[k]) }

// Swap swaps the elements with indexes i and j.
func (items Items) Swap(i, k int) { items[i], items[k] = items[k], items[i] }

// Less returns whether item should be sorted before b
func (item ListItem) Less(b ListItem) bool { return item.Key.Less(b.Key) }

// Less returns whether key should be sorted before b
func (key Key) Less(b Key) bool { return bytes.Compare([]byte(key), []byte(b)) < 0 }

// Equal returns whether key and b are equal
func (key Key) Equal(b Key) bool { return bytes.Equal([]byte(key), []byte(b)) }
