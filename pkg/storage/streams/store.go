// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package streams

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/paths"
	ranger "storj.io/storj/pkg/ranger"
	"storj.io/storj/pkg/storage/segments"
	streamspb "storj.io/storj/protos/streams"
)

var mon = monkit.Package()

// Meta info about a segment
type Meta struct {
	Modified   time.Time
	Expiration time.Time
	Size       int64
	Data       []byte
}

// toMeta converts segment metadata to stream metadata
func toMeta(m segments.Meta) Meta {
	return Meta{
		Modified:   m.Modified,
		Expiration: m.Expiration,
		Size:       m.Size,
		Data:       m.Data,
	}
}

// Store for streams
type Store interface {
	Meta(ctx context.Context, path paths.Path) (Meta, error)
	Get(ctx context.Context, path paths.Path) (ranger.RangeCloser, Meta, error)
	Put(ctx context.Context, path paths.Path, data io.Reader,
		metadata []byte, expiration time.Time) (Meta, error)
	Delete(ctx context.Context, path paths.Path) error
	List(ctx context.Context, prefix, startAfter, endBefore paths.Path,
		recursive bool, limit int, metaFlags uint32) (items []ListItem,
		more bool, err error)
}

type streamStore struct {
	segments    segments.Store
	segmentSize int64
}

// NewStreams stuff
func NewStreams(segments segments.Store, segmentSize int64) (Store, error) {
	if segmentSize <= 0 {
		return nil, errors.New("segment size must be larger than 0")
	}
	return &streamStore{segments: segments, segmentSize: segmentSize}, nil
}

// Put breaks up data as it comes in into s.segmentSize length pieces, then
// store the first piece at s0/<path>, second piece at s1/<path>, and the
// *last* piece at l/<path>. Store the given metadata, along with the number
// of segments, in a new protobuf, in the metadata of l/<path>.
func (s *streamStore) Put(ctx context.Context, path paths.Path, data io.Reader,
	metadata []byte, expiration time.Time) (m Meta, err error) {
	defer mon.Task()(&ctx)(&err)

	identitySlice := make([]byte, 0)
	identityMeta := Meta{}
	var totalSegments int64
	var totalSize int64
	var lastSegmentSize int64

	awareLimitReader := EOFAwareReader(data)

	for !awareLimitReader.isEOF() {
		segmentPath := path.Prepend(fmt.Sprintf("s%d", totalSegments))
		segmentData := io.LimitReader(awareLimitReader, s.segmentSize)
		segmentMetatdata := identitySlice
		putMeta, err := s.segments.Put(ctx, segmentPath, segmentData,
			segmentMetatdata, expiration)
		if err != nil {
			return identityMeta, err
		}
		lastSegmentSize = putMeta.Size
		totalSize = totalSize + putMeta.Size
		totalSegments = totalSegments + 1
	}

	identitySegmentData := data
	lastSegmentPath := path.Prepend("l")

	md := streamspb.MetaStreamInfo{
		NumberOfSegments: totalSegments,
		SegmentsSize:     s.segmentSize,
		LastSegmentSize:  lastSegmentSize,
		MetaData:         metadata,
	}
	lastSegmentMetadata, err := proto.Marshal(&md)
	if err != nil {
		return identityMeta, err
	}

	putMeta, err := s.segments.Put(ctx, lastSegmentPath, identitySegmentData,
		lastSegmentMetadata, expiration)
	if err != nil {
		return identityMeta, err
	}
	totalSize = totalSize + putMeta.Size

	resultMeta := Meta{
		Modified:   putMeta.Modified,
		Expiration: expiration,
		Size:       totalSize,
		Data:       lastSegmentMetadata,
	}

	return resultMeta, nil
}

// Get returns a ranger that knows what the overall size is (from l/<path>)
// and then returns the appropriate data from segments s0/<path>, s1/<path>,
// ..., l/<path>.
func (s *streamStore) Get(ctx context.Context, path paths.Path) (
	rr ranger.RangeCloser, meta Meta, err error) {
	defer mon.Task()(&ctx)(&err)

	lastRangerCloser, lastSegmentMeta, err := s.segments.Get(ctx, path.Prepend("l"))
	if err != nil {
		return nil, Meta{}, err
	}

	msi := streamspb.MetaStreamInfo{}
	err = proto.Unmarshal(lastSegmentMeta.Data, &msi)
	if err != nil {
		return nil, Meta{}, err
	}

	newMeta := toMeta(lastSegmentMeta)

	var resRanger ranger.Ranger

	for i := 0; i < int(msi.NumberOfSegments); i++ {
		currentPath := fmt.Sprintf("s%d", i)
		rangeCloser, _, err := s.segments.Get(ctx, path.Prepend(currentPath))
		if err != nil {
			return nil, Meta{}, err
		}

		resRanger = ranger.Concat(resRanger, rangeCloser)
	}

	resRanger = ranger.Concat(resRanger, lastRangerCloser)

	return ranger.NopCloser(resRanger), newMeta, nil

}

func (s *streamStore) Meta(ctx context.Context, path paths.Path) (Meta, error) {
	segmentMeta, err := s.segments.Meta(ctx, path)
	if err != nil {
		return Meta{}, err
	}
	meta := toMeta(segmentMeta)

	return meta, nil
}

// Delete all the segments, with the last one last
func (s *streamStore) Delete(ctx context.Context, path paths.Path) (err error) {
	defer mon.Task()(&ctx)(&err)

	lastSegmentMeta, err := s.segments.Meta(ctx, path.Prepend("l"))
	if err != nil {
		return err
	}

	msi := streamspb.MetaStreamInfo{}
	err = proto.Unmarshal(lastSegmentMeta.Data, &msi)
	if err != nil {
		return err
	}

	for i := 0; i < int(msi.NumberOfSegments-1); i++ {
		currentPath := fmt.Sprintf("s%d", i)
		worked := s.segments.Delete(ctx, path.Prepend(currentPath))
		if worked != nil {
			return worked
		}
	}

	return s.segments.Delete(ctx, path.Prepend("l"))
}

// ListItem is a single item in a listing
type ListItem struct {
	Path paths.Path
	Meta Meta
}

// List all the paths inside l/, stripping off the l/ prefix
func (s *streamStore) List(ctx context.Context, prefix, startAfter, endBefore paths.Path,
	recursive bool, limit int, metaFlags uint32) (items []ListItem,
	more bool, err error) {
	defer mon.Task()(&ctx)(&err)

	lastSegmentMeta, err := s.segments.Meta(ctx, prefix.Prepend("l"))
	if err != nil {
		return nil, false, err
	}

	msi := streamspb.MetaStreamInfo{}
	err = proto.Unmarshal(lastSegmentMeta.Data, &msi)
	if err != nil {
		return nil, false, err
	}

	var resItems []ListItem
	var resMore bool

	for i := 0; i < int(msi.NumberOfSegments); i++ {
		items, more, err := s.segments.List(ctx, prefix, startAfter, endBefore, recursive, limit, metaFlags)
		if err != nil {
			return nil, false, err
		}
		for _, item := range items {
			newPath := strings.Split(item.Path.String(), fmt.Sprintf("s%d", i))
			newMeta := toMeta(item.Meta)

			resItems = append(resItems, ListItem{Path: newPath, Meta: newMeta})
			resMore = more
		}
	}

	return resItems, resMore, nil
}
