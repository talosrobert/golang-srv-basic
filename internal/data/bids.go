package data

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
	CREATE OR REPLACE FUNCTION appl.add_bid(item_id uuid, bid_amount numeric, bid_by uuid)
	 RETURNS integer
	 LANGUAGE plpgsql
	AS $function$
	DECLARE
	    new_bid_id integer;
	BEGIN
	    -- Debugging: Print out the input values
	    RAISE NOTICE 'item_id: %, bid_amount: %, bid_by: %', item_id, bid_amount, bid_by;

	    -- Check if the item has expired
	    IF now() > (SELECT expires_at FROM appl.auction_items WHERE id = item_id) THEN
		RAISE EXCEPTION 'Auction Item with ID % has already expired', item_id;
	    ELSE
		-- Insert the bid and return the new bid's ID
		INSERT INTO appl.auction_bids (item, bid_amount, bid_by)
		VALUES (item_id, bid_amount, bid_by)
		RETURNING id INTO new_bid_id;

		-- Debugging: Print out the new bid ID
		RAISE NOTICE 'New bid ID: %', new_bid_id;
		
		RETURN new_bid_id;
	    END IF;
	END;
	$function$`

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
