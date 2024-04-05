package service

import (
	"backend_course/rent_car/api/models"
	"backend_course/rent_car/pkg/hash"
	"backend_course/rent_car/pkg/logger"
	"backend_course/rent_car/storage"
	"context"
)

type customerService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewCustomerService(storage storage.IStorage, logger logger.ILogger) customerService {
	return customerService{
		storage: storage,
		logger:  logger,
	}
}

func (s customerService) Create(ctx context.Context, customer models.CreateCustomer) (string, error) {
	pKey, err := s.storage.Customer().Create(ctx, customer)
	if err != nil {
		s.logger.Error("failed to create customer", logger.Error(err))
		return "", err
	}
	return pKey, nil
}

func (s customerService) Update(ctx context.Context, customer models.UpdateCustomer, id string) (string, error) {
	id, err := s.storage.Customer().Update(ctx, customer, id)
	if err != nil {
		s.logger.Error("failed to update customer", logger.Error(err))
		return "", err
	}
	return id, nil
}

func (s customerService) Login(ctx context.Context, req models.LoginCustomer) (string, error) {

	hashedPswd, err := s.storage.Customer().GetPassword(ctx, req.Phone)
	if err != nil {
		s.logger.Error("error while getting customer password", logger.Error(err))
		return "", err
	}

	err = hash.CompareHashAndPassword(hashedPswd, req.Password)
	if err != nil {
		s.logger.Error("incorrect password", logger.Error(err))
		return "", err
	}
	return "Login successfully", nil
}

func (s customerService) ChangePassword(ctx context.Context, pass models.ChangePassword) (string, error) {
	msg, err := s.storage.Customer().ChangePassword(ctx, pass)
	if err != nil {
		s.logger.Error("failed to change customer password", logger.Error(err))
		return "", err
	}
	return msg, nil
}

func (s customerService) GetByID(ctx context.Context, id string) (models.Customer, error) {
	customer, err := s.storage.Customer().GetByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to get customer by ID", logger.Error(err))
		return models.Customer{}, err
	}
	return customer, nil
}

func (s customerService) GetAll(ctx context.Context, req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {

	customers, err := s.storage.Customer().GetAll(ctx, req)
	if err != nil {
		s.logger.Error("failed to get all customers", logger.Error(err))
		return models.GetAllCustomersResponse{}, err
	}

	return customers, nil
}

func (s customerService) GetCustomerCars(ctx context.Context, name string, id string, boolean bool) (models.GetCustomerCarsResponse, error) {

	customerCars, err := s.storage.Customer().GetCustomerCars(ctx, name, id, boolean)
	if err != nil {
		s.logger.Error("failed to get customer cars", logger.Error(err))
		return models.GetCustomerCarsResponse{}, err
	}

	return customerCars, nil
}

func (s customerService) Delete(ctx context.Context, id string) error {

	err := s.storage.Customer().Delete(ctx, id)
	if err != nil {
		s.logger.Error("failed to delete customer", logger.Error(err))
		return err
	}

	return nil
}
