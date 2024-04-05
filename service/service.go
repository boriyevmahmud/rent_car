package service

import (
	"rent-car/pkg/logger"
	"rent-car/storage"
)

type IServiceManager interface {
	Car() carService
}

type Service struct {
	carService carService
}

func New(storage storage.IStorage, log logger.ILogger) Service {
	services := Service{}
	services.carService = NewCarService(storage, log)

	return services
}

func (s Service) Car() carService {
	return s.carService
}
