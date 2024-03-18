package storage

import "rent-car/api/models"

type IStorage interface {
	CloseDB()
	Car() ICarStorage
}

type ICarStorage interface {
	Create(models.Car) (string, error)
	// GetByID(models.PrimaryKey) (models.User, error)
	GetAll(request models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
	Update(models.Car) (string, error)
	Delete(string) error
}
