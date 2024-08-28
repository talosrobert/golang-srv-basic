package auctions

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrRecordNotFound = errors.New("record not found")
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
	INSERT INTO appl.auction_items (starting_price, reserve_price, user_id)
	VALUES ($1, $2, $3)
	RETURNING id, version
	`
	args := []interface{}{ai.StartingPrice, ai.ReservePrice, ai.Seller}

	return m.DB.QueryRow(query, args...).Scan(&ai.ID, &ai.CreatedAt, &ai.ExpiresAt, &ai.IsActive, &ai.Version)
}
func (m AuctionItemModel) Read(id uuid.UUID) (*AuctionItem, error) {
	query := `
	SELECT id, starting_price, reserve_price, is_active, created_at, expires_at, seller, version
	FROM appl.auction_items
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
	query := `
	UPDATE appl.auction_items
	SET starting_price = $1, reserver_price = $2, version = version + 1
	WHERE id = $3
	RETURNING version
	`

	args := []interface{}{ai.StartingPrice, ai.ReservePrice, ai.ID}

	return m.DB.QueryRow(query, args...).Scan(&ai.Version)
}

func (m AuctionItemModel) Delete(id uuid.UUID) error {
	query := `
	DELETE FROM appl.auction_items
	WHERE id = $1
	`

	result, err := m.DB.Exec(query)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m AuctionUserModel) Create(au *AuctionUser) error {
	query := `
	INSERT INTO appl.auction_users (first_name, last_name, display_name, email)
	VALUES ($1, $2, $3, $4)
	RETURNING id, is_active, created_at 
	`

	args := []interface{}{au.FirstName, au.LastName, au.DisplayName, au.EMail}

	return m.DB.QueryRow(query, args...).Scan(&au.ID, &au.IsActive, &au.CreatedAt)
}

func (m AuctionUserModel) Read(id uuid.UUID) (*AuctionUser, error) {
	query := `
	SELECT id, is_active, created_at, first_name, last_name, display_name, email
	FROM appl.auction_users
	WHERE id = $1
	`

	var au AuctionUser

	err := m.DB.QueryRow(query, id).Scan(
		&au.ID,
		&au.IsActive,
		&au.CreatedAt,
		&au.FirstName,
		&au.LastName,
		&au.DisplayName,
		&au.EMail,
	)

	if err != nil {
		return nil, err
	}

	return &au, nil
}
func (m AuctionUserModel) Update(au *AuctionUser) error {
	return nil
}
func (m AuctionUserModel) Delete(id uuid.UUID) error {
	return nil
}
