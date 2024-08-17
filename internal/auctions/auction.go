package auctions

import (
	"time"

	"github.com/google/uuid"
)

type AuctionItem struct {
	ID                   uuid.UUID
	AbsoluteAuction      bool // absolute auction, is an auction in which the item for sale will be sold regardless of price
	StartingPrice        float64
	ReservePrice         float64 // if ReservePrice is set, the item for sale may not be sold if the final bid is not high enough to satisfy the seller (opposite of AbsoluteAuction))
	ActivateAt           time.Time
	CreatedAt            time.Time
	ModifiedAt           time.Time
	ExpiresAt            time.Time
	SoldAt               time.Time
	SetDownAt            time.Time
	HighestBid           *AuctionBid
	AdministrativelyDown *ItemModeration
	Seller               *AuctionUser
	Comments             []*ItemComment
}

type AuctionUser struct {
	ID            uuid.UUID
	FirstName     string
	LastName      string
	DisplayName   string
	EMail         string
	RegisteredAt  time.Time
	DeactivatedAt time.Time
	DeactivatedBy *AuctionAdmin
}

type AuctionAdmin struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	DisplayName string
	EMail       string
	CreatedAt   time.Time
}

type ItemModeration struct {
	Reason      string
	ModeratedAt time.Time
	ModeratedBy *AuctionAdmin
}

type ItemComment struct {
	CommentText   string
	CommentSentBy *AuctionUser
	CommentSentAt time.Time
	ParentComment *ItemComment
	ChildComments []*ItemComment
}
