package auctions

import "time"

type Bids struct {
	CurrentBid    float64
	CurrentBidder *AuctionUser
	BiddedAt      time.Time
}
