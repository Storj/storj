// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package trust

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/zeebo/errs"

	"storj.io/storj/pkg/storj"
)

// SatelliteURL represents a Satellite URL
type SatelliteURL struct {
	ID   storj.NodeID `json:"id"`
	Host string       `json:"host"`
	Port int          `json:"port"`
}

// Address returns the address (i.e. host:port) of the Satellite
func (u *SatelliteURL) Address() string {
	return fmt.Sprintf("%s:%d", u.Host, u.Port)
}

// NodeURL returns a full Node URL to the Satellite
func (u *SatelliteURL) NodeURL() storj.NodeURL {
	return storj.NodeURL{
		ID:      u.ID,
		Address: u.Address(),
	}
}

// String returns a string representation of the Satellite URL
func (u *SatelliteURL) String() string {
	return fmt.Sprintf("%s@%s:%d", u.ID.String(), u.Host, u.Port)
}

// ParseSatelliteURL parses a Satellite URL. For the purposes of the trust list,
// the Satellite URL MUST contain both an ID and port designation.
func ParseSatelliteURL(s string) (SatelliteURL, error) {
	url, err := storj.ParseNodeURL(s)
	if err != nil {
		return SatelliteURL{}, Error.New("invalid satellite URL: %v", err)
	}
	if url.ID.IsZero() {
		return SatelliteURL{}, Error.New("invalid satellite URL: must contain an ID")
	}

	// storj.ParseNodeURL will have already verified that the address is
	// well-formed, so if SplitHostPort fails it should be due to the address
	// not having a port
	host, portStr, err := net.SplitHostPort(url.Address)
	if err != nil {
		return SatelliteURL{}, Error.New("invalid satellite URL: must specify the port")
	}

	// Port should already be numeric so this shouldn't fail, but just in case.
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return SatelliteURL{}, Error.New("invalid satellite URL: port is not numeric")
	}

	return SatelliteURL{
		ID:   url.ID,
		Host: host,
		Port: port,
	}, nil
}

// ParseSatelliteURLList parses a newline separated list of Satellite URLs.
// Empty lines or lines starting with '#' (comments) are ignored.
func ParseSatelliteURLList(ctx context.Context, r io.Reader) (urls []SatelliteURL, err error) {
	defer mon.Task()(&ctx)(&err)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}

		url, err := ParseSatelliteURL(line)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	if err := scanner.Err(); err != nil {
		return nil, Error.Wrap(err)
	}

	return urls, nil
}

// LoadSatelliteURLList loads a list of Satellite URLs from a path on disk
func LoadSatelliteURLList(ctx context.Context, path string) (_ []SatelliteURL, err error) {
	defer mon.Task()(&ctx)(&err)

	f, err := os.Open(path)
	if err != nil {
		return nil, Error.Wrap(err)
	}
	defer func() { err = errs.Combine(err, f.Close()) }()

	return ParseSatelliteURLList(ctx, f)
}
