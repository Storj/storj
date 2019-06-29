// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information

package rewards

import (
	"context"
	"fmt"
	"time"
)

// ToCents converts USD credit amounts to cents.
func ToCents(dollars int) int {
	return dollars * 100
}

// ToDollars converts credit amounts in cents to USD.
func ToDollars(cents int) string {
	formattedAmount := fmt.Sprintf("%d.%d0", (cents / 100), (cents % 100))
	return formattedAmount
}

// DB holds information about offer
type DB interface {
	ListAll(ctx context.Context) ([]Offer, error)
	GetCurrentByType(ctx context.Context, offerType OfferType) (*Offer, error)
	Create(ctx context.Context, offer *NewOffer) (*Offer, error)
	Redeem(ctx context.Context, offerID int, isDefault bool) error
	Finish(ctx context.Context, offerID int) error
}

// NewOffer holds information that's needed for creating a new offer
type NewOffer struct {
	Name        string
	Description string

	AwardCreditInCents   int
	InviteeCreditInCents int

	RedeemableCap int

	AwardCreditDurationDays   int
	InviteeCreditDurationDays int

	ExpiresAt time.Time

	Status OfferStatus
	Type   OfferType
}

// UpdateOffer holds fields needed for update an offer
type UpdateOffer struct {
	ID        int
	Status    OfferStatus
	ExpiresAt time.Time
}

// OfferType indicates the type of an offer
type OfferType int

const (
	// FreeCredit is a type of offers used for Free Credit Program
	FreeCredit = OfferType(iota)
	// Referral is a type of offers used for Referral Program
	Referral
)

// OfferStatus indicates the status of an offer
type OfferStatus int

const (

	// Active is a offer status when an offer is currently being used
	Active = OfferStatus(iota)
	// Default is a offer status when an offer is used as a default offer
	Default
	// Done is a default offer status when an offer is not being used currently
	Done
)

// Offer contains info needed for giving users free credits through different offer programs
type Offer struct {
	ID          int
	Name        string
	Description string

	AwardCreditInCents   int
	InviteeCreditInCents int

	AwardCreditDurationDays   int
	InviteeCreditDurationDays int

	RedeemableCap int
	NumRedeemed   int

	ExpiresAt time.Time
	CreatedAt time.Time

	Status OfferStatus
	Type   OfferType
}

// Offers contains a slice of offers.
type Offers []Offer

// OrganizedOffers contains a list of offers organized by status.
type OrganizedOffers struct {
	Active  Offer
	Default Offer
	Done    Offers
}

// OfferSet provides a separation of marketing offers by type.
type OfferSet struct {
	ReferralOffers OrganizedOffers
	FreeCredits    OrganizedOffers
}

// OganizeOffersByStatus organizes offers by OfferStatus.
func (offers Offers) OganizeOffersByStatus() (oo OrganizedOffers) {
	for _, offer := range offers {
		switch offer.Status {
		case Active:
			oo.Active = offer
		case Default:
			oo.Default = offer
		case Done:
			oo.Done = append(oo.Done, offer)
		}
	}
	return oo
}

// OrganizeOffersByType organizes offers by OfferType.
func (offers Offers) OrganizeOffersByType() (os OfferSet) {
	var fc, ro Offers

	for _, offer := range offers {
		switch offer.Type {
		case FreeCredit:
			fc = append(fc, offer)
		case Referral:
			ro = append(ro, offer)
		default:
			continue
		}
	}
	os.FreeCredits = fc.OganizeOffersByStatus()
	os.ReferralOffers = ro.OganizeOffersByStatus()
	return os
}
