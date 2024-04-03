package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"rent-car/api/models"
	"rent-car/pkg"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type carRepo struct {
	db *pgxpool.Pool
}

func NewCar(db *pgxpool.Pool) carRepo {
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

func (c *carRepo) Create(ctx context.Context, car models.Car) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO cars (
		id,
		name,
		brand,
		model,
		hourse_power,
		colour,
		engine_cap,
		year)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8) 
	`

	_, err := c.db.Exec(ctx, query,
		id.String(),
		car.Name, car.Brand,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap,
		car.Year)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (c *carRepo) Update(ctx context.Context, car models.Car) (string, error) {

	fmt.Println("car.id", car.Id)
	query := ` UPDATE cars set
			name=$1,
			brand=$2,
			model=$3,
			hourse_power=$4,
			colour=$5,
			engine_cap=$6,
			year = $7,
			updated_at=CURRENT_TIMESTAMP
		WHERE id = $8 AND deleted_at=0
	`

	_, err := c.db.Exec(ctx, query,
		car.Name,
		car.Brand,
		car.Model,
		car.HoursePower,
		car.Colour,
		car.EngineCap,
		car.Year,
		car.Id)

	if err != nil {
		return "", err
	}

	return car.Id, nil
}

func (c carRepo) GetAll(req models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	var (
		resp   = models.GetAllCarsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter: ", filter)
	rows, err := c.db.Query(context.Background(), `select 
				count(id) OVER(),
				id, 
				name,
				brand,
				model,
				year,
				hourse_power,
				colour,
				engine_cap,
				--created_at::date,
				updated_at
	  FROM cars WHERE deleted_at = 0 `+filter+`
	  `)
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
			// &car.CreatedAt,
			&updateAt); err != nil {
			return resp, err
		}

		car.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Cars = append(resp.Cars, car)
	}
	return resp, nil
}
func (c carRepo) Get(ctx context.Context, id string) (models.Car, error) {
	var (
		car      = models.Car{}
		updateAt sql.NullString
	)

	err := c.db.QueryRow(ctx, `select 
				id, 
				name,
				brand,
				model,
				year,
				hourse_power,
				colour,
				engine_cap,
				updated_at
	  FROM cars WHERE deleted_at = 0 and id =$1
	  `, id).Scan(&car.Id,
		&car.Name,
		&car.Brand,
		&car.Model,
		&car.Year,
		&car.HoursePower,
		&car.Colour,
		&car.EngineCap,
		&updateAt)
	if err != nil {
		return car, err
	}
	car.UpdatedAt = pkg.NullStringToString(updateAt)

	return car, nil
}

func (c *carRepo) Delete(id string) error {

	query := ` UPDATE cars set
			deleted_at = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE id = $1 AND deleted_at=0
	`

	_, err := c.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}
