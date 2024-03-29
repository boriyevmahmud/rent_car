package service

import (
	"context"
	"fmt"
	"rent-car/api/models"
	"rent-car/storage"
)

type carService struct {
	storage storage.IStorage
}

func NewCarService(storage storage.IStorage) carService {
	return carService{
		storage: storage,
	}
}
func (u carService) Create(ctx context.Context, car models.Car) (string, error) {

	pKey, err := u.storage.Car().Create(ctx, car)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return "", err
	}

	return pKey, nil
}
