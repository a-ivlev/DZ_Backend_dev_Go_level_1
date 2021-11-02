package postgresDB

import (
	"context"
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client struct {
	Id           ClientID
	FirstName    string
	MiddleName   string
	LastName     string
	Phone        string
	RegisteredAt pgtype.Date
}

type ClientID int

func InsertClient(ctx context.Context, dbpool *pgxpool.Pool, client Client) (ClientID, error) {
	var (
		id  ClientID
		sql string
		err error
	)
	if client.MiddleName != "" {
		sql = `insert into client (first_name, middle_name, last_name, phone) values ($1, $2, $3, $4) returning id;`
		err = dbpool.QueryRow(ctx, sql, client.FirstName, client.MiddleName, client.LastName, client.Phone).Scan(&id)
	}
	if client.MiddleName == "" {
		sql = `insert into client (first_name, last_name, phone) values ($1, $2, $3) returning id;`
		err = dbpool.QueryRow(ctx, sql, client.FirstName, client.LastName, client.Phone).Scan(&id)
	}
	if err != nil {
		return id, fmt.Errorf("failed to insert client: %w", err)
	}

	return id, nil
}
