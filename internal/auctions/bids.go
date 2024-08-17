package auctions

import "time"

type AuctionBid struct {
	Item      *AuctionItem
	BidAmount float64 //money?? postgresql monetary or use one of the golang money libraries?
	BidBy     *AuctionUser
	BidAt     time.Time
}
