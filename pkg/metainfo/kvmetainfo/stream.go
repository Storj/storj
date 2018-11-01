// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package kvmetainfo

import (
	"context"
	"errors"

	"github.com/gogo/protobuf/proto"

	"storj.io/storj/pkg/encryption"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storage/streams"
	"storj.io/storj/pkg/storj"
)

type readonlyStream struct {
	db *Objects

	info          storj.Object
	encryptedPath storj.Path
	streamKey     *storj.Key // lazySegmentReader derivedKey

	lastSegment     pb.StreamMeta
	lastSegmentSize int64
}

func (stream *readonlyStream) Info() storj.Object { return stream.info }

func (stream *readonlyStream) SegmentsAt(ctx context.Context, byteOffset int64, limit int64) (infos []storj.Segment, more bool, err error) {
	if stream.info.FixedSegmentSize <= 0 {
		return nil, false, errors.New("not implemented")
	}

	index := byteOffset / stream.info.FixedSegmentSize
	return stream.Segments(ctx, index, limit)
}

func (stream *readonlyStream) segment(ctx context.Context, index int64) (storj.Segment, error) {
	segment := storj.Segment{
		Index: index,
	}

	isLastSegment := segment.Index+1 == stream.info.SegmentCount
	if !isLastSegment {
		segmentPath := streams.GetSegmentPath(stream.encryptedPath, index+1)
		_, meta, err := stream.db.segments.Get(ctx, segmentPath)
		if err != nil {
			return segment, err
		}

		segmentMeta := pb.SegmentMeta{}
		err = proto.Unmarshal(meta.Data, &segmentMeta)
		if err != nil {
			return segment, err
		}
		encryptedKey, encryptedKeyNonce := streams.GetEncryptedKeyAndNonce(&segmentMeta)

		segment.Size = stream.info.FixedSegmentSize
		segment.EncryptedKeyNonce = *encryptedKeyNonce
		segment.EncryptedKey = encryptedKey
	} else {
		lastEncryptedKey, lastEncryptedKeyNonce := streams.GetEncryptedKeyAndNonce(stream.lastSegment.LastSegmentMeta)

		segment.Size = stream.lastSegmentSize
		segment.EncryptedKeyNonce = *lastEncryptedKeyNonce
		segment.EncryptedKey = lastEncryptedKey
	}

	var nonce storj.Nonce
	_, err := encryption.Increment(&nonce, index+1)
	if err != nil {
		return segment, err
	}

	return segment, nil
}

func (stream *readonlyStream) Segments(ctx context.Context, index int64, limit int64) (infos []storj.Segment, more bool, err error) {
	if index < 0 {
		return nil, false, errors.New("invalid argument")
	}
	if limit <= 0 {
		limit = defaultSegmentLimit
	}
	if index >= stream.info.SegmentCount {
		return nil, false, nil
	}

	infos = make([]storj.Segment, 0, limit)
	for ; index < stream.info.SegmentCount && limit > 0; index++ {
		limit--
		segment, err := stream.segment(ctx, index)
		if err != nil {
			return nil, false, err
		}
		infos = append(infos, segment)
	}

	more = index+limit >= stream.info.SegmentCount
	return infos, more, nil
}
