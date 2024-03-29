package postgres

import (
	"context"
	"fmt"
	"rent-car/api/models"
	"testing"

	"github.com/go-faker/faker/v4"
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
		createdCar, err := carRepo.Get(id)
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
