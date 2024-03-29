package storage

import (
	"context"
	"rent-car/api/models"
)

type IStorage interface {
	CloseDB()
	Car() ICarStorage
}

type ICarStorage interface {
	Create(context.Context, models.Car) (string, error)
	Get(string) (models.Car, error)
	GetAll(request models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
	Update(models.Car) (string, error)
	Delete(string) error
}
