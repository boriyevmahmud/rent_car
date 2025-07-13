package service

import (
	"backend_course/rent_car/api/models"
	"backend_course/rent_car/pkg/logger"
	"backend_course/rent_car/storage"
	"context"
)

type carService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewCarService(storage storage.IStorage, logger logger.ILogger) carService {
	return carService{
		storage: storage,
		logger:  logger,
	}
}

func (s carService) Create(ctx context.Context, car models.CreateCarRequest) (string, error) {

	pKey, err := s.storage.Car().Create(ctx, car)
	if err != nil {
		s.logger.Error("failed to create car", logger.Error(err))
		return "", err
	}

	return pKey, nil
}

func (s carService) Update(ctx context.Context, car models.UpdateCarRequest) (string, error) {

	id, err := s.storage.Car().Update(ctx, car)
	if err != nil {
		s.logger.Error("failed to update car", logger.Error(err))
		return "", err
	}
	return id, nil
}

func (s carService) GetByID(ctx context.Context, id string) (models.GetCarByIDResponse, error) {

	car, err := s.storage.Car().GetByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to get car by ID", logger.Error(err))
		return models.GetCarByIDResponse{}, err
	}

	return car, nil

}

func (s carService) GetAll(ctx context.Context, req models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {

	cars, err := s.storage.Car().GetAll(ctx, req)
	if err != nil {
		s.logger.Error("failed to get all cars", logger.Error(err))
		return models.GetAllCarsResponse{}, err
	}

	return cars, nil
}

func (s carService) GetAvailable(ctx context.Context, req models.GetAvailableCarsRequest) (models.GetAvailableCarsResponse, error) {

	car, err := s.storage.Car().GetAvailable(ctx, req)
	if err != nil {
		s.logger.Error("failed to get available cars", logger.Error(err))
		return models.GetAvailableCarsResponse{}, err
	}

	return car, nil
}

func (s carService) Delete(ctx context.Context, id string) error {

	err := s.storage.Car().Delete(ctx, id)
	if err != nil {
		s.logger.Error("failed to delete car", logger.Error(err))
		return err
	}

	return nil
}
