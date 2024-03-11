package postgres

import (
	"database/sql"
	"fmt"
	"rent-car/models"
	"rent-car/pkg"

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

	_, err := c.db.Exec(query,
		id.String(),
		car.Name, car.Brand,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (c *carRepo) Update(car models.Car) (string, error) {

	query := ` UPDATE cars set
			name=$1,
			brand=$2,
			model=$3,
			hourse_power=$4,
			colour=$5,
			engine_cap=$6,
			updated_at=CURRENT_TIMESTAMP
		WHERE id = $7 AND deleted_at=0
	`

	_, err := c.db.Exec(query,
		car.Name, car.Brand,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap, car.Id)

	if err != nil {
		return "", err
	}

	return car.Id, nil
}

func (c carRepo) GetAll(search string) (models.GetAllCarsResponse, error) {
	var (
		resp   = models.GetAllCarsResponse{}
		filter = ""
	)

	if search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, search)
	}

	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(`select 
				count(id) OVER(),
				id, 
				name,
				brand,
				model,
				year,
				hourse_power,
				colour,
				engine_cap,
				created_at::date,
				updated_at
	  FROM cars WHERE deleted_at = 0 ` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			car      = models.Car{}
			updateAt sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&car.Id,
			&car.Name,
			&car.Brand,
			&car.Model,
			&car.Year,
			&car.HoursePower,
			&car.Colour,
			&car.EngineCap,
			&car.CreatedAt,
			&updateAt); err != nil {
			return resp, err
		}

		car.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Cars = append(resp.Cars, car)
	}
	return resp, nil
}

func (c *carRepo) Delete(id string) error {

	query := ` UPDATE cars set
			deleted_at = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE id = $1 AND deleted_at=0
	`

	_, err := c.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
