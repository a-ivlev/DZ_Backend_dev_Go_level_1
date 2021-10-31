package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("POSTGRES")
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatal(err)
	}

	// В метод Scan передаётся ссылка на переменную greeting
	// туда будет записан результат работы запроса.
	// Если выборка идет из нескольких столбцов,
	// то для каждого столбца в Scan передаётся по одной ссылке
	var greeting string
	err = db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(greeting)
}
