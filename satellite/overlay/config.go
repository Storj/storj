// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package overlay

import (
	"time"

	"github.com/spacemonkeygo/monkit/v3"
	"github.com/zeebo/errs"

	"storj.io/common/memory"
)

var (
	mon = monkit.Package()
	// Error represents an overlay error
	Error = errs.Class("overlay error")
)

// Config is a configuration for overlay service.
type Config struct {
	Node                 NodeSelectionConfig
	UpdateStatsBatchSize int       `help:"number of update requests to process per transaction" default:"100"`
	QueueMaxSize         int       `help:"the max size of the select storage node queue" default:"1000"`
	QueueExpiredTime     time.Time `help:"how long until items in the select storaage node queue are stale" default:"60s"`
}

// NodeSelectionConfig is a configuration struct to determine the minimum
// values for nodes to select
type NodeSelectionConfig struct {
	UptimeCount      int64         `help:"the number of times a node's uptime has been checked to not be considered a New Node" releaseDefault:"100" devDefault:"0"`
	AuditCount       int64         `help:"the number of times a node has been audited to not be considered a New Node" releaseDefault:"100" devDefault:"0"`
	NewNodeFraction  float64       `help:"the fraction of new nodes allowed per request" default:"0.05"`
	MinimumVersion   string        `help:"the minimum node software version for node selection queries" default:""`
	OnlineWindow     time.Duration `help:"the amount of time without seeing a node before its considered offline" default:"4h"`
	DistinctIP       bool          `help:"require distinct IPs when choosing nodes for upload" releaseDefault:"true" devDefault:"false"`
	MinimumDiskSpace memory.Size   `help:"how much disk space a node at minimum must have to be selected for upload" default:"100MB"`

	AuditReputationRepairWeight float64 `help:"weight to apply to audit reputation for total repair reputation calculation" default:"1.0"`
	AuditReputationUplinkWeight float64 `help:"weight to apply to audit reputation for total uplink reputation calculation" default:"1.0"`
	AuditReputationAlpha0       float64 `help:"the initial shape 'alpha' used to calculate audit SNs reputation" default:"1.0"`
	AuditReputationBeta0        float64 `help:"the initial shape 'beta' value used to calculate audit SNs reputation" default:"0.0"`
	AuditReputationLambda       float64 `help:"the forgetting factor used to calculate the audit SNs reputation" default:"0.95"`
	AuditReputationWeight       float64 `help:"the normalization weight used to calculate the audit SNs reputation" default:"1.0"`
	AuditReputationDQ           float64 `help:"the reputation cut-off for disqualifying SNs based on audit history" default:"0.6"`
}
