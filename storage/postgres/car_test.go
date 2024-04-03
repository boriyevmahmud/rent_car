package postgres

import (
	"context"
	"errors"
	"fmt"
	"rent-car/api/models"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateCar(t *testing.T) {
	carRepo := NewCar(db)

	reqCar := models.Car{
		Name:  faker.Name(),
		Year:  2010,
		Brand: faker.Word(),
	}

	id, err := carRepo.Create(context.Background(), reqCar)
	if assert.NoError(t, err) {
		createdCar, err := carRepo.Get(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCar.Name, createdCar.Name)
			assert.Equal(t, reqCar.Year, createdCar.Year)
			assert.Equal(t, reqCar.Brand, createdCar.Brand)
		} else {
			return
		}
		fmt.Println("Created car", createdCar)
	}
}

func TestUpdateCar(t *testing.T) {
	carRepo := NewCar(db)

	reqCar := models.Car{
		Name:        faker.Name(),
		Year:        2010,
		Brand:       faker.Word(),
		Model:       faker.Word(),
		HoursePower: 200,
		Colour:      "Blue",
		EngineCap:   2.0,
	}
	id, err := carRepo.Create(context.Background(), reqCar)
	assert.NoError(t, err)

	reqCar.Id = id
	reqCar.Name = faker.Email()
	reqCar.Year = 2050
	reqCar.Colour = "green"
	id, err = carRepo.Update(context.Background(), reqCar)
	if assert.NoError(t, err) {
		createdCar, err := carRepo.Get(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqCar.Name, createdCar.Name)
			assert.Equal(t, reqCar.Year, createdCar.Year)
			assert.Equal(t, reqCar.Brand, createdCar.Brand)
		} else {
			return
		}
		fmt.Println("Created car", createdCar)
	}
}

func TestGetCar(t *testing.T) {
	carRepo := NewCar(db)

	reqCar := models.Car{
		Name:        faker.Name(),
		Year:        2010,
		Brand:       faker.Word(),
		Model:       faker.Word(),
		HoursePower: 200,
		Colour:      "Blue",
		EngineCap:   2.0,
	}
	id, err := carRepo.Create(context.Background(), reqCar)
	assert.NoError(t, err)

	car, err := carRepo.Get(context.Background(), id)
	assert.NoError(t, err)

	assert.Equal(t, reqCar.Name, car.Name)
	assert.Equal(t, reqCar.Year, car.Year)
	assert.Equal(t, reqCar.Brand, car.Brand)
	assert.Equal(t, reqCar.Model, car.Model)
	assert.Equal(t, reqCar.HoursePower, car.HoursePower)
	assert.Equal(t, reqCar.Colour, car.Colour)
	assert.Equal(t, reqCar.EngineCap, car.EngineCap)
}

func TestDeleteCar(t *testing.T) {
	carRepo := NewCar(db)

	reqCar := models.Car{
		Name:        faker.Name(),
		Year:        2010,
		Brand:       faker.Word(),
		Model:       faker.Word(),
		HoursePower: 200,
		Colour:      "Blue",
		EngineCap:   2.0,
	}
	id, err := carRepo.Create(context.Background(), reqCar)
	assert.NoError(t, err)

	err = carRepo.Delete(id)
	if assert.NoError(t, err) {
		_, err := carRepo.Get(context.Background(), id)
		if assert.ErrorIs(t, err, pgx.ErrNoRows) {
			fmt.Println("Test passed: ")
		}
	}
}

func TestGetAllCar(t *testing.T) {
	carRepo := NewCar(db)

	reqCar := models.Car{
		Name:        faker.Name(),
		Year:        2010,
		Brand:       faker.Word(),
		Model:       faker.Word(),
		HoursePower: 200,
		Colour:      "Blue",
		EngineCap:   2.0,
	}
	_, err := carRepo.Create(context.Background(), reqCar)
	assert.NoError(t, err)

	reqGetAll := models.GetAllCarsRequest{
		Search: reqCar.Name,
		Page:   1,
		Limit:  10,
	}

	cars, err := carRepo.GetAll(reqGetAll)
	if assert.NoError(t, err) {
		if !(len(cars.Cars) > 0 && len(cars.Cars) <= 10) {
			assert.Error(t, errors.New("Invalid number of cars"), fmt.Sprintf("length of cars: %d", len(cars.Cars)))
			return
		}
		for _, car := range cars.Cars {
			if assert.Contains(t, car.Name, reqGetAll.Search) {
				fmt.Println("test passed")
			}
		}
	}

	// assert.Equal(t, reqCar.Name, car.Name)
	// assert.Equal(t, reqCar.Year, car.Year)
	// assert.Equal(t, reqCar.Brand, car.Brand)
	// assert.Equal(t, reqCar.Model, car.Model)
	// assert.Equal(t, reqCar.HoursePower, car.HoursePower)
	// assert.Equal(t, reqCar.Colour, car.Colour)
	// assert.Equal(t, reqCar.EngineCap, car.EngineCap)
}
