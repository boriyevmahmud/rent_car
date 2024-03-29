package postgres

import (
	"context"
	"fmt"
	"os"
	"rent-car/config"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	db *pgxpool.Pool
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
	))
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
