package data

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"time"
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

type AuctionUserModel struct {
	DB *sql.DB
}

func (m AuctionUserModel) Create(au *AuctionUser) error {
	query := `
	INSERT INTO appl.auction_users (first_name, last_name, display_name, email)
	VALUES ($1, $2, $3, $4)
	RETURNING id, is_active, created_at 
	`

	args := []interface{}{au.FirstName, au.LastName, au.DisplayName, au.EMail}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&au.ID, &au.IsActive, &au.CreatedAt)
}

func (m AuctionUserModel) Read(id uuid.UUID) (*AuctionUser, error) {
	query := `
	SELECT id, is_active, created_at, first_name, last_name, display_name, email
	FROM appl.auction_users
	WHERE id = $1
	`

	var au AuctionUser

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
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
