package auctions

import (
	"time"

	"github.com/google/uuid"
)

type AuctionItem struct {
	ID            uuid.UUID
	StartingPrice float64
	ReservePrice  float64 // if ReservePrice is set, the item for sale may not be sold if the final bid is not high enough to satisfy the seller (opposite of AbsoluteAuction))
	IsActive      bool
	CreatedAt     time.Time
	ExpiresAt     time.Time
	HighestBid    *AuctionBid
	Seller        *AuctionUser
	Comments      []*ItemComment
	Version       int16
}

type AuctionUser struct {
	ID          uuid.UUID
	IsActive    bool
	CreatedAt   time.Time
	FirstName   string
	LastName    string
	DisplayName string
	EMail       string
}

type ItemComment struct {
	CommentText   string
	CommentSentBy *AuctionUser
	CommentSentAt time.Time
	ParentComment *ItemComment
	ChildComments []*ItemComment
}

func (ai *AuctionItem) getAuctionItem(id uuid.UUID) (*AuctionItem, error) {
	return nil, nil
}

func (ai *AuctionItem) createAuctionItem(id uuid.UUID) (*AuctionItem, error) {
	return nil, nil
}

func (ai *AuctionItem) removeAuctionItem(id uuid.UUID) (*AuctionItem, error) {
	return nil, nil
}

func (ai *AuctionItem) updateAuctionItem(id uuid.UUID) (*AuctionItem, error) {
	return nil, nil
}

func (au *AuctionUser) getAuctionUser(id uuid.UUID) (*AuctionUser, error) {
	return nil, nil
}

func (au *AuctionUser) createAuctionUser(id uuid.UUID) (*AuctionUser, error) {
	return nil, nil
}

func (au *AuctionUser) removeAuctionUser(id uuid.UUID) (*AuctionUser, error) {
	return nil, nil
}

func (au *AuctionUser) updateAuctionUser(id uuid.UUID) (*AuctionUser, error) {
	return nil, nil
}

func (ic *ItemComment) getItemComment(id uuid.UUID, user *AuctionUser, parent *ItemComment) (*ItemComment, error) {
	return nil, nil
}

func (ic *ItemComment) createItemComment(id uuid.UUID, user *AuctionUser, parent *ItemComment) (*ItemComment, error) {
	return nil, nil
}

func (ic *ItemComment) removeItemComment(id uuid.UUID, user *AuctionUser, parent *ItemComment) (*ItemComment, error) {
	return nil, nil
}

func (ic *ItemComment) updateItemComment(id uuid.UUID, user *AuctionUser, parent *ItemComment) (*ItemComment, error) {
	return nil, nil
}