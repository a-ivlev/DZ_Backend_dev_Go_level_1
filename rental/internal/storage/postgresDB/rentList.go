package postgresDB

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type RentList struct {
	ClientID     ClientID
	ItemID       ItemID
	RentPriceID  RentPriceID
	Duration     int
	RentalAmount int
	StartRentAt  pgtype.Date
}

type RentListID int

func InsertRentlist(ctx context.Context, dbpool *pgxpool.Pool, rentList RentList) (RentListID, error) {
	const sql = `insert into rent_list (client_id, item_id, rent_price_id, duration, rental_amount) values ($1, $2, $3, $4, $5) returning id;`

	var id RentListID
	err := dbpool.QueryRow(ctx, sql,
		rentList.ClientID,
		rentList.ItemID,
		rentList.RentPriceID,
		rentList.Duration,
		rentList.RentalAmount,
	).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("failed to insert rent_list: %w", err)
	}

	expiresAT := pgtype.Date{
		Time:   time.Now().Add(time.Duration(rentList.Duration) * time.Hour),
		Status: pgtype.Present,
	}
	err = updateItemExpiresAt(ctx, dbpool, rentList.ItemID, expiresAT)
	if err != nil {
		return id, err
	}
	return id, nil
}
