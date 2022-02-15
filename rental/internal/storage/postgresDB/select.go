package postgresDB

import (
	"context"
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// SearchClientByPhone ищет клиента по номеру телефона.
func SearchClientByPhone(ctx context.Context, dbpool *pgxpool.Pool, phone string) (*Client, error) {
	client := &Client{}

	const sql = `select id, first_name, last_name, phone, registered_at from client where phone = $1;`
	row := dbpool.QueryRow(ctx, sql, phone)

	// Scan записывает значения столбцов в свойства структуры client
	err := row.Scan(&client.Id, &client.FirstName, &client.LastName, &client.Phone, &client.RegisteredAt.Time)

	if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return client, nil
}

type RentItems struct {
	RentListID    RentListID
	Firstname     string
	Lastname      string
	ItemID        ItemID
	ItemName      string
	RentPriceName string
	RentDuration  int
	RentAmount    int
	StartRent     pgtype.Date
	ExpiresAt     pgtype.Date
}

// SearchRentItemsByPhone возвращает список клиентов которые взяли вещи в аренду, что они взяли и оринтировочный
// срок возврата. Из функции возвращается список RentItems, отсортированный по id записи в таблицу rent_list.
func SearchRentItemsByPhone(ctx context.Context, dbpool *pgxpool.Pool, phone string) ([]RentItems, error) {
	var (
		sql  string
		rows pgx.Rows
		err  error
	)
	if phone != "" {
		sql = `select rent_list.id, client.first_name, client.last_name, item.id, item.name, rent_price.name as rental, rent_list.duration,  rent_list.rental_amount, rent_list.start_rent_at, item.expires_at from rent_list
join client ON rent_list.client_id = client.id
join item ON rent_list.item_id = item.id
join rent_price ON rent_list.rent_price_id = rent_price.id
WHERE client.phone = $1 ORDER BY rent_list.id;`
		rows, err = dbpool.Query(ctx, sql, phone)
	}
	if phone == "" {
		sql = `select rent_list.id, client.first_name, client.last_name, item.id, item.name, rent_price.name as rental, rent_list.duration,  rent_list.rental_amount, rent_list.start_rent_at, item.expires_at from rent_list
join client ON rent_list.client_id = client.id
join item ON rent_list.item_id = item.id
join rent_price ON rent_list.rent_price_id = rent_price.id
ORDER BY rent_list.id;`
		rows, err = dbpool.Query(ctx, sql)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	// Вызов Close нужен, чтобы вернуть соединение в пул
	defer rows.Close()
	// В слайс hints будут собраны все строки, полученные из базы
	var hints []RentItems
	// rows.Next() итерируется по всем строкам, полученным из базы.
	for rows.Next() {
		var hint RentItems
		// Scan записывает значения столбцов в свойства структуры hint
		err = rows.Scan(&hint.RentListID, &hint.Firstname, &hint.Lastname, &hint.ItemID, &hint.ItemName, &hint.RentPriceName, &hint.RentDuration, &hint.RentAmount, &hint.StartRent.Time, &hint.ExpiresAt.Time)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		hints = append(hints, hint)
	}
	// Проверка, что во время выборки данных не происходило ошибок
	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to read response: %w", rows.Err())
	}
	return hints, nil
}

type ClientsNotReturnItems struct {
	RentListID    RentListID
	Firstname     string
	Lastname      string
	ItemID        ItemID
	ItemName      string
	RentPriceName string
	RentDuration  int
	RentAmount    int
	StartRent     pgtype.Date
	ExpiresAt     pgtype.Date
	NowAt         pgtype.Date
}

// SearchClientsNotReturnItems показывает список клиентов не вернувщих вовремя вещи взятые в аренду.
// Список вещей находящихся у данных клиентов в аренде, дату и время когда вещь должны вернуть.
// Из функции возвращается список ClientsNotReturnItems.
func SearchClientsNotReturnItems(ctx context.Context, dbpool *pgxpool.Pool) ([]ClientsNotReturnItems, error) {
	const sql = `select rent_list.id, client.first_name, client.last_name, item.id, item.name, rent_price.name as rental, rent_list.duration,  rent_list.rental_amount, rent_list.start_rent_at, item.expires_at, now() as now from rent_list 
join client ON rent_list.client_id = client.id 
join item ON rent_list.item_id = item.id
join rent_price ON rent_list.rent_price_id = rent_price.id
WHERE item.expires_at < now() and rent_list.end_rent_at IS NULL;`

	rows, err := dbpool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	// Вызов Close нужен, чтобы вернуть соединение в пул
	defer rows.Close()
	// В слайс hints будут собраны все строки, полученные из базы
	var hints []ClientsNotReturnItems
	// rows.Next() итерируется по всем строкам, полученным из базы.
	for rows.Next() {
		var hint ClientsNotReturnItems
		// Scan записывает значения столбцов в свойства структуры hint
		err = rows.Scan(&hint.RentListID, &hint.Firstname, &hint.Lastname, &hint.ItemID, &hint.ItemName, &hint.RentPriceName, &hint.RentDuration, &hint.RentAmount, &hint.StartRent.Time, &hint.ExpiresAt.Time, &hint.NowAt.Time)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		hints = append(hints, hint)
	}
	// Проверка, что во время выборки данных не происходило ошибок
	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to read response: %w", rows.Err())
	}
	return hints, nil
}
