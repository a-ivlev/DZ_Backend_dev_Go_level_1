package main

import (
	"fmt"
)

func main() {
	fmt.Println("Старт разработки 8 урок")
	// ctx := context.Background()
	// cfg := configs.LoadConfDB()
	//
	// dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
	// if err != nil {
	//	log.Fatal(err)
	// }
	// defer dbpool.Close()
	//
	// db := postgresDB.NewPostgresDB(dbpool)
	//
	//// Поиск клиента по номеру телефона
	// phone := "+79993452776"
	// client, err := db.SearchClientByPhone(ctx, phone)
	// if err != nil {
	//	fmt.Println("Ошибка: ", err)
	// }
	// fmt.Printf("Клиент с номером телефона %s\n %v", phone, client)
	//
	//// Поиск клиента по фамилии
	// lastName := "Гришухин"
	// clients, err := db.SearchClientByLastName(ctx, lastName)
	// if err != nil {
	//	fmt.Println("Ошибка: ", err)
	// }
	// for i, client := range clients {
	//	fmt.Printf("%d Клиент с фамилией %s %v\n", i+1, lastName, client)
	// }
	//
	//// Если указываем в запросе номер телефона, получаем информацию по конкретному клиенту.
	//// Если в место номера телефона пустая строка, получаем информацию по всем клиентам оформившим аренду.
	// items, err := db.SearchRentItemsByPhone(ctx, "+7 411 923 8377")
	// if err != nil {
	//	fmt.Println("Ошибка: ", err)
	// }
	// fmt.Println("Информация что находиться в аренде и у кого:")
	// for _, item := range items {
	//	fmt.Println(item)
	// }
	//
	//// Получает информацию из БД по клиентам кто не вернул арендованную вещь вовремя.
	// SearchClientsNotReturnItemsOnTime, err := db.SearchClientsNotReturnItemsOnTime(ctx)
	// if err != nil {
	//	fmt.Println("Ошибка: ", err)
	// }
	// fmt.Println("В аренде, но должны были уже вернуть.")
	// for _, client := range SearchClientsNotReturnItemsOnTime {
	//	fmt.Println(client)
	// }
}
