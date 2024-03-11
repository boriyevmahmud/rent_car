package postgres

import (
	"database/sql"
	"fmt"
	"rent-car/config"
	"rent-car/storage"

	_ "github.com/lib/pq"
)

type Store struct {
	DB *sql.DB
}

func New(cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return Store{
		DB: db,
	}, nil
}
func (s Store) CloseDB() {
	s.DB.Close()
}

func (s Store) Car() storage.ICarStorage {
	newCar := NewCar(s.DB)

	return &newCar
}

