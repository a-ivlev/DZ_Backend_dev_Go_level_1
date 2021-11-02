package postgresDB

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type RentPrice struct {
	ItemID ItemID
	Name   string
	Price  int
}

type RentPriceID int

func InsertRentPrice(ctx context.Context, dbpool *pgxpool.Pool, rentPrice RentPrice, itemID ItemID) (RentPriceID, error) {
	const sql = `insert into rent_price (item_id, name, price) values ($1, $2, $3) returning id;`

	var id RentPriceID
	err := dbpool.QueryRow(ctx, sql,
		itemID,
		rentPrice.Name,
		rentPrice.Price,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert rent_price: %w", err)
	}
	return id, nil
}
