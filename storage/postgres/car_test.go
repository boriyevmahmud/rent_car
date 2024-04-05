package postgres

import (
	"backend_course/rent_car/api/models"
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCar(t *testing.T) {
	carRepo := NewCarRepo(db, log)

	reqCar := models.CreateCarRequest{
		Name:       gofakeit.RandomString([]string{"Tesla", "Hundai", "BMW", "Jiguli", "Gentra", "Challenger"}),
		Year:       2010,
		Brand:      faker.Word(),
		Model:      faker.Word(),
		HorsePower: 200,
		Colour:     "Blue",
		EngineCap:  2.0,
	}

	id, err := carRepo.Create(context.Background(), reqCar)
	assert.NoError(t, err)

	createdCar, err := carRepo.GetByID(context.Background(), id)
	assert.NoError(t, err)

	assert.Equal(t, reqCar.Name, createdCar.Name)
	assert.Equal(t, reqCar.Year, createdCar.Year)
	assert.Equal(t, reqCar.Brand, createdCar.Brand)
	assert.Equal(t, reqCar.Model, createdCar.Model)
	assert.Equal(t, reqCar.HorsePower, createdCar.HorsePower)
	assert.Equal(t, reqCar.Colour, createdCar.Colour)
	assert.Equal(t, reqCar.EngineCap, createdCar.EngineCap)
}

func TestUpdateCar(t *testing.T) {
	carRepo := NewCarRepo(db, log)

	reqCar := models.CreateCarRequest{
		Name:       gofakeit.RandomString([]string{"Tesla", "Hundai", "BMW", "Jiguli", "Gentra", "Challenger"}),
		Year:       2010,
		Brand:      faker.Word(),
		Model:      faker.Word(),
		HorsePower: 200,
		Colour:     "Blue",
		EngineCap:  2.0,
	}
	id, err := carRepo.Create(context.Background(), reqCar)
	assert.NoError(t, err)

	updateCar := models.UpdateCarRequest{
		ID:         id,
		Name:       faker.Name(),
		Year:       2015,
		Brand:      faker.Word(),
		Model:      faker.Word(),
		HorsePower: 250,
		Colour:     "Red",
		EngineCap:  2.5,
	}
	updatedID, err := carRepo.Update(context.Background(), updateCar)
	assert.NoError(t, err)
	assert.Equal(t, id, updatedID)

	updatedCar, err := carRepo.GetByID(context.Background(), id)
	assert.NoError(t, err)

	assert.Equal(t, updateCar.Name, updatedCar.Name)
	assert.Equal(t, updateCar.Year, updatedCar.Year)
	assert.Equal(t, updateCar.Brand, updatedCar.Brand)
	assert.Equal(t, updateCar.Model, updatedCar.Model)
	assert.Equal(t, updateCar.HorsePower, updatedCar.HorsePower)
	assert.Equal(t, updateCar.Colour, updatedCar.Colour)
	assert.Equal(t, updateCar.EngineCap, updatedCar.EngineCap)
}
func TestGetByIDCar(t *testing.T) {
	carRepo := NewCarRepo(db, log)

	reqCar := models.CreateCarRequest{
		Name:       gofakeit.RandomString([]string{"Tesla", "Hundai", "BMW", "Jiguli", "Gentra", "Challenger"}),
		Year:       2010,
		Brand:      faker.Word(),
		Model:      faker.Word(),
		HorsePower: 200,
		Colour:     "Blue",
		EngineCap:  2.0,
	}
	id, err := carRepo.Create(context.Background(), reqCar)
	assert.NoError(t, err)

	car, err := carRepo.GetByID(context.Background(), id)
	assert.NoError(t, err)

	assert.Equal(t, reqCar.Name, car.Name)
	assert.Equal(t, reqCar.Year, car.Year)
	assert.Equal(t, reqCar.Brand, car.Brand)
	assert.Equal(t, reqCar.Model, car.Model)
	assert.Equal(t, reqCar.HorsePower, car.HorsePower)
	assert.Equal(t, reqCar.Colour, car.Colour)
	assert.Equal(t, reqCar.EngineCap, car.EngineCap)
}

func TestGetAllCar(t *testing.T) {
	carRepo := NewCarRepo(db, log)

	for i := 0; i < 3; i++ {
		reqCar := models.CreateCarRequest{
			Name:       "LucidAir",
			Year:       2010 + int64(i),
			Brand:      faker.Word(),
			Model:      faker.Word(),
			HorsePower: 200 + int64(i),
			Colour:     "Yellow",
			EngineCap:  2.0 + float32(i)*0.1,
		}
		_, err := carRepo.Create(context.Background(), reqCar)
		assert.NoError(t, err)
	}

	testCases := []struct {
		name     string
		req      models.GetAllCarsRequest
		expected int
	}{
		{"Get 1st page with limit 3", models.GetAllCarsRequest{Search: "", Page: 1, Limit: 2}, 2},
		{"Get 2nd page with limit 2", models.GetAllCarsRequest{Search: "", Page: 2, Limit: 1}, 1},
		{"Search for 'LucidAir' cars", models.GetAllCarsRequest{Search: "LucidAir", Page: 1, Limit: 3}, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cars, err := carRepo.GetAll(context.Background(), tc.req)
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, len(cars.Cars))
			assert.NoError(t, err)
		})
	}
}

func TestGetAvailableCar(t *testing.T) {
	carRepo := NewCarRepo(db, log)

	for i := 0; i < 3; i++ {
		reqCar := models.CreateCarRequest{
			Name:       gofakeit.RandomString([]string{"Tesla", "Hundai", "BMW", "Jiguli", "Gentra", "Challenger"}),
			Year:       2000 + int64(i),
			Brand:      faker.Word(),
			Model:      faker.Word(),
			HorsePower: 200 + int64(i),
			Colour:     "Yellow",
			EngineCap:  2.0 + float32(i)*0.1,
		}
		_, err := carRepo.Create(context.Background(), reqCar)
		assert.NoError(t, err)
	}

	testCases := []struct {
		name     string
		req      models.GetAvailableCarsRequest
		expected int
	}{
		{"Get first page with limit 5", models.GetAvailableCarsRequest{Page: 1, Limit: 5}, 5},
		{"Get second page with limit 3", models.GetAvailableCarsRequest{Page: 2, Limit: 3}, 3},
		{"Search for 'Hundai' available cars", models.GetAvailableCarsRequest{Search: "Hundai", Page: 1, Limit: 3}, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cars, err := carRepo.GetAvailable(context.Background(), tc.req)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, len(cars.Cars))
			assert.NoError(t, err)
		})
	}
}
func TestDeleteCar(t *testing.T) {
	carRepo := NewCarRepo(db, log)

	reqCar := models.CreateCarRequest{
		Name:       gofakeit.RandomString([]string{"Tesla", "Hundai", "BMW", "Jiguli", "Gentra", "Challenger"}),
		Year:       2010,
		Brand:      faker.Word(),
		Model:      faker.Word(),
		HorsePower: 200,
		Colour:     "Blue",
		EngineCap:  2.0,
	}
	id, err := carRepo.Create(context.Background(), reqCar)
	assert.NoError(t, err)

	err = carRepo.Delete(context.Background(), id)
	assert.NoError(t, err)

	_, err = carRepo.GetByID(context.Background(), id)
	assert.Error(t, err)
}
