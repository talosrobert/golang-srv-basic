package data

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"time"
)

type AuctionItem struct {
	ID             uuid.UUID
	StartingPrice  float64
	ReservePrice   float64 // if ReservePrice is set, the item for sale may not be sold if the final bid is not high enough to satisfy the seller (opposite of AbsoluteAuction))
	IsActive       bool
	CreatedAt      time.Time
	ExpiresAt      time.Time
	Seller         uuid.UUID
	Comments       []*ItemComment
	LastMinuteBids int16
	Version        int16
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&ai.ID, &ai.StartingPrice, &ai.ReservePrice, &ai.IsActive, &ai.CreatedAt, &ai.ExpiresAt, &ai.Seller, &ai.LastMinuteBids, &ai.Version)
}
func (m AuctionItemModel) Read(id uuid.UUID) (*AuctionItem, error) {
	query := `
	SELECT id, starting_price, reserve_price, is_active, created_at, expires_at, seller, version
	FROM appl.auction_items
	WHERE id = $1
	`

	var ai AuctionItem
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
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
	var query string
	var args []interface{}

	if &ai.ReservePrice != nil {
		query = `
		UPDATE appl.auction_items
		SET starting_price = $1, reserve_price = $2, version = version + 1
		WHERE id = $3 AND version = $4
		RETURNING version
		`
		args = []interface{}{ai.StartingPrice, ai.ReservePrice, ai.ID, ai.Version}
	} else {
		query = `
		UPDATE appl.auction_items
		SET starting_price = $1, version = version + 1
		WHERE id = $2 AND version = $3
		RETURNING version
		`
		args = []interface{}{ai.StartingPrice, ai.ID, ai.Version}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&ai.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrVersionConflict
		default:
			return err
		}
	}
	return nil
}

func (m AuctionItemModel) Delete(id uuid.UUID) error {
	query := `
	DELETE FROM appl.auction_items
	WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query)
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
