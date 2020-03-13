// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package date_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"storj.io/storj/private/date"
)

func TestMonthBoundary(t *testing.T) {
	now := time.Now()

	start, end := date.MonthBoundary(now)
	assert.Equal(t, start, time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()))
	assert.Equal(t, end, time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, -1, now.Location()))
}

func TestDayBoundary(t *testing.T) {
	now := time.Now()

	start, end := date.DayBoundary(now)
	assert.Equal(t, start, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
	assert.Equal(t, end, time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, -1, now.Location()))
}
