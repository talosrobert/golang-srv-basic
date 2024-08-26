package auctions

import (
	"context"
	"database/sql"
	"time"
)

type AuctionBid struct {
	Item      *AuctionItem
	BidAmount float64 //money?? postgresql monetary or use one of the golang money libraries?
	BidBy     *AuctionUser
	BidAt     time.Time
}

type AuctionBidModel struct {
	DB *sql.DB
}

func (m AuctionBidModel) create(ab *AuctionBid) error {
	query := `
	INSERT INTO appl.auction_bids (item, bid_amount, bid_by)
	VALUES ($1, $2, $3)
	RETURN id, created_at
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []interface{}{&ab.Item, &ab.BidAmount, &ab.BidBy}

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&ab.Item, &ab.BidAmount, &ab.BidBy, &ab.BidAt)
}

func (m AuctionBidModel) read(id int) (*AuctionBid, error) {
	return nil, nil
}

func (m AuctionBidModel) update(bid *AuctionBid) (*AuctionBid, error) {
	return nil, nil
}

func (m AuctionBidModel) delete(id int) error {
	return nil
}
