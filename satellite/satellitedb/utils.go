// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"database/sql/driver"

	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/zeebo/errs"

	"storj.io/storj/pkg/storj"
)

// bytesToUUID is used to convert []byte to UUID
func bytesToUUID(data []byte) (uuid.UUID, error) {
	var id uuid.UUID

	copy(id[:], data)
	if len(id) != len(data) {
		return uuid.UUID{}, errs.New("Invalid uuid")
	}

	return id, nil
}

type postgresNodeIDList storj.NodeIDList

// Value converts a NodeIDList to a postgres array
func (nodes postgresNodeIDList) Value() (driver.Value, error) {
	const hextable = "0123456789abcdef"

	if nodes == nil {
		return nil, nil
	}
	if len(nodes) == 0 {
		return []byte("{}"), nil
	}

	var wp, x int
	out := make([]byte, 2+len(nodes)*(6+storj.NodeIDSize*2)-1)

	x = copy(out[wp:], []byte(`{"\\x`))
	wp += x

	for i := range nodes {
		for _, v := range nodes[i] {
			out[wp] = hextable[v>>4]
			out[wp+1] = hextable[v&0xf]
			wp += 2
		}

		if i+1 < len(nodes) {
			x = copy(out[wp:], []byte(`","\\x`))
			wp += x
		}
	}

	x = copy(out[wp:], `"}`)
	wp += x

	if wp != len(out) {
		panic("unreachable")
	}

	return out, nil
}

type postgresArrayList [][]byte

// Value converts an array of byte arrays to a postgres array
func (values postgresArrayList) Value() (driver.Value, error) {
	const hextable = "0123456789abcdef"

	if values == nil {
		return nil, nil
	}
	if len(values) == 0 {
		return []byte("{}"), nil
	}

	var total int
	for _, path := range values {
		total += len(path)
	}
	var wp, x int
	out := make([]byte, 2+(6*len(values))+(total*2)-1)

	x = copy(out[wp:], []byte(`{"\\x`))
	wp += x

	for i := range values {
		for _, v := range values[i] {
			out[wp] = hextable[v>>4]
			out[wp+1] = hextable[v&0xf]
			wp += 2
		}

		if i+1 < len(values) {
			x = copy(out[wp:], []byte(`","\\x`))
			wp += x
		}
	}

	x = copy(out[wp:], `"}`)
	wp += x

	if wp != len(out) {
		panic("unreachable")
	}

	return out, nil
}

// uuidScan used to represent uuid scan struct
type uuidScan struct {
	uuid *uuid.UUID
}

// Scan is used to wrap logic of db scan with uuid conversion
func (s *uuidScan) Scan(src interface{}) (err error) {
	b, ok := src.([]byte)
	if !ok {
		return Error.New("unexpected type %T for uuid", src)
	}

	*s.uuid, err = bytesToUUID(b)
	if err != nil {
		return Error.Wrap(err)
	}

	return nil
}
