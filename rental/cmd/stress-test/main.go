package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"rental-app/rental/configs"
	"rental-app/rental/pkg/attack"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	if tz := os.Getenv("TZ"); tz != "" {
		var err error
		// Загрузка таймзоны в приложение.
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			log.Printf("error loading location '%s': %v\n", tz, err)
		}
	}

	// Выводит текущую таймзону.
	tnow := time.Now()
	tz, _ := tnow.Zone()
	log.Printf("Local time zone %s. Service started at %s", tz,
		tnow.Format("2006-01-02T15:04:05.000 MST"))

	ctx := context.Background()
	cfg := configs.LoadConfDB()

	dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Для проведения стресс-теста для БД Postgres перейдите на вкладку attack")

	})

	http.HandleFunc("/attack", func(w http.ResponseWriter, r *http.Request) {
		duration := 10 * time.Second
		threads := 1000

		res := attack.Attack(ctx, duration, threads, dbpool)
		qps := res.QueriesPerformed / uint64(res.Duration.Seconds())

		fmt.Fprintf(w, "\tstart attack\n duration: %s\n threads: %d\n queries: %d\n QPS: %d\n", res.Duration, res.Threads, res.QueriesPerformed, qps)
	})

	port := os.Getenv("PORT")
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println(err)
	}
}
