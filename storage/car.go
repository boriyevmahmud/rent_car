package storage

import (
	"database/sql"
	"fmt"
	"rent-car/models"

	"github.com/google/uuid"
)

type carRepo struct {
	db *sql.DB
}

func NewCar(db *sql.DB) carRepo {
	return carRepo{
		db: db,
	}
}

/*
create (body) id,err
update (body) id,err
delete (id) err
get (id) body,err
getAll (search) []body,count,err
*/

func (c *carRepo) Create(car models.Car) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO cars (
		id,
		name,
		brand,
		model,
		hourse_power,
		colour,
		engine_cap)
		VALUES($1,$2,$3,$4,$5,$6,$7) 
	`

	res, err := c.db.Exec(query,
		id.String(),
		car.Name, car.Brand,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap)

	if err != nil {
		return "", err
	}

	fmt.Printf("%+v\n", res)

	return id.String(), nil
}
