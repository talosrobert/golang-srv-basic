package data

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrVersionConflict = errors.New("version conflict")
	ErrRecordNotFound  = errors.New("record not found")
)

type Models struct {
	AuctionUser interface {
		Create(au *AuctionUser) error
		Read(id uuid.UUID) (*AuctionUser, error)
		Update(au *AuctionUser) error
		Delete(id uuid.UUID) error
	}
	AuctionItems interface {
		Create(ai *AuctionItem) error
		Read(id uuid.UUID) (*AuctionItem, error)
		Update(ai *AuctionItem) error
		Delete(id uuid.UUID) error
	}
	AuctionBids interface {
		Create(ab *AuctionBid) error
		Read(id int) (*AuctionBid, error)
		Update(ab *AuctionBid) (*AuctionBid, error)
		Delete(id int) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		AuctionUser:  AuctionUserModel{DB: db},
		AuctionItems: AuctionItemModel{DB: db},
		AuctionBids:  AuctionBidModel{DB: db},
	}
}
