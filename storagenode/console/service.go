// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package console

import (
	"context"
	"time"

	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/internal/date"
	"storj.io/storj/internal/memory"
	"storj.io/storj/internal/version"
	"storj.io/storj/pkg/kademlia"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/storagenode/bandwidth"
	"storj.io/storj/storagenode/pieces"
	"storj.io/storj/storagenode/reputation"
	"storj.io/storj/storagenode/storageusage"
	"storj.io/storj/storagenode/trust"
)

var (
	// SNOServiceErr defines sno service error.
	SNOServiceErr = errs.Class("storage node dashboard service error")

	mon = monkit.Package()
)

// DB exposes methods for managing SNO dashboard related data.
type DB interface {
	// Bandwidth is a getter for Bandwidth db.
	Bandwidth() Bandwidth
}

// Service is handling storage node operator related logic.
type Service struct {
	log *zap.Logger

	trust          *trust.Pool
	consoleDB      DB
	bandwidthDB    bandwidth.DB
	reputationDB   reputation.DB
	storageUsageDB storageusage.DB
	pieceStore     *pieces.Store
	kademlia       *kademlia.Kademlia
	version        *version.Service

	allocatedBandwidth memory.Size
	allocatedDiskSpace memory.Size
	walletAddress      string
	startedAt          time.Time
	versionInfo        version.Info
}

// NewService returns new instance of Service.
func NewService(log *zap.Logger, consoleDB DB, bandwidth bandwidth.DB, pieceStore *pieces.Store, kademlia *kademlia.Kademlia, version *version.Service,
	allocatedBandwidth, allocatedDiskSpace memory.Size, walletAddress string, versionInfo version.Info, trust *trust.Pool,
	reputationDB reputation.DB, storageUsageDB storageusage.DB) (*Service, error) {
	if log == nil {
		return nil, errs.New("log can't be nil")
	}

	if consoleDB == nil {
		return nil, errs.New("consoleDB can't be nil")
	}

	if bandwidth == nil {
		return nil, errs.New("bandwidth can't be nil")
	}

	if pieceStore == nil {
		return nil, errs.New("pieceStore can't be nil")
	}

	if version == nil {
		return nil, errs.New("version can't be nil")
	}

	if kademlia == nil {
		return nil, errs.New("kademlia can't be nil")
	}

	return &Service{
		log:                log,
		trust:              trust,
		consoleDB:          consoleDB,
		bandwidthDB:        bandwidth,
		reputationDB:       reputationDB,
		storageUsageDB:     storageUsageDB,
		pieceStore:         pieceStore,
		kademlia:           kademlia,
		version:            version,
		allocatedBandwidth: allocatedBandwidth,
		allocatedDiskSpace: allocatedDiskSpace,
		walletAddress:      walletAddress,
		startedAt:          time.Now(),
		versionInfo:        versionInfo,
	}, nil
}

// SatelliteInfo encapsulates satellite ID and disqualification
type SatelliteInfo struct {
	ID           storj.NodeID `json:"id"`
	Disqualified *time.Time   `json:"disqualified"`
}

// Dashboard encapsulates dashboard stale data.
type Dashboard struct {
	NodeID storj.NodeID `json:"nodeID"`
	Wallet string       `json:"wallet"`

	Satellites []SatelliteInfo `json:"satellites"`

	DiskSpace DiskSpaceInfo `json:"diskSpace"`
	Bandwidth BandwidthInfo `json:"bandwidth"`

	LastPinged  time.Time `json:"lastPinged"`
	LastQueried time.Time `json:"lastQueried"`

	Version  version.SemVer `json:"version"`
	UpToDate bool           `json:"upToDate"`
}

// GetDashboardData returns stale dashboard data.
func (s *Service) GetDashboardData(ctx context.Context) (_ *Dashboard, err error) {
	defer mon.Task()(&ctx)(&err)
	data := new(Dashboard)

	data.NodeID = s.kademlia.Local().Id
	data.Wallet = s.walletAddress
	data.Version = s.versionInfo.Version
	data.UpToDate = s.version.IsAllowed()
	data.LastPinged = s.kademlia.LastPinged()
	data.LastQueried = s.kademlia.LastQueried()

	stats, err := s.reputationDB.All(ctx)
	if err != nil {
		return nil, SNOServiceErr.Wrap(err)
	}

	for _, rep := range stats {
		data.Satellites = append(data.Satellites,
			SatelliteInfo{
				ID:           rep.SatelliteID,
				Disqualified: rep.Disqualified,
			},
		)
	}

	spaceUsage, err := s.pieceStore.SpaceUsedForPieces(ctx)
	if err != nil {
		return nil, SNOServiceErr.Wrap(err)
	}

	bandwidthUsage, err := bandwidth.TotalMonthlySummary(ctx, s.bandwidthDB)
	if err != nil {
		return nil, SNOServiceErr.Wrap(err)
	}

	data.DiskSpace = DiskSpaceInfo{
		Used:      memory.Size(spaceUsage).GB(),
		Available: s.allocatedDiskSpace.GB(),
	}

	data.Bandwidth = BandwidthInfo{
		Used:      memory.Size(bandwidthUsage.Total()).GB(),
		Available: s.allocatedBandwidth.GB(),
	}

	return data, nil
}

// Satellite encapsulates satellite related data.
type Satellite struct {
	ID               storj.NodeID         `json:"id"`
	StorageDaily     []storageusage.Stamp `json:"storageDaily"`
	BandwidthDaily   []BandwidthUsed      `json:"bandwidthDaily"`
	StorageSummary   float64              `json:"storageSummary"`
	BandwidthSummary int64                `json:"bandwidthSummary"`
	Audit            reputation.Metric    `json:"audit"`
	Uptime           reputation.Metric    `json:"uptime"`
}

// GetSatelliteData returns satellite related data.
func (s *Service) GetSatelliteData(ctx context.Context, satelliteID storj.NodeID) (_ *Satellite, err error) {
	defer mon.Task()(&ctx)(&err)
	from, to := date.MonthBoundary(time.Now().UTC())

	bandwidthDaily, err := s.consoleDB.Bandwidth().GetDaily(ctx, satelliteID, from, to)
	if err != nil {
		return nil, SNOServiceErr.Wrap(err)
	}

	storageDaily, err := s.storageUsageDB.GetDaily(ctx, satelliteID, from, to)
	if err != nil {
		return nil, SNOServiceErr.Wrap(err)
	}

	bandwidthSummary, err := s.bandwidthDB.SatelliteSummary(ctx, satelliteID, from, to)
	if err != nil {
		return nil, SNOServiceErr.Wrap(err)
	}

	storageSummary, err := s.storageUsageDB.SatelliteSummary(ctx, satelliteID, from, to)
	if err != nil {
		return nil, SNOServiceErr.Wrap(err)
	}

	rep, err := s.reputationDB.Get(ctx, satelliteID)
	if err != nil {
		return nil, SNOServiceErr.Wrap(err)
	}

	return &Satellite{
		ID:               satelliteID,
		StorageDaily:     storageDaily,
		BandwidthDaily:   bandwidthDaily,
		StorageSummary:   storageSummary,
		BandwidthSummary: bandwidthSummary.Total(),
		Audit:            rep.Audit,
		Uptime:           rep.Uptime,
	}, nil
}

// Satellites represents consolidated data across all satellites.
type Satellites struct {
	StorageDaily     []storageusage.Stamp `json:"storageDaily"`
	BandwidthDaily   []BandwidthUsed      `json:"bandwidthDaily"`
	StorageSummary   float64              `json:"storageSummary"`
	BandwidthSummary int64                `json:"bandwidthSummary"`
}

// GetAllSatellitesData returns bandwidth and storage daily usage consolidate
// among all satellites from the node's trust pool.
func (s *Service) GetAllSatellitesData(ctx context.Context) (_ *Satellites, err error) {
	defer mon.Task()(&ctx)(nil)
	from, to := date.MonthBoundary(time.Now().UTC())

	bandwidthDaily, err := s.consoleDB.Bandwidth().GetDailyTotal(ctx, from, to)
	if err != nil {
		return nil, SNOServiceErr.Wrap(err)
	}

	storageDaily, err := s.storageUsageDB.GetDailyTotal(ctx, from, to)
	if err != nil {
		return nil, SNOServiceErr.Wrap(err)
	}

	bandwidthSummary, err := s.bandwidthDB.Summary(ctx, from, to)
	if err != nil {
		return nil, SNOServiceErr.Wrap(err)
	}

	storageSummary, err := s.storageUsageDB.Summary(ctx, from, to)
	if err != nil {
		return nil, SNOServiceErr.Wrap(err)
	}

	return &Satellites{
		StorageDaily:     storageDaily,
		BandwidthDaily:   bandwidthDaily,
		StorageSummary:   storageSummary,
		BandwidthSummary: bandwidthSummary.Total(),
	}, nil
}

// VerifySatelliteID verifies if the satellite belongs to the trust pool.
func (s *Service) VerifySatelliteID(ctx context.Context, satelliteID storj.NodeID) (err error) {
	defer mon.Task()(&ctx)(&err)

	err = s.trust.VerifySatelliteID(ctx, satelliteID)
	if err != nil {
		return SNOServiceErr.Wrap(err)
	}

	return nil
}
