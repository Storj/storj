// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package consoleapi

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeebo/errs"
	"go.uber.org/zap"

	"storj.io/storj/pkg/storj"
	"storj.io/storj/storagenode/heldamount"
)

// ErrHeldAmountPI - console heldAmount api error type.
var ErrHeldAmountPI = errs.Class("heldAmount console web error")

// HeldAmount is an api controller that exposes all held amount related api.
type HeldAmount struct {
	service *heldamount.Service

	log *zap.Logger
}

// NewHeldAmount is a constructor for heldAmount controller.
func NewHeldAmount(log *zap.Logger, service *heldamount.Service) *HeldAmount {
	return &HeldAmount{
		log:     log,
		service: service,
	}
}

// SatellitePayStubMonthly returns heldamount, storage holding and prices data for specific month from satellite.
func (heldAmount *HeldAmount) SatellitePayStubMonthly(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	w.Header().Set(contentType, applicationJSON)

	params := mux.Vars(r)

	period, ok := params["period"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}

	id, ok := params["satelliteID"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}
	satelliteID, err := storj.NodeIDFromString(id)
	if err != nil {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrHeldAmountPI.Wrap(err))
		return
	}

	payStub, err := heldAmount.service.SatellitePayStubMonthlyCached(ctx, satelliteID, period)
	if err != nil {
		heldAmount.serveJSONError(w, http.StatusInternalServerError, ErrHeldAmountPI.Wrap(err))
		return
	}

	if err := json.NewEncoder(w).Encode(payStub); err != nil {
		heldAmount.log.Error("failed to encode json response", zap.Error(ErrHeldAmountPI.Wrap(err)))
		return
	}
}

// AllPayStubsMonthly returns heldAmount, storage holding and prices data for specific month from all satellites.
func (heldAmount *HeldAmount) AllPayStubsMonthly(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	w.Header().Set(contentType, applicationJSON)

	params := mux.Vars(r)

	period, ok := params["period"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}

	payStubs, err := heldAmount.service.AllPayStubsMonthlyCached(ctx, period)
	if err != nil {
		heldAmount.serveJSONError(w, http.StatusInternalServerError, ErrHeldAmountPI.Wrap(err))
		return
	}

	if err := json.NewEncoder(w).Encode(payStubs); err != nil {
		heldAmount.log.Error("failed to encode json response", zap.Error(ErrHeldAmountPI.Wrap(err)))
		return
	}
}

// SatellitePayStubPeriod retrieves held amount for all satellites for selected months from storagenode database.
func (heldAmount *HeldAmount) SatellitePayStubPeriod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	w.Header().Set(contentType, applicationJSON)

	params := mux.Vars(r)

	id, ok := params["satelliteID"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}
	satelliteID, err := storj.NodeIDFromString(id)
	if err != nil {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrHeldAmountPI.Wrap(err))
		return
	}

	start, ok := params["start"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}

	end, ok := params["end"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}

	payStubs, err := heldAmount.service.SatellitePayStubPeriodCached(ctx, satelliteID, start, end)
	if err != nil {
		heldAmount.serveJSONError(w, http.StatusInternalServerError, ErrHeldAmountPI.Wrap(err))
		return
	}

	if err := json.NewEncoder(w).Encode(payStubs); err != nil {
		heldAmount.log.Error("failed to encode json response", zap.Error(ErrHeldAmountPI.Wrap(err)))
		return
	}
}

// AllPayStubsPeriod retrieves held amount for all satellites for selected range of months from storagenode database.
func (heldAmount *HeldAmount) AllPayStubsPeriod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	w.Header().Set(contentType, applicationJSON)

	params := mux.Vars(r)

	start, ok := params["start"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}

	end, ok := params["end"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}

	payStubs, err := heldAmount.service.AllPayStubsPeriodCached(ctx, start, end)
	if err != nil {
		heldAmount.serveJSONError(w, http.StatusInternalServerError, ErrHeldAmountPI.Wrap(err))
		return
	}

	if err := json.NewEncoder(w).Encode(payStubs); err != nil {
		heldAmount.log.Error("failed to encode json response", zap.Error(ErrHeldAmountPI.Wrap(err)))
		return
	}
}

// SatellitePaymentMonthly returns payment data from satellite for specific month.
func (heldAmount *HeldAmount) SatellitePaymentMonthly(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	w.Header().Set(contentType, applicationJSON)

	params := mux.Vars(r)

	period, ok := params["period"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}

	id, ok := params["satelliteID"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}
	satelliteID, err := storj.NodeIDFromString(id)
	if err != nil {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrHeldAmountPI.Wrap(err))
		return
	}

	paymentData, err := heldAmount.service.SatellitePaymentMonthlyCached(ctx, satelliteID, period)
	if err != nil {
		heldAmount.serveJSONError(w, http.StatusInternalServerError, ErrHeldAmountPI.Wrap(err))
		return
	}

	if err := json.NewEncoder(w).Encode(paymentData); err != nil {
		heldAmount.log.Error("failed to encode json response", zap.Error(ErrHeldAmountPI.Wrap(err)))
		return
	}
}

// AllPaymentsMonthly returns payments for specific month from all satellites.
func (heldAmount *HeldAmount) AllPaymentsMonthly(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	w.Header().Set(contentType, applicationJSON)

	params := mux.Vars(r)

	period, ok := params["period"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}

	payStubs, err := heldAmount.service.AllPaymentsMonthlyCached(ctx, period)
	if err != nil {
		heldAmount.serveJSONError(w, http.StatusInternalServerError, ErrHeldAmountPI.Wrap(err))
		return
	}

	if err := json.NewEncoder(w).Encode(payStubs); err != nil {
		heldAmount.log.Error("failed to encode json response", zap.Error(ErrHeldAmountPI.Wrap(err)))
		return
	}
}

// SatellitePaymentPeriod retrieves payment for selected satellite for selected period from storagenode database.
func (heldAmount *HeldAmount) SatellitePaymentPeriod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	w.Header().Set(contentType, applicationJSON)

	params := mux.Vars(r)

	id, ok := params["satelliteID"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}
	satelliteID, err := storj.NodeIDFromString(id)
	if err != nil {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrHeldAmountPI.Wrap(err))
		return
	}

	start, ok := params["start"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}

	end, ok := params["end"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}

	payments, err := heldAmount.service.SatellitePaymentPeriodCached(ctx, satelliteID, start, end)
	if err != nil {
		heldAmount.serveJSONError(w, http.StatusInternalServerError, ErrHeldAmountPI.Wrap(err))
		return
	}

	if err := json.NewEncoder(w).Encode(payments); err != nil {
		heldAmount.log.Error("failed to encode json response", zap.Error(ErrHeldAmountPI.Wrap(err)))
		return
	}
}

// AllPaymentsPeriod retrieves payment for all satellites for selected range of months from storagenode database.
func (heldAmount *HeldAmount) AllPaymentsPeriod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	w.Header().Set(contentType, applicationJSON)

	params := mux.Vars(r)

	start, ok := params["start"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}

	end, ok := params["end"]
	if !ok {
		heldAmount.serveJSONError(w, http.StatusBadRequest, ErrNotificationsAPI.Wrap(err))
		return
	}

	payStubs, err := heldAmount.service.AllPaymentsPeriodCached(ctx, start, end)
	if err != nil {
		heldAmount.serveJSONError(w, http.StatusInternalServerError, ErrHeldAmountPI.Wrap(err))
		return
	}

	if err := json.NewEncoder(w).Encode(payStubs); err != nil {
		heldAmount.log.Error("failed to encode json response", zap.Error(ErrHeldAmountPI.Wrap(err)))
		return
	}
}

// serveJSONError writes JSON error to response output stream.
func (heldAmount *HeldAmount) serveJSONError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		heldAmount.log.Error("failed to write json error response", zap.Error(ErrHeldAmountPI.Wrap(err)))
		return
	}
}
