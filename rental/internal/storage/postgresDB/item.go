package postgresDB

import (
	"context"
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Item struct {
	Name        string
	Description string
	ExpiresAt   pgtype.Date
}

type ItemID int

func InsertItem(ctx context.Context, dbpool *pgxpool.Pool, item Item) (ItemID, error) {
	const sql = `insert into item (name, description) values ($1, $2) returning id;`

	var id ItemID
	err := dbpool.QueryRow(ctx, sql,
		item.Name,
		item.Description,
	).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("failed to insert item: %w", err)
	}
	return id, nil
}

func updateItemExpiresAt(ctx context.Context, dbpool *pgxpool.Pool, id ItemID, expiresAT pgtype.Date) error {
	const sql = `UPDATE item SET expires_at = $2 WHERE id = $1;`
	_, err := dbpool.Exec(ctx, sql,
		id,
		expiresAT.Time,
	)
	if err != nil {
		return fmt.Errorf("failed to update expires_at in item id = %d: %w", id, err)
	}
	return nil
}
