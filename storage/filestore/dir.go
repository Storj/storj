// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package filestore

import (
	"context"
	"encoding/base32"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/zeebo/errs"

	"storj.io/storj/storage"
)

const (
	blobPermission = 0600
	dirPermission  = 0700

	v0PieceFileSuffix      = ""
	v1PieceFileSuffix      = ".sj1"
	unknownPieceFileSuffix = "/..error_unknown_format../"
)

var pathEncoding = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567").WithPadding(base32.NoPadding)

// Dir represents single folder for storing blobs
type Dir struct {
	path string

	mu          sync.Mutex
	deleteQueue []string
}

// NewDir returns folder for storing blobs
func NewDir(path string) (*Dir, error) {
	dir := &Dir{
		path: path,
	}

	return dir, errs.Combine(
		os.MkdirAll(dir.blobsdir(), dirPermission),
		os.MkdirAll(dir.tempdir(), dirPermission),
		os.MkdirAll(dir.garbagedir(), dirPermission),
	)
}

// Path returns the directory path
func (dir *Dir) Path() string { return dir.path }

func (dir *Dir) blobsdir() string   { return filepath.Join(dir.path, "blobs") }
func (dir *Dir) tempdir() string    { return filepath.Join(dir.path, "temp") }
func (dir *Dir) garbagedir() string { return filepath.Join(dir.path, "garbage") }

// CreateTemporaryFile creates a preallocated temporary file in the temp directory
// prealloc preallocates file to make writing faster
func (dir *Dir) CreateTemporaryFile(ctx context.Context, prealloc int64) (_ *os.File, err error) {
	const preallocLimit = 5 << 20 // 5 MB
	if prealloc > preallocLimit {
		prealloc = preallocLimit
	}

	file, err := ioutil.TempFile(dir.tempdir(), "blob-*.partial")
	if err != nil {
		return nil, err
	}

	if prealloc >= 0 {
		if err := file.Truncate(prealloc); err != nil {
			return nil, errs.Combine(err, file.Close())
		}
	}
	return file, nil
}

// DeleteTemporary deletes a temporary file
func (dir *Dir) DeleteTemporary(ctx context.Context, file *os.File) (err error) {
	defer mon.Task()(&ctx)(&err)
	closeErr := file.Close()
	return errs.Combine(closeErr, os.Remove(file.Name()))
}

// blobToBasePath converts a blob reference to a filepath in permanent storage. This may not be the
// entire path; blobPathForFormatVersion() must also be used. This is a separate call because this
// part of the filepath is constant, and blobPathForFormatVersion may need to be called multiple
// times with different storage.FormatVersion values.
func (dir *Dir) blobToBasePath(ref storage.BlobRef) (string, error) {
	if !ref.IsValid() {
		return "", storage.ErrInvalidBlobRef.New("")
	}

	namespace := pathEncoding.EncodeToString(ref.Namespace)
	key := pathEncoding.EncodeToString(ref.Key)
	if len(key) < 3 {
		// ensure we always have enough characters to split [:2] and [2:]
		key = "11" + key
	}
	return filepath.Join(dir.blobsdir(), namespace, key[:2], key[2:]), nil
}

// blobPathForFormatVersion adjusts a bare blob path (as might have been generated by a call to
// blobToBasePath()) to what it should be for the given storage format version.
func blobPathForFormatVersion(path string, formatVersion storage.FormatVersion) string {
	switch formatVersion {
	case FormatV0:
		return path + v0PieceFileSuffix
	case FormatV1:
		return path + v1PieceFileSuffix
	}
	return path + unknownPieceFileSuffix
}

// blobToTrashPath converts a blob reference to a filepath in transient storage.
// The files in trash are deleted on an interval (in case the initial deletion didn't work for
// some reason).
func (dir *Dir) blobToTrashPath(ref storage.BlobRef) string {
	var name []byte
	name = append(name, ref.Namespace...)
	name = append(name, ref.Key...)
	return filepath.Join(dir.garbagedir(), pathEncoding.EncodeToString(name))
}

// Commit commits the temporary file to permanent storage.
func (dir *Dir) Commit(ctx context.Context, file *os.File, ref storage.BlobRef, formatVersion storage.FormatVersion) (err error) {
	defer mon.Task()(&ctx)(&err)
	position, seekErr := file.Seek(0, io.SeekCurrent)
	truncErr := file.Truncate(position)
	syncErr := file.Sync()
	chmodErr := os.Chmod(file.Name(), blobPermission)
	closeErr := file.Close()

	if seekErr != nil || truncErr != nil || syncErr != nil || chmodErr != nil || closeErr != nil {
		removeErr := os.Remove(file.Name())
		return errs.Combine(seekErr, truncErr, syncErr, chmodErr, closeErr, removeErr)
	}

	path, err := dir.blobToBasePath(ref)
	if err != nil {
		removeErr := os.Remove(file.Name())
		return errs.Combine(err, removeErr)
	}
	path = blobPathForFormatVersion(path, formatVersion)

	mkdirErr := os.MkdirAll(filepath.Dir(path), dirPermission)
	if os.IsExist(mkdirErr) {
		mkdirErr = nil
	}

	if mkdirErr != nil {
		removeErr := os.Remove(file.Name())
		return errs.Combine(mkdirErr, removeErr)
	}

	renameErr := rename(file.Name(), path)
	if renameErr != nil {
		removeErr := os.Remove(file.Name())
		return errs.Combine(renameErr, removeErr)
	}

	return nil
}

// Open opens the file with the specified ref. It may need to check in more than one location in
// order to find the blob, if it was stored with an older version of the storage node software.
// In cases where the storage format version of a blob is already known, OpenWithStorageFormat()
// will generally be a better choice.
func (dir *Dir) Open(ctx context.Context, ref storage.BlobRef) (_ *os.File, _ storage.FormatVersion, err error) {
	defer mon.Task()(&ctx)(&err)
	path, err := dir.blobToBasePath(ref)
	if err != nil {
		return nil, FormatV0, err
	}
	for formatVer := MaxFormatVersionSupported; formatVer >= MinFormatVersionSupported; formatVer-- {
		vPath := blobPathForFormatVersion(path, formatVer)
		file, err := openFileReadOnly(vPath, blobPermission)
		if err == nil {
			return file, formatVer, nil
		}
		if !os.IsNotExist(err) {
			return nil, FormatV0, Error.New("unable to open %q: %v", vPath, err)
		}
	}
	return nil, FormatV0, os.ErrNotExist
}

// OpenWithStorageFormat opens an already-located blob file with a known storage format version,
// which avoids the potential need to search through multiple storage formats to find the blob.
func (dir *Dir) OpenWithStorageFormat(ctx context.Context, blobRef storage.BlobRef, formatVer storage.FormatVersion) (_ *os.File, err error) {
	defer mon.Task()(&ctx)(&err)
	path, err := dir.blobToBasePath(blobRef)
	if err != nil {
		return nil, err
	}
	vPath := blobPathForFormatVersion(path, formatVer)
	file, err := openFileReadOnly(vPath, blobPermission)
	if err == nil {
		return file, nil
	}
	if os.IsNotExist(err) {
		return nil, err
	}
	return nil, Error.New("unable to open %q: %v", vPath, err)
}

// Stat looks up disk metadata on the blob file. It may need to check in more than one location
// in order to find the blob, if it was stored with an older version of the storage node software.
// In cases where the storage format version of a blob is already known, StatWithStorageFormat()
// will generally be a better choice.
func (dir *Dir) Stat(ctx context.Context, ref storage.BlobRef) (_ storage.BlobInfo, err error) {
	defer mon.Task()(&ctx)(&err)
	path, err := dir.blobToBasePath(ref)
	if err != nil {
		return nil, err
	}
	for formatVer := MaxFormatVersionSupported; formatVer >= MinFormatVersionSupported; formatVer-- {
		vPath := blobPathForFormatVersion(path, formatVer)
		stat, err := os.Stat(vPath)
		if err == nil {
			return newBlobInfo(ref, vPath, stat, formatVer), nil
		}
		if !os.IsNotExist(err) {
			return nil, Error.New("unable to stat %q: %v", vPath, err)
		}
	}
	return nil, os.ErrNotExist
}

// StatWithStorageFormat looks up disk metadata on the blob file with the given storage format
// version. This avoids the need for checking for the file in multiple different storage format
// types.
func (dir *Dir) StatWithStorageFormat(ctx context.Context, ref storage.BlobRef, formatVer storage.FormatVersion) (_ storage.BlobInfo, err error) {
	defer mon.Task()(&ctx)(&err)
	path, err := dir.blobToBasePath(ref)
	if err != nil {
		return nil, err
	}
	vPath := blobPathForFormatVersion(path, formatVer)
	stat, err := os.Stat(vPath)
	if err == nil {
		return newBlobInfo(ref, vPath, stat, formatVer), nil
	}
	if os.IsNotExist(err) {
		return nil, err
	}
	return nil, Error.New("unable to stat %q: %v", vPath, err)
}

// Delete deletes blobs with the specified ref (in all supported storage formats).
func (dir *Dir) Delete(ctx context.Context, ref storage.BlobRef) (err error) {
	defer mon.Task()(&ctx)(&err)
	pathBase, err := dir.blobToBasePath(ref)
	if err != nil {
		return err
	}
	trashPath := dir.blobToTrashPath(ref)

	var (
		moveErr        error
		combinedErrors errs.Group
	)

	// Try deleting all possible paths, starting with the oldest format version. It is more
	// likely, in the general case, that we will find the piece with the newest format version
	// instead, but if we iterate backward here then we run the risk of a race condition: the
	// piece might have existed with _SomeOldVer before the Delete call, and could then have
	// been updated atomically with _MaxVer concurrently while we were iterating. If we iterate
	// _forwards_, this race should not occur because it is assumed that pieces are never
	// rewritten with an _older_ storage format version.
	for i := MinFormatVersionSupported; i <= MaxFormatVersionSupported; i++ {
		verPath := blobPathForFormatVersion(pathBase, i)

		// move to trash folder, this is allowed for some OS-es
		moveErr = rename(verPath, trashPath)
		if os.IsNotExist(moveErr) {
			// no piece at that path; either it has a different storage format version or there
			// was a concurrent delete. (this function is expected by callers to return a nil
			// error in the case of concurrent deletes.)
			continue
		}
		if moveErr != nil {
			// piece could not be moved into the trash dir; we'll try removing it directly
			trashPath = verPath
		}

		// try removing the file
		err = os.Remove(trashPath)

		// ignore concurrent deletes
		if os.IsNotExist(err) {
			// something is happening at the same time as this; possibly a concurrent delete,
			// or possibly a rewrite of the blob. keep checking for more versions.
			continue
		}

		// the remove may have failed because of an open file handle. put it in a queue to be
		// retried later.
		if err != nil {
			dir.mu.Lock()
			dir.deleteQueue = append(dir.deleteQueue, trashPath)
			dir.mu.Unlock()
		}

		// ignore is-busy errors, they are still in the queue
		// but no need to notify
		if isBusy(err) {
			err = nil
		}
		combinedErrors.Add(err)
	}

	return combinedErrors.Err()
}

// GarbageCollect collects files that are pending deletion.
func (dir *Dir) GarbageCollect(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)
	offset := int(math.MaxInt32)
	// limited deletion loop to avoid blocking `Delete` for too long
	for offset >= 0 {
		dir.mu.Lock()
		limit := 100
		if offset >= len(dir.deleteQueue) {
			offset = len(dir.deleteQueue) - 1
		}
		for offset >= 0 && limit > 0 {
			path := dir.deleteQueue[offset]
			err := os.Remove(path)
			if os.IsNotExist(err) {
				err = nil
			}
			if err == nil {
				dir.deleteQueue = append(dir.deleteQueue[:offset], dir.deleteQueue[offset+1:]...)
			}

			offset--
			limit--
		}
		dir.mu.Unlock()
	}

	// remove anything left in the trashdir
	_ = removeAllContent(ctx, dir.garbagedir())
	return nil
}

const nameBatchSize = 1024

// ListNamespaces finds all known namespace IDs in use in local storage. They are not
// guaranteed to contain any blobs.
func (dir *Dir) ListNamespaces(ctx context.Context) (ids [][]byte, err error) {
	defer mon.Task()(&ctx)(&err)
	topBlobDir := dir.blobsdir()
	openDir, err := os.Open(topBlobDir)
	if err != nil {
		return nil, err
	}
	defer func() { err = errs.Combine(err, openDir.Close()) }()
	for {
		dirNames, err := openDir.Readdirnames(nameBatchSize)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if len(dirNames) == 0 {
			return ids, nil
		}
		for _, name := range dirNames {
			namespace, err := pathEncoding.DecodeString(name)
			if err != nil {
				// just an invalid directory entry, and not a namespace. probably
				// don't need to pass on this error
				continue
			}
			ids = append(ids, namespace)
		}
	}
}

// WalkNamespace executes walkFunc for each locally stored blob, stored with storage format V1 or
// greater, in the given namespace. If walkFunc returns a non-nil error, WalkNamespace will stop
// iterating and return the error immediately. The ctx parameter is intended specifically to allow
// canceling iteration early.
func (dir *Dir) WalkNamespace(ctx context.Context, namespace []byte, walkFunc func(storage.BlobInfo) error) (err error) {
	namespaceDir := pathEncoding.EncodeToString(namespace)
	nsDir := filepath.Join(dir.blobsdir(), namespaceDir)
	openDir, err := os.Open(nsDir)
	if err != nil {
		if os.IsNotExist(err) {
			// job accomplished: there are no blobs in this namespace!
			return nil
		}
		return err
	}
	defer func() { err = errs.Combine(err, openDir.Close()) }()
	for {
		// check for context done both before and after our readdir() call
		if err := ctx.Err(); err != nil {
			return err
		}
		subdirNames, err := openDir.Readdirnames(nameBatchSize)
		if err != nil && err != io.EOF {
			return err
		}
		if os.IsNotExist(err) || len(subdirNames) == 0 {
			return nil
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		for _, keyPrefix := range subdirNames {
			if len(keyPrefix) != 2 {
				// just an invalid subdir; could be garbage of many kinds. probably
				// don't need to pass on this error
				continue
			}
			err := dir.walkNamespaceWithPrefix(ctx, namespace, nsDir, keyPrefix, walkFunc)
			if err != nil {
				return err
			}
		}
	}
}

func (dir *Dir) walkNamespaceWithPrefix(ctx context.Context, namespace []byte, nsDir, keyPrefix string, walkFunc func(storage.BlobInfo) error) (err error) {
	keyDir := filepath.Join(nsDir, keyPrefix)
	openDir, err := os.Open(keyDir)
	if err != nil {
		return err
	}
	defer func() { err = errs.Combine(err, openDir.Close()) }()
	for {
		// check for context done both before and after our readdir() call
		if err := ctx.Err(); err != nil {
			return err
		}
		keyInfos, err := openDir.Readdir(nameBatchSize)
		if err != nil && err != io.EOF {
			return err
		}
		if os.IsNotExist(err) || len(keyInfos) == 0 {
			return nil
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		for _, keyInfo := range keyInfos {
			if keyInfo.Mode().IsDir() {
				continue
			}
			blobFileName := keyInfo.Name()
			encodedKey := keyPrefix + blobFileName
			formatVer := FormatV0
			if strings.HasSuffix(blobFileName, v1PieceFileSuffix) {
				formatVer = FormatV1
				encodedKey = encodedKey[0 : len(encodedKey)-len(v1PieceFileSuffix)]
			}
			key, err := pathEncoding.DecodeString(encodedKey)
			if err != nil {
				continue
			}
			ref := storage.BlobRef{
				Namespace: namespace,
				Key:       key,
			}
			fullPath := filepath.Join(keyDir, blobFileName)
			err = walkFunc(newBlobInfo(ref, fullPath, keyInfo, formatVer))
			if err != nil {
				return err
			}
			// also check for context done between every walkFunc callback.
			if err := ctx.Err(); err != nil {
				return err
			}
		}
	}
}

// removeAllContent deletes everything in the folder
func removeAllContent(ctx context.Context, path string) (err error) {
	defer mon.Task()(&ctx)(&err)
	dir, err := os.Open(path)
	if err != nil {
		return err
	}

	for {
		files, err := dir.Readdirnames(100)
		for _, file := range files {
			// the file might be still in use, so ignore the error
			_ = os.RemoveAll(filepath.Join(path, file))
		}
		if err == io.EOF || len(files) == 0 {
			return dir.Close()
		}
		if err != nil {
			return err
		}
	}
}

// DiskInfo contains statistics about this dir
type DiskInfo struct {
	ID             string
	AvailableSpace int64
}

// Info returns information about the current state of the dir
func (dir *Dir) Info() (DiskInfo, error) {
	path, err := filepath.Abs(dir.path)
	if err != nil {
		return DiskInfo{}, err
	}
	return diskInfoFromPath(path)
}

type blobInfo struct {
	ref           storage.BlobRef
	path          string
	fileInfo      os.FileInfo
	formatVersion storage.FormatVersion
}

func newBlobInfo(ref storage.BlobRef, path string, fileInfo os.FileInfo, formatVer storage.FormatVersion) storage.BlobInfo {
	return &blobInfo{
		ref:           ref,
		path:          path,
		fileInfo:      fileInfo,
		formatVersion: formatVer,
	}
}

func (info *blobInfo) BlobRef() storage.BlobRef {
	return info.ref
}

func (info *blobInfo) StorageFormatVersion() storage.FormatVersion {
	return info.formatVersion
}

func (info *blobInfo) Stat(ctx context.Context) (os.FileInfo, error) {
	return info.fileInfo, nil
}

func (info *blobInfo) FullPath(ctx context.Context) (string, error) {
	return info.path, nil
}
