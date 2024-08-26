package auctions

import (
	"database/sql"

	"github.com/google/uuid"
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
}

func NewModels(db *sql.DB) Models {
	return Models{
		AuctionUser:  AuctionUserModel{DB: db},
		AuctionItems: AuctionItemModel{DB: db},
	}
}

type AuctionUserModel struct {
	DB *sql.DB
}

type AuctionItemModel struct {
	DB *sql.DB
}

func (m AuctionItemModel) Create(ai *AuctionItem) error {
	query := `
	INSERT INTO auction_items (starting_price, reserve_price, user_id)
	VALUES ($1, $2, $3)
	RETURNING id, version
	`
	args := []interface{}{ai.StartingPrice, ai.ReservePrice, ai.Seller}

	return m.DB.QueryRow(query, args).Scan(&ai.ID, &ai.CreatedAt, &ai.ExpiresAt, &ai.IsActive, &ai.Version)
}
func (m AuctionItemModel) Read(id uuid.UUID) (*AuctionItem, error) {
	query := `
	SELECT id, starting_price, reserve_price, is_active, created_at, expires_at, seller, version
	FROM auction_items
	WHERE id = $1
	`

	var ai AuctionItem

	err := m.DB.QueryRow(query, id).Scan(
		&ai.ID,
		&ai.StartingPrice,
		&ai.ReservePrice,
		&ai.IsActive,
		&ai.CreatedAt,
		&ai.ExpiresAt,
		&ai.Seller,
		&ai.Version,
	)

	if err != nil {
		return nil, err
	}

	return &ai, nil
}
func (m AuctionItemModel) Update(ai *AuctionItem) error {
	return nil
}
func (m AuctionItemModel) Delete(id uuid.UUID) error {
	return nil
}

func (m AuctionUserModel) Create(au *AuctionUser) error {
	return nil
}
func (m AuctionUserModel) Read(id uuid.UUID) (*AuctionUser, error) {
	return nil, nil
}
func (m AuctionUserModel) Update(au *AuctionUser) error {
	return nil
}
func (m AuctionUserModel) Delete(id uuid.UUID) error {
	return nil
}
