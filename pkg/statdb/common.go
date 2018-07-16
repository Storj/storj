// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package statdb

import (
	"github.com/zeebo/errs"
)

// Error is the default boltdb errs class
var Error = errs.Class("statdb error")
