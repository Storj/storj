// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package filestore

import (
	"context"
	"os"

	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/storage"
)

var (
	// Error is the default filestore error class
	Error = errs.Class("filestore error")

	mon = monkit.Package()

	_ storage.Blobs = (*Store)(nil)
)

// Store implements a blob store
type Store struct {
	dir *Dir
	log *zap.Logger

	cache storage.Blobs
}

// New creates a new disk blob store in the specified directory
func New(dir *Dir, log *zap.Logger) *Store {
	return &Store{dir: dir, log: log}
}

// NewAt creates a new disk blob store in the specified directory
func NewAt(path string, log *zap.Logger) (*Store, error) {
	dir, err := NewDir(path)
	if err != nil {
		return nil, Error.Wrap(err)
	}
	return &Store{dir: dir, log: log}, nil
}

// Close closes the store.
func (store *Store) Close() error { return nil }

// Cache returns the storage cache
func (store *Store) Cache(ctx context.Context) storage.BlobsUsageCache {
	return store.cache
}

// Open loads blob with the specified hash
func (store *Store) Open(ctx context.Context, ref storage.BlobRef) (_ storage.BlobReader, err error) {
	defer mon.Task()(&ctx)(&err)
	file, formatVer, err := store.dir.Open(ctx, ref)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, err
		}
		return nil, Error.Wrap(err)
	}
	return newBlobReader(file, formatVer), nil
}

// OpenSpecific loads the already-located blob, avoiding the potential need to check multiple
// storage formats to find the blob.
func (store *Store) OpenSpecific(ctx context.Context, blobRef storage.BlobRef, formatVer storage.FormatVersion) (_ storage.BlobReader, err error) {
	defer mon.Task()(&ctx)(&err)
	file, err := store.dir.OpenSpecific(ctx, blobRef, formatVer)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, err
		}
		return nil, Error.Wrap(err)
	}
	return newBlobReader(file, formatVer), nil
}

// Lookup looks up disk metadata on the blob file
func (store *Store) Lookup(ctx context.Context, ref storage.BlobRef) (_ storage.StoredBlobAccess, err error) {
	defer mon.Task()(&ctx)(&err)
	access, err := store.dir.Lookup(ctx, ref)
	return access, Error.Wrap(err)
}

// LookupSpecific looks up disk metadata on the blob file with the given storage format version
func (store *Store) LookupSpecific(ctx context.Context, ref storage.BlobRef, formatVer storage.FormatVersion) (_ storage.StoredBlobAccess, err error) {
	defer mon.Task()(&ctx)(&err)
	access, err := store.dir.LookupSpecific(ctx, ref, formatVer)
	return access, Error.Wrap(err)
}

// Delete deletes blobs with the specified ref
func (store *Store) Delete(ctx context.Context, ref storage.BlobRef) (err error) {
	defer mon.Task()(&ctx)(&err)
	err = store.dir.Delete(ctx, ref)
	return Error.Wrap(err)
}

// GarbageCollect tries to delete any files that haven't yet been deleted
func (store *Store) GarbageCollect(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)
	err = store.dir.GarbageCollect(ctx)
	return Error.Wrap(err)
}

// Create creates a new blob that can be written
// optionally takes a size argument for performance improvements, -1 is unknown size
func (store *Store) Create(ctx context.Context, ref storage.BlobRef, size int64) (_ storage.BlobWriter, err error) {
	defer mon.Task()(&ctx)(&err)
	file, err := store.dir.CreateTemporaryFile(ctx, size)
	if err != nil {
		return nil, Error.Wrap(err)
	}
	return newBlobWriter(ref, store, storage.MaxStorageFormatVersionSupported, file), nil
}

// SpaceUsed adds up the space used in all namespaces for blob storage
func (store *Store) SpaceUsed(ctx context.Context) (space int64, err error) {
	defer mon.Task()(&ctx)(&err)

	var totalSpaceUsed int64
	namespaces, err := store.GetAllNamespaces(ctx)
	if err != nil {
		return 0, Error.New("failed to enumerate namespaces: %v", err)
	}
	for _, namespace := range namespaces {
		used, err := store.SpaceUsedInNamespace(ctx, namespace)
		if err != nil {
			return 0, Error.New("failed to sum space used: %v", err)
		}
		totalSpaceUsed += used
	}
	return totalSpaceUsed, nil
}

// SpaceUsedTotalAndByNamespace adds up the space used by and for all namespaces for blob storage
func (store *Store) SpaceUsedTotalAndByNamespace(ctx context.Context) (_ int64, _ map[string]int64, err error) {
	defer mon.Task()(&ctx)(&err)

	var totalSpaceUsed int64
	var totalSpaceUsedByNamespace = map[string]int64{}
	namespaces, err := store.GetAllNamespaces(ctx)
	if err != nil {
		return totalSpaceUsed, totalSpaceUsedByNamespace, Error.New("failed to enumerate namespaces: %v", err)
	}
	for _, namespace := range namespaces {
		used, err := store.SpaceUsedInNamespace(ctx, namespace)
		if err != nil {
			return totalSpaceUsed, totalSpaceUsedByNamespace, Error.New("failed to sum space used: %v", err)
		}
		totalSpaceUsed += used
		totalSpaceUsedByNamespace[string(namespace)] = used
	}
	return totalSpaceUsed, totalSpaceUsedByNamespace, nil
}

// SpaceUsedInNamespace adds up how much is used in the given namespace for blob storage
func (store *Store) SpaceUsedInNamespace(ctx context.Context, namespace []byte) (int64, error) {
	var totalUsed int64
	err := store.ForAllKeysInNamespace(ctx, namespace, func(access storage.StoredBlobAccess) error {
		statInfo, statErr := access.Stat(ctx)
		if statErr != nil {
			store.log.Error("failed to stat blob", zap.Binary("namespace", namespace), zap.Binary("key", access.BlobRef().Key), zap.Error(statErr))
			// keep iterating; we want a best effort total here.
			return nil
		}
		totalUsed += statInfo.Size()
		return nil
	})
	if err != nil {
		return 0, err
	}
	return totalUsed, nil
}

// FreeSpace returns how much space left in underlying directory
func (store *Store) FreeSpace() (int64, error) {
	info, err := store.dir.Info()
	if err != nil {
		return 0, err
	}
	return info.AvailableSpace, nil
}

// GetAllNamespaces finds all known namespace IDs in use in local storage. They are not
// guaranteed to contain any blobs.
func (store *Store) GetAllNamespaces(ctx context.Context) (ids [][]byte, err error) {
	return store.dir.GetAllNamespaces(ctx)
}

// ForAllKeysInNamespace executes doForEach for each locally stored blob in the given
// namespace. If doForEach returns a non-nil error, ForAllKeysInNamespace will stop
// iterating and return the error immediately.
func (store *Store) ForAllKeysInNamespace(ctx context.Context, namespace []byte, doForEach func(storage.StoredBlobAccess) error) (err error) {
	return store.dir.ForAllKeysInNamespace(ctx, namespace, doForEach)
}

// StoreForTest is a wrapper for Store that also allows writing new V0 blobs (in order to test
// situations involving those)
type StoreForTest struct {
	*Store
}

// CreateV0 creates a new V0 blob that can be written. This is only appropriate in test situations.
func (testStore *StoreForTest) CreateV0(ctx context.Context, ref storage.BlobRef) (_ storage.BlobWriter, err error) {
	defer mon.Task()(&ctx)(&err)

	file, err := testStore.dir.CreateTemporaryFile(ctx, -1)
	if err != nil {
		return nil, Error.Wrap(err)
	}
	return newBlobWriter(ref, testStore.Store, storage.FormatV0, file), nil
}
