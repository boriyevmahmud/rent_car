package postgres

import (
	"backend_course/rent_car/config"
	"backend_course/rent_car/pkg/logger"
	"context"
	"fmt"
	"os"

	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	db  *pgxpool.Pool
	log logger.ILogger
)

var (
	CariD      = "d286f61b-c92e-47ea-94e4-8fea75edb941"
	CustomeriD = "17dec8e0-fd9a-45f6-890c-ef76d37347b1"
	OrderiD    = "b51c185e-0a23-4f3b-aa36-d1d6b89dd80e"
)

func TestMain(m *testing.M) {
	cfg := config.Load()
	conf, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
		cfg.ServiceName,
	))
	if err != nil {
		panic(err)
	}

	conf.MaxConns = 10

	db, err = pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		panic(err)
	}

	conf.MaxConns = 10

	db, err = pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
