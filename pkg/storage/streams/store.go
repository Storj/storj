// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package streams

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	proto "github.com/gogo/protobuf/proto"
	"github.com/zeebo/errs"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/paths"
	"storj.io/storj/pkg/pb"
	ranger "storj.io/storj/pkg/ranger"
	"storj.io/storj/pkg/storage/meta"
	"storj.io/storj/pkg/storage/segments"
)

var mon = monkit.Package()

// Meta info about a segment
type Meta struct {
	Modified   time.Time
	Expiration time.Time
	Size       int64
	Data       []byte
}

// convertMeta converts segment metadata to stream metadata
func convertMeta(segmentMeta segments.Meta) (Meta, error) {
	msi := pb.MetaStreamInfo{}
	err := proto.Unmarshal(segmentMeta.Data, &msi)
	if err != nil {
		return Meta{}, err
	}

	return Meta{
		Modified:   segmentMeta.Modified,
		Expiration: segmentMeta.Expiration,
		Size:       ((msi.NumberOfSegments - 1) * msi.SegmentsSize) + msi.LastSegmentSize,
		Data:       msi.Metadata,
	}, nil
}

// Store interface methods for streams to satisfy to be a store
type Store interface {
	Meta(ctx context.Context, path paths.Path) (Meta, error)
	Get(ctx context.Context, path paths.Path) (ranger.Ranger, Meta, error)
	Put(ctx context.Context, path paths.Path, data io.Reader,
		metadata []byte, expiration time.Time) (Meta, error)
	Delete(ctx context.Context, path paths.Path) error
	List(ctx context.Context, prefix, startAfter, endBefore paths.Path,
		recursive bool, limit int, metaFlags uint32) (items []ListItem,
		more bool, err error)
}

// streamStore is a store for streams
type streamStore struct {
	segments    segments.Store
	segmentSize int64
}

// NewStreamStore stuff
func NewStreamStore(segments segments.Store, segmentSize int64) (Store, error) {
	if segmentSize <= 0 {
		return nil, errs.New("segment size must be larger than 0")
	}
	return &streamStore{segments: segments, segmentSize: segmentSize}, nil
}

func collectErrors(errs <-chan error, size int) []error {
	var result []error
	for i := 0; i < size; i++ {
		err := <-errs
		if err != nil {
			result = append(result, err)
		}
	}
	return result
}

// Put breaks up data as it comes in into s.segmentSize length pieces, then
// store the first piece at s0/<path>, second piece at s1/<path>, and the
// *last* piece at l/<path>. Store the given metadata, along with the number
// of segments, in a new protobuf, in the metadata of l/<path>.
func (s *streamStore) Put(ctx context.Context, path paths.Path, data io.Reader,
	metadata []byte, expiration time.Time) (m Meta, err error) {
	defer mon.Task()(&ctx)(&err)

	var totalSegments int64
	var totalSize int64
	var lastSegmentSize int64

	//ctx, cancel := context.WithCancel(ctx)
	log.Println("inside object store ")
	/* create a signal of type os.Signal */
	c := make(chan os.Signal, 0x01)

	/* register for the os signals */
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// go func() {
	// 	<-c
	// 	log.Println("cancelling .......")
	// 	signal.Stop(c)
	// 	log.Println("totalSeg....", totalSegments)
	// 	log.Println("path....", path)
	// }()

	defer func() {
		select {
		case <-ctx.Done():
			log.Println("cancelling .......")
			signal.Stop(c)
			log.Println("totalSeg....", totalSegments)
			log.Println("path....", path)
			break
		default:
			return
		}

		segErrs := make(chan error, int(totalSegments))
		for i := 0; i < int(totalSegments); i++ {
			log.Println("inside for loop seg#....", i)
			currentPath := fmt.Sprintf("s%d", i)
			log.Println("deleting segment and path ... ", currentPath, path.Prepend(currentPath))

			//go func() {
			segErrs <- s.segments.Delete(ctx, path.Prepend(currentPath))
			log.Println("----> KISHORE <---deleted path", path.Prepend(currentPath))
			//segErrs <- err
			//}()
		}
		//go func() {
		for i := 0; i < int(totalSegments); i++ {
			fmt.Println("err from deleting", <-segErrs, i)
		}
		//}()

		//cancel()
		log.Printf("cleaned up the partial uploads !!!!!!!!!ctx.Done()")
		return
	}()

	awareLimitReader := EOFAwareReader(data)

	for !awareLimitReader.isEOF() && !awareLimitReader.hasError() {
		segmentPath := path.Prepend(fmt.Sprintf("s%d", totalSegments))
		segmentData := io.LimitReader(awareLimitReader, s.segmentSize)

		putMeta, err := s.segments.Put(ctx, segmentPath, segmentData, nil, expiration)
		if err != nil {
			return Meta{}, err
		}
		lastSegmentSize = putMeta.Size
		totalSize = totalSize + putMeta.Size
		totalSegments = totalSegments + 1
	}
	if awareLimitReader.hasError() {
		return Meta{}, awareLimitReader.err
	}

	lastSegmentPath := path.Prepend("l")

	md := pb.MetaStreamInfo{
		NumberOfSegments: totalSegments,
		SegmentsSize:     s.segmentSize,
		LastSegmentSize:  lastSegmentSize,
		Metadata:         metadata,
	}
	lastSegmentMetadata, err := proto.Marshal(&md)
	if err != nil {
		return Meta{}, err
	}

	putMeta, err := s.segments.Put(ctx, lastSegmentPath, data,
		lastSegmentMetadata, expiration)
	if err != nil {
		return Meta{}, err
	}
	totalSize = totalSize + putMeta.Size

	resultMeta := Meta{
		Modified:   putMeta.Modified,
		Expiration: expiration,
		Size:       totalSize,
		Data:       metadata,
	}

	return resultMeta, nil
}

// Get returns a ranger that knows what the overall size is (from l/<path>)
// and then returns the appropriate data from segments s0/<path>, s1/<path>,
// ..., l/<path>.
func (s *streamStore) Get(ctx context.Context, path paths.Path) (
	rr ranger.Ranger, meta Meta, err error) {
	defer mon.Task()(&ctx)(&err)

	lastRangerCloser, lastSegmentMeta, err := s.segments.Get(ctx, path.Prepend("l"))
	if err != nil {
		return nil, Meta{}, err
	}

	msi := pb.MetaStreamInfo{}
	err = proto.Unmarshal(lastSegmentMeta.Data, &msi)
	if err != nil {
		return nil, Meta{}, err
	}

	newMeta, err := convertMeta(lastSegmentMeta)
	if err != nil {
		return nil, Meta{}, err
	}

	var rangers []ranger.Ranger

	for i := int64(0); i < msi.NumberOfSegments; i++ {
		currentPath := fmt.Sprintf("s%d", i)
		size := msi.SegmentsSize
		if i == msi.NumberOfSegments-1 {
			size = msi.LastSegmentSize
		}
		rr := &lazySegmentRanger{
			segments: s.segments,
			path:     path.Prepend(currentPath),
			size:     size,
		}
		rangers = append(rangers, rr)
	}

	rangers = append(rangers, lastRangerCloser)

	catRangers := ranger.Concat(rangers...)

	return catRangers, newMeta, nil
}

// Meta implements Store.Meta
func (s *streamStore) Meta(ctx context.Context, path paths.Path) (Meta, error) {
	segmentMeta, err := s.segments.Meta(ctx, path.Prepend("l"))
	if err != nil {
		return Meta{}, err
	}

	meta, err := convertMeta(segmentMeta)
	if err != nil {
		return Meta{}, err
	}

	return meta, nil
}

// Delete all the segments, with the last one last
func (s *streamStore) Delete(ctx context.Context, path paths.Path) (err error) {
	defer mon.Task()(&ctx)(&err)

	lastSegmentMeta, err := s.segments.Meta(ctx, path.Prepend("l"))
	if err != nil {
		return err
	}

	msi := pb.MetaStreamInfo{}
	err = proto.Unmarshal(lastSegmentMeta.Data, &msi)
	if err != nil {
		return err
	}

	for i := 0; i < int(msi.NumberOfSegments); i++ {
		currentPath := fmt.Sprintf("s%d", i)
		err := s.segments.Delete(ctx, path.Prepend(currentPath))
		if err != nil {
			return err
		}
	}

	return s.segments.Delete(ctx, path.Prepend("l"))
}

// ListItem is a single item in a listing
type ListItem struct {
	Path     paths.Path
	Meta     Meta
	IsPrefix bool
}

// List all the paths inside l/, stripping off the l/ prefix
func (s *streamStore) List(ctx context.Context, prefix, startAfter, endBefore paths.Path,
	recursive bool, limit int, metaFlags uint32) (items []ListItem,
	more bool, err error) {
	defer mon.Task()(&ctx)(&err)

	if metaFlags&meta.Size != 0 {
		// Calculating the stream's size require also the user-defined metadata,
		// where stream store keeps info about the number of segments and their size.
		metaFlags |= meta.UserDefined
	}

	segments, more, err := s.segments.List(ctx, prefix.Prepend("l"), startAfter, endBefore, recursive, limit, metaFlags)
	if err != nil {
		return nil, false, err
	}

	items = make([]ListItem, len(segments))
	for i, item := range segments {
		newMeta, err := convertMeta(item.Meta)
		if err != nil {
			return nil, false, err
		}
		items[i] = ListItem{Path: item.Path, Meta: newMeta, IsPrefix: item.IsPrefix}
	}

	return items, more, nil
}

type lazySegmentRanger struct {
	ranger   ranger.Ranger
	segments segments.Store
	path     paths.Path
	size     int64
}

// Size implements Ranger.Size
func (lr *lazySegmentRanger) Size() int64 {
	return lr.size
}

// Range implements Ranger.Range to be lazily connected
func (lr *lazySegmentRanger) Range(ctx context.Context, offset, length int64) (io.ReadCloser, error) {
	if lr.ranger == nil {
		rr, _, err := lr.segments.Get(ctx, lr.path)
		if err != nil {
			return nil, err
		}
		lr.ranger = rr
	}
	return lr.ranger.Range(ctx, offset, length)
}
