// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package marketingweb_test

import (
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/satellite/marketing"
)

type CreateRequest struct {
	Path   string
	Values url.Values
}

func TestCreateOffer(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {

		requests := []CreateRequest{
			{
				Path: "/create/referral-offer",
				Values: url.Values{
					"Name":                      {"Referral Credit"},
					"Description":               {"desc"},
					"ExpiresAt":                 {"2019-06-27"},
					"InviteeCreditInCents":      {"50"},
					"InviteeCreditDurationDays": {"50"},
					"AwardCreditInCents":        {"50"},
					"AwardCreditDurationDays":   {"50"},
					"RedeemableCap":             {"150"},
				},
			}, {
				Path: "/create/free-credit-offer",
				Values: url.Values{
					"Name":                      {"Free Credit Credit"},
					"Description":               {"desc"},
					"ExpiresAt":                 {"2019-06-27"},
					"InviteeCreditInCents":      {"50"},
					"InviteeCreditDurationDays": {"50"},
					"RedeemableCap":             {"150"},
				},
			},
		}

		addr := planet.Satellites[0].Marketing.Listener.Addr()

		var group errgroup.Group
		for _, offer := range requests {
			o := offer
			group.Go(func() error {
				baseURL := "http://" + addr.String()

				_, err := http.PostForm(baseURL+o.Path, o.Values)
				if err != nil {
					return err
				}

				_, err = http.Get(baseURL)
				if err != nil {
					return err
				}

				return nil
			})
			err := group.Wait()
			require.NoError(t, err)
		}
	})
}

func TestStopOffer(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {

		offersToStop := []marketing.Offer{
			{
				ID:   1,
				Type: marketing.Referral,
			}, {
				ID:   1,
				Type: marketing.FreeCredit,
			},
		}

		addr := planet.Satellites[0].Marketing.Listener.Addr()

		var group errgroup.Group
		for _, offer := range offersToStop {
			o := offer
			group.Go(func() error {
				baseURL := "http://" + addr.String()
				endpoint := "/stop/"
				param := strconv.Itoa(o.ID)

				_, err := http.Get(baseURL + endpoint + param)
				if err != nil {
					return err
				}

				_, err = http.Get(baseURL)
				if err != nil {
					return err
				}

				return nil
			})
			err := group.Wait()
			require.NoError(t, err)
		}
	})
}
