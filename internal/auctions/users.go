package auctions

import (
	"time"

	"github.com/google/uuid"
)

type AuctionUser struct {
	ID          uuid.UUID
	IsActive    bool
	CreatedAt   time.Time
	FirstName   string
	LastName    string
	DisplayName string
	EMail       string
}
