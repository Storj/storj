// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package audit

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/vivint/infectious"

	"storj.io/storj/pkg/eestream"
	"storj.io/storj/pkg/paths"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/pointerdb/pdbclient"
	"storj.io/storj/pkg/storage/meta"
)

// Stripe is a struct that contains stripe info
type Stripe struct {
	Index int
}

// NextStripe returns a random stripe to be audited
func (auditor *Auditor) NextStripe(ctx context.Context) (stripe *Stripe, pointer *pb.Pointer, more bool, err error) {
	auditor.mutex.Lock()
	defer auditor.mutex.Unlock()

	var pointerItems []pdbclient.ListItem
	var path paths.Path

	if auditor.lastPath == nil {
		pointerItems, more, err = auditor.pointers.List(ctx, nil, nil, nil, true, 0, meta.None)
	} else {
		pointerItems, more, err = auditor.pointers.List(ctx, nil, *auditor.lastPath, nil, true, 0, meta.None)
	}

	if err != nil {
		return nil, nil, more, err
	}

	// get random pointer
	pointerItem, err := getRandomPointer(pointerItems)
	if err != nil {
		return nil, nil, more, err
	}

	// get path
	path = pointerItem.Path

	// keep track of last path listed
	if !more {
		auditor.lastPath = nil
	} else {
		auditor.lastPath = &pointerItems[len(pointerItems)-1].Path
	}

	// get pointer info
	pointer, err = auditor.pointers.Get(ctx, path)
	if err != nil {
		return nil, nil, more, err
	}

	// create the erasure scheme so we can get the stripe size
	es, err := makeErasureScheme(pointer.GetRemote().GetRedundancy())
	if err != nil {
		return nil, nil, more, err
	}

	//get random stripe
	index, err := getRandomStripe(es, pointer)
	if err != nil {
		return nil, nil, more, err
	}

	return &Stripe{Index: index}, pointer, more, nil
}

// create the erasure scheme
func makeErasureScheme(rs *pb.RedundancyScheme) (eestream.ErasureScheme, error) {
	fc, err := infectious.NewFEC(int(rs.GetMinReq()), int(rs.GetTotal()))
	if err != nil {
		return nil, err
	}
	es := eestream.NewRSScheme(fc, int(rs.GetErasureShareSize()))
	return es, nil
}

func getRandomStripe(es eestream.ErasureScheme, pointer *pb.Pointer) (index int, err error) {
	stripeSize := es.StripeSize()
	randomStripeIndex, err := rand.Int(rand.Reader, big.NewInt(pointer.GetSize()/int64(stripeSize)))
	if err != nil {
		return -1, err
	}
	return int(randomStripeIndex.Int64()), nil
}

func getRandomPointer(pointerItems []pdbclient.ListItem) (pointer pdbclient.ListItem, err error) {
	randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(len(pointerItems))))
	if err != nil {
		return pdbclient.ListItem{}, err
	}
	randomNumInt64 := randomNum.Int64()
	pointerItem := pointerItems[randomNumInt64]
	return pointerItem, nil
}
