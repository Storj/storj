// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package console_test

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/satellite"
	"storj.io/storj/satellite/console"
	"storj.io/storj/satellite/satellitedb/satellitedbtest"
)

func TestProjectInvoiceStamps(t *testing.T) {
	satellitedbtest.Run(t, func(t *testing.T, db satellite.DB) {
		ctx := testcontext.New(t)
		consoleDB := db.Console()

		startDate := time.Now().UTC()
		endDate := startDate.Add(time.Hour * 24)

		var invoiceID [8]byte
		_, err := rand.Read(invoiceID[:])
		if err != nil {
			t.Fatal(fmt.Sprintf("can not create invoice id: %s", err))
		}

		//create project
		proj, err := consoleDB.Projects().Insert(ctx, &console.Project{
			Name: "test",
		})
		if err != nil {
			t.Fatal(fmt.Sprintf("can not create project: %s", err))
		}

		t.Run("create project invoice stamp", func(t *testing.T) {
			stamp, err := consoleDB.ProjectInvoiceStamps().Create(ctx, console.ProjectInvoiceStamp{
				ProjectID: proj.ID,
				InvoiceID: invoiceID[:],
				StartDate: startDate,
				EndDate:   endDate,
			})

			assert.NoError(t, err)
			assert.Equal(t, proj.ID, stamp.ProjectID)
			assert.Equal(t, invoiceID[:], stamp.InvoiceID)
			assert.Equal(t, startDate, stamp.StartDate)
			assert.Equal(t, endDate, stamp.EndDate)
		})

		t.Run("get by project id and start date", func(t *testing.T) {
			stamp, err := consoleDB.ProjectInvoiceStamps().GetByProjectIDStartDate(ctx, proj.ID, startDate)

			assert.NoError(t, err)
			assert.Equal(t, proj.ID, stamp.ProjectID)
			assert.Equal(t, invoiceID[:], stamp.InvoiceID)
			assert.Equal(t, startDate, stamp.StartDate)
			assert.Equal(t, endDate, stamp.EndDate)
		})
	})
}
