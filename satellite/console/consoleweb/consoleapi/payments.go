// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package consoleapi

import (
	"context"
	"encoding/json"
	"github.com/zeebo/errs"
	"io/ioutil"
	"net/http"
	"strings"

	"go.uber.org/zap"
	"gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/auth"
	"storj.io/storj/satellite/console"
)

var (
	Error = errs.Class("satellite console payments api error")
	mon = monkit.Package()
)


// Payments is an api controller that exposes all payment related functionality
type Payments struct {
	log     *zap.Logger
	service *console.Service
}

// NewPayments is a constructor for api payments controller.
func NewPayments(log *zap.Logger, service *console.Service) *Payments {
	return &Payments{
		log:     log,
		service: service,
	}
}

// SetupAccount creates a payment account for the user.
func (p *Payments) SetupAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ctx = p.authorize(ctx, r)

	err = p.service.Payments().SetupAccount(ctx)
	if err != nil {
		if console.ErrUnauthorized.Has(err) {
			p.serveJSONError(w, http.StatusUnauthorized, err)
			return
		}

		p.serveJSONError(w, http.StatusInternalServerError, err)
		return
	}
}

// AccountBalance returns an integer amount in cents that represents the current balance of payment account.
func (p *Payments) AccountBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ctx = p.authorize(ctx, r)

	balance, err := p.service.Payments().AccountBalance(ctx)
	if err != nil {
		if console.ErrUnauthorized.Has(err) {
			p.serveJSONError(w, http.StatusUnauthorized, err)
			return
		}

		p.serveJSONError(w, http.StatusInternalServerError, err)
		return
	}

	err = json.NewEncoder(w).Encode(balance)
	if err != nil {
		p.log.Error("failed to write json balance response", zap.Error(Error.Wrap(err)))
	}
}

// CreditCards selects what to do depends on http method type.
func (p *Payments) CreditCards(w http.ResponseWriter, r *http.Request) {
	p.log.Error("AAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	switch r.Method {
	case http.MethodPatch:
		p.MakeCreditCardDefault(w, r)
	case http.MethodGet:
		p.ListCreditCards(w, r)
	case http.MethodPost:
		p.AddCreditCard(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// AddCreditCard is used to save new credit card and attach it to payment account.
func (p *Payments) AddCreditCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	ctx = p.authorize(ctx, r)

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		p.serveJSONError(w, http.StatusBadRequest, err)
		return
	}

	token := string(bodyBytes)

	err = p.service.Payments().AddCreditCard(ctx, token)
	if err != nil {
		if console.ErrUnauthorized.Has(err) {
			p.serveJSONError(w, http.StatusUnauthorized, err)
			return
		}

		p.serveJSONError(w, http.StatusInternalServerError, err)
		return
	}
}

// ListCreditCards returns a list of credit cards for a given payment account.
func (p *Payments) ListCreditCards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	ctx = p.authorize(ctx, r)

	cards, err := p.service.Payments().ListCreditCards(ctx)
	if err != nil {
		if console.ErrUnauthorized.Has(err) {
			p.serveJSONError(w, http.StatusUnauthorized, err)
			return
		}

		p.serveJSONError(w, http.StatusInternalServerError, err)
		return
	}

	err = json.NewEncoder(w).Encode(cards)
	if err != nil {
		p.log.Error("failed to write json list cards response", zap.Error(Error.Wrap(err)))
	}
}

// MakeCreditCardDefault makes a credit card default payment method.
func (p *Payments) MakeCreditCardDefault(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	ctx = p.authorize(ctx, r)

	cardID, err := ioutil.ReadAll(r.Body)
	if err != nil {
		p.serveJSONError(w, http.StatusBadRequest, err)
		return
	}

	err = p.service.Payments().MakeCreditCardDefault(ctx, cardID)
	if err != nil {
		if console.ErrUnauthorized.Has(err) {
			p.serveJSONError(w, http.StatusUnauthorized, err)
			return
		}

		p.serveJSONError(w, http.StatusInternalServerError, err)
		return
	}
}

// RemoveCreditCard is used to detach a credit card from payment account.
func (p *Payments) RemoveCreditCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ctx = p.authorize(ctx, r)

	p.log.Error(r.URL.Path[len("/api/v0/payments/cards/"):])

	cardID := r.URL.Path[len("/api/v0/payments/cards/"):]

	err = p.service.Payments().RemoveCreditCard(ctx, []byte(cardID))
	if err != nil {
		if console.ErrUnauthorized.Has(err) {
			p.serveJSONError(w, http.StatusUnauthorized, err)
			return
		}

		p.serveJSONError(w, http.StatusInternalServerError, err)
		return
	}
}

// serveJSONError writes JSON error to response output stream.
func (p *Payments) serveJSONError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		p.log.Error("failed to write json error response", zap.Error(Error.Wrap(err)))
	}
}

// authorize checks request for authorization token, validates it and updates context with auth data.
func (p *Payments) authorize(ctx context.Context, r *http.Request) context.Context {
	authHeaderValue := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeaderValue, "Bearer ")

	auth, err := p.service.Authorize(auth.WithAPIKey(ctx, []byte(token)))
	if err != nil {
		return console.WithAuthFailure(ctx, err)
	}

	return console.WithAuth(ctx, auth)
}
