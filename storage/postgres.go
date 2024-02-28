package storage

import (
	"database/sql"
	"fmt"
	"rent-car/config"

	_ "github.com/lib/pq"
)

type Store struct {
	DB  *sql.DB
	Car carRepo
}

func New(cfg config.Config) (Store, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	db, err := sql.Open("postgres", url)
	fmt.Println("err: ", err)
	if err != nil {
		return Store{}, err
	}

	newCar := NewCar(db)
	return Store{
		DB:  db,
		Car: newCar,
	}, nil

}
