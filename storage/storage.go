package storage

import "rent-car/models"

type IStorage interface {
	CloseDB()
	Car() ICarStorage
}

type ICarStorage interface {
	Create(models.Car) (string, error)
	// GetByID(models.PrimaryKey) (models.User, error)
	GetAll(search string) (models.GetAllCarsResponse, error)
	Update(models.Car) (string, error)
	Delete(string) error
}
