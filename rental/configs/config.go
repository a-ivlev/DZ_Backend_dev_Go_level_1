package configs

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func LoadConfDB() *pgxpool.Config {
	// Получаем строку с параметрами для подключения к базе данных
	dsn := os.Getenv("POSTGRES_DSN")
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Pool соединений обязательно ограничивать сверху,
	// так как иначе есть потенциальная опасность превысить лимит соединений с	базой.
	cfg.MaxConns = 64
	cfg.MinConns = 64

	// HealthCheckPeriod - частота проверки работоспособности
	// соединения с Postgres
	cfg.HealthCheckPeriod = 1 * time.Minute

	// MaxConnLifetime - сколько времени будет жить соединение.
	// Так как большого смысла удалять живые соединения нет,
	// можно устанавливать большие значения
	cfg.MaxConnLifetime = 24 * time.Hour

	// MaxConnIdleTime - время жизни неиспользуемого соединения,
	// если запросов не поступало, то соединение закроется.
	cfg.MaxConnIdleTime = 30 * time.Minute

	// ConnectTimeout устанавливает ограничение по времени
	// на весь процесс установки соединения и аутентификации.
	cfg.ConnConfig.ConnectTimeout = 1 * time.Second

	// Лимиты в net.Dialer позволяют достичь предсказуемого
	// поведения в случае обрыва сети.
	cfg.ConnConfig.DialFunc = (&net.Dialer{
		KeepAlive: cfg.HealthCheckPeriod,
		// Timeout на установку соединения гарантирует,
		// что не будет зависаний при попытке установить соединение.
		Timeout: cfg.ConnConfig.ConnectTimeout,
	}).DialContext

	// pgx предоставляет набор адаптеров для популярных логеров
	// это позволяет организовать сбор ошибок при работе с базой
	// @see: https://github.com/jackc/pgx/tree/master/log
	// cfg.ConnConfig = zerologadapter.NewLogger(logger)

	return cfg
}
