// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package streams

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/paths"
	"storj.io/storj/pkg/ranger"
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

// ListItem is a single item in a listing
type ListItem struct {
	Path paths.Path
	Meta Meta
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
	segments    segment.Store
	segmentSize int64
}

// NewStreams stuff
func NewStreams(segments segment.Store, segmentSize int64) (Store, error) {
	if segmentSize < 0 {
		return nil, errors.New("Segment size must be larger than 0")
	}
	return &streamStore{segments: segments, segmentSize: segmentSize}, nil
}

func (s *streamStore) Put(ctx context.Context, path paths.Path, data io.Reader,
	metadata []byte, expiration time.Time) (m Meta, err error) {
	defer mon.Task()(&ctx)(&err)

	// TODO: break up data as it comes in into s.segmentSize length pieces, then
	// store the first piece at s0/<path>, second piece at s1/<path>, and the
	// *last* piece at l/<path>. Store the given metadata, along with the number
	// of segments, in a new protobuf, in the metadata of l/<path>.

	identitySlice := make([]byte, 0)
	identityMeta := Meta{}
	var totalSegments int64
	var totalSize int64
	var lastSegmentSize int64
	totalSegments = 0
	totalSize = 0
	lastSegmentSize = 0

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
	}

	totalSegments = totalSegments + 1
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
		Modified:   time.Now(),
		Expiration: expiration,
		Size:       totalSize,
		Data:       lastSegmentMetadata,
	}

	return resultMeta, nil
}

// EOFAwareLimitReader holds reader and status of EOF
type EOFAwareLimitReader struct {
	reader io.Reader
	eof    bool
}

// EOFAwareReader keeps track of the state, has the internal reader reached EOF
func EOFAwareReader(r io.Reader) *EOFAwareLimitReader {
	return &EOFAwareLimitReader{reader: r, eof: false}
}

func (r *EOFAwareLimitReader) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	if err == io.EOF {
		r.eof = true
	}
	return n, err
}

func (r *EOFAwareLimitReader) isEOF() bool {
	return r.eof
}

func (s *streamStore) Get(ctx context.Context, path paths.Path) (ranger.Ranger, Meta, error) {
	defer mon.Task()(&ctx)(&err)

	// TODO: return a ranger that knows what the overall size is (from l/<path>)
	// and then returns the appropriate data from segments s0/<path>, s1/<path>,
	// ..., l/<path>.

	lastRanger, lastMeta, err := s.segments.Get(ctx, path.Prepend("l"))
	if err != nil {
		return nil, m, err
	}
	totalSize = lastMeta.Size
	sizePerSegment = float64(totalSize) / float64(s.segmentSize)
	stringSegmentsSize = fmt.Sprintf("%f", sizePerSegment)
	segmentSizeSlice = strings.Split(stringSegmentsSize, ".")
	perfectSizedSegments, err = strconv.ParseInt(segmentSizeSlice[0], 10, 64)
	lastSegmentSize, err = strconv.ParseInt(segmentSizeSlice[1], 10, 64)

	rv, meta, err := s.segments.Get(ctx, path)
	return rv, meta.Meta, nil
}

/*
func (s *streamStore) Meta(ctx context.Context, path paths.Path) (Meta, error) {

}

func (s *streamStore) Delete(ctx context.Context, path dtypes.Path) (err error) {
	defer mon.Task()(&ctx)(&err)

	// TODO: delete all the segments, with the last one last

	return s.store.Delete(ctx, path)
}

func (s *streamStore) List(ctx context.Context, startingPath, endingPath dtypes.Path) (paths []dtypes.Path, truncated bool, err error) {
	defer mon.Task()(&ctx)(&err)

	// TODO: list all the paths inside l/, stripping off the l/ prefix

	return s.store.List(ctx, startingPath, endingPath)
}
*/
