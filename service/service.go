package service

import (
	"backend_course/rent_car/pkg/logger"
	"backend_course/rent_car/storage"
)

type IServiceManager interface {
	Car() carService
	Customer() customerService
	Order() orderService
}

type Service struct {
	carService      carService
	customerService customerService
	orderService    orderService

	logger logger.ILogger
}

func New(storage storage.IStorage, log logger.ILogger) Service {
	return Service{
		carService:      NewCarService(storage, log),
		customerService: NewCustomerService(storage, log),
		orderService:    NewOrderService(storage, log),
		logger:          log,
	}
}

func (s Service) Car() carService {
	return s.carService
}

func (s Service) Customer() customerService {
	return s.customerService
}

func (s Service) Order() orderService {
	return s.orderService
}
