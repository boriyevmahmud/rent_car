package service

import (
	"rent-car/storage"
)

type IServiceManager interface {
	Car() carService
}

type Service struct {
	carService carService
}

func New(storage storage.IStorage) Service {
	services := Service{}
	services.carService = NewCarService(storage)

	return services
}

func (s Service) Car() carService {
	return s.carService
}
