package auctions

import (
	"time"

	"github.com/google/uuid"
)

type AuctionItem struct {
	ID             uuid.UUID
	StartingPrice  float64
	ReservePrice   float64 // if ReservePrice is set, the item for sale may not be sold if the final bid is not high enough to satisfy the seller (opposite of AbsoluteAuction))
	IsActive       bool
	CreatedAt      time.Time
	ExpiresAt      time.Time
	Seller         uuid.UUID
	Comments       []*ItemComment
	LastMinuteBids int16
	Version        int16
}
