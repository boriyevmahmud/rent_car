package postgres

import (
	"backend_course/rent_car/api/models"
	"backend_course/rent_car/pkg/logger"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CarRepo struct {
	db     *pgxpool.Pool
	logger logger.ILogger
}

func NewCarRepo(db *pgxpool.Pool, log logger.ILogger) CarRepo {
	return CarRepo{
		db:     db,
		logger: log,
	}
}

func (c *CarRepo) Create(ctx context.Context, car models.CreateCarRequest) (string, error) {
	id := uuid.New().String()

	query := `INSERT INTO cars (
		id,
		name,
		year,
		brand,
		model,
		horse_power,
		colour,
		engine_cap,
		created_at,
		updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	_, err := c.db.Exec(ctx, query,
		id,
		car.Name,
		car.Year,
		car.Brand,
		car.Model,
		car.HorsePower,
		car.Colour,
		car.EngineCap,
	)

	if err != nil {
		c.logger.Error("failed to create car in database", logger.Error(err))
		return "", err
	}

	return id, nil
}

func (c *CarRepo) Update(ctx context.Context, car models.UpdateCarRequest) (string, error) {
	query := `UPDATE cars SET
		name = $1,
		year = $2,
		brand = $3,
		model = $4,
		horse_power = $5,
		colour = $6,
		engine_cap = $7,
		updated_at = CURRENT_TIMESTAMP
	WHERE id = $8`

	_, err := c.db.Exec(ctx, query,
		car.Name,
		car.Year,
		car.Brand,
		car.Model,
		car.HorsePower,
		car.Colour,
		car.EngineCap,
		car.ID,
	)

	if err != nil {
		c.logger.Error("failed to update car in database", logger.Error(err))
		return "", err
	}

	return car.ID, nil
}

func (c *CarRepo) GetByID(ctx context.Context, id string) (models.GetCarByIDResponse, error) {
	var (
		car        = models.GetCarByIDResponse{}
		name       sql.NullString
		year       sql.NullInt64
		brand      sql.NullString
		model      sql.NullString
		horsepower sql.NullInt64
		colour     sql.NullString
		enginecap  sql.NullFloat64
		createdat  sql.NullString
		updatedat  sql.NullString
	)

	query := `SELECT
		id,
		name,
		year,
		brand,
		model,
		horse_power,
		colour,
		engine_cap,
		created_at,
		updated_at
	FROM cars
	WHERE id = $1 and deleted_at = 0`

	row := c.db.QueryRow(ctx, query, id)

	err := row.Scan(
		&car.ID,
		&name,
		&year,
		&brand,
		&model,
		&horsepower,
		&colour,
		&enginecap,
		&createdat,
		&updatedat,
	)

	if err != nil {
		c.logger.Error("failed to get car by ID from database", logger.Error(err))
		return models.GetCarByIDResponse{}, err
	}

	car.Name = name.String
	car.Year = year.Int64
	car.Brand = brand.String
	car.Model = model.String
	car.HorsePower = horsepower.Int64
	car.Colour = colour.String
	car.EngineCap = float32(enginecap.Float64)
	car.CreatedAt = createdat.String
	car.UpdatedAt = updatedat.String

	return car, nil
}

func (c *CarRepo) GetAll(ctx context.Context, req models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	var (
		resp       = models.GetAllCarsResponse{}
		name       sql.NullString
		year       sql.NullInt64
		brand      sql.NullString
		model      sql.NullString
		horsepower sql.NullInt64
		colour     sql.NullString
		enginecap  sql.NullFloat64
		createdat  sql.NullString
		updatedat  sql.NullString
		filter     string
	)

	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = fmt.Sprintf(` AND (name ILIKE '%%%v%%' OR brand ILIKE '%%%v%%' OR model ILIKE '%%%v%%')`, req.Search, req.Search, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)

	query := `SELECT 
		id, 
		name, 
		year, 
		brand, 
		model, 
		horse_power, 
		colour, 
		engine_cap, 
		created_at, 
		updated_at
	FROM cars WHERE deleted_at = 0` + filter

	rows, err := c.db.Query(ctx, query)
	if err != nil {
		c.logger.Error("failed to get all cars from database", logger.Error(err))
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var car models.Car

		err := rows.Scan(
			&car.ID,
			&name,
			&year,
			&brand,
			&model,
			&horsepower,
			&colour,
			&enginecap,
			&createdat,
			&updatedat,
		)

		if err != nil {
			c.logger.Error("failed to scan  all cars from database", logger.Error(err))
			return models.GetAllCarsResponse{}, err
		}

		resp.Cars = append(resp.Cars, models.Car{
			ID:         car.ID,
			Name:       name.String,
			Year:       year.Int64,
			Brand:      brand.String,
			Model:      model.String,
			HorsePower: horsepower.Int64,
			Colour:     colour.String,
			EngineCap:  float32(enginecap.Float64),
			CreatedAt:  createdat.String,
			UpdatedAt:  updatedat.String,
		})
	}

	countQuery := `SELECT COUNT(id) FROM cars`
	err = c.db.QueryRow(ctx, countQuery).Scan(&resp.Count)
	if err != nil {
		c.logger.Error("failed to get cars count from database", logger.Error(err))
		return resp, err
	}

	return resp, nil
}

func (c *CarRepo) GetAvailable(ctx context.Context, req models.GetAvailableCarsRequest) (models.GetAvailableCarsResponse, error) {
	var (
		cars       models.GetAvailableCarsResponse
		count      uint64
		filter     string
		name       sql.NullString
		year       sql.NullInt64
		brand      sql.NullString
		model      sql.NullString
		horsepower sql.NullInt64
		colour     sql.NullString
		enginecap  sql.NullFloat64
		createdat  sql.NullString
		updatedat  sql.NullString
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = fmt.Sprintf(` AND (name ILIKE '%%%v%%' OR brand ILIKE '%%%v%%' OR model ILIKE '%%%v%%')`, req.Search, req.Search, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)

	query := `SELECT
			id,
			name,
			year,
			brand,
			model,
			horse_power,
			colour,
			engine_cap,
			created_at,
			updated_at
		FROM cars
		WHERE id NOT IN (
			SELECT DISTINCT car_id
			FROM orders
			WHERE from_date <= NOW() AND to_date >= NOW()
		)
	` + filter

	rows, err := c.db.Query(ctx, query)
	if err != nil {
		c.logger.Error("failed to get available cars from database", logger.Error(err))
		return models.GetAvailableCarsResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var car models.Car
		err := rows.Scan(
			&car.ID,
			&name,
			&year,
			&brand,
			&model,
			&horsepower,
			&colour,
			&enginecap,
			&createdat,
			&updatedat,
		)

		if err != nil {
			c.logger.Error("failed to scan available cars from database", logger.Error(err))
			return models.GetAvailableCarsResponse{}, err
		}

		cars.Cars = append(cars.Cars, models.Car{
			ID:         car.ID,
			Name:       name.String,
			Year:       year.Int64,
			Brand:      brand.String,
			Model:      model.String,
			HorsePower: horsepower.Int64,
			Colour:     colour.String,
			EngineCap:  float32(enginecap.Float64),
			CreatedAt:  createdat.String,
			UpdatedAt:  updatedat.String,
		})
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS subquery", query)
	err = c.db.QueryRow(ctx, countQuery).Scan(&count)
	cars.Count = count
	if err != nil {
		c.logger.Error("failed to get count of available cars", logger.Error(err))
		return models.GetAvailableCarsResponse{}, err
	}

	return cars, nil
}

func (c *CarRepo) Delete(ctx context.Context, id string) error {
	query := `UPDATE cars SET deleted_at = date_part('epoch', CURRENT_TIMESTAMP)::int WHERE id = $1 AND deleted_at = 0`

	_, err := c.db.Exec(ctx, query, id)
	if err != nil {
		c.logger.Error("failed to delete car", logger.Error(err), logger.String("car_id", id))
		return err
	}

	return nil
}

func (c *CarRepo) DeleteHard(ctx context.Context, id string) error {
	query := `DELETE FROM cars WHERE id = $1`

	_, err := c.db.Exec(ctx, query, id)
	if err != nil {
		c.logger.Error("failed to hard delete car", logger.Error(err), logger.String("car_id", id))
		return err
	}

	return nil
}
