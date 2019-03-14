// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package repairer

import (
	"context"
	"time"

	"storj.io/storj/internal/memory"
	"storj.io/storj/pkg/auth/signing"
	"storj.io/storj/pkg/overlay"
	"storj.io/storj/pkg/pointerdb"
	ecclient "storj.io/storj/pkg/storage/ec"
	"storj.io/storj/pkg/storage/segments"
	"storj.io/storj/pkg/transport"
)

// Config contains configurable values for repairer
type Config struct {
	MaxRepair    int           `help:"maximum segments that can be repaired concurrently" default:"100"`
	Interval     time.Duration `help:"how frequently checker should audit segments" default:"3600s"`
	MaxBufferMem memory.Size   `help:"maximum buffer memory (in bytes) to be allocated for read buffers" default:"4M"`
}

// GetSegmentRepairer creates a new segment repairer from storeConfig values
func (c Config) GetSegmentRepairer(ctx context.Context, tc transport.Client, pointerdb *pointerdb.Service, allocation *pointerdb.AllocationSigner, cache *overlay.Cache, signer signing.Signer, selectionPreferences *overlay.NodeSelectionConfig) (ss SegmentRepairer, err error) {
	defer mon.Task()(&ctx)(&err)

	ec := ecclient.NewClient(tc, c.MaxBufferMem.Int())

	return segments.NewSegmentRepairer(pointerdb, allocation, cache, ec, signer, selectionPreferences), nil
}
