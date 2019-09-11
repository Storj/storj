// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package storagenodedbtest

var v5 = Snapshots.Add(&MultiDBSnapshot{
	Version:   5,
	Databases: v4.Databases,
})
