package service

import (
	"backend_course/rent_car/api/models"
	"backend_course/rent_car/pkg/logger"
	"backend_course/rent_car/storage"
	"context"
)

type orderService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewOrderService(storage storage.IStorage, logger logger.ILogger) orderService {
	return orderService{
		storage: storage,
		logger:  logger,
	}
}

func (s orderService) Create(ctx context.Context, order models.CreateOrder) (string, error) {
	pKey, err := s.storage.Order().Create(ctx, order)
	if err != nil {
		s.logger.Error("failed to create order", logger.Error(err))
		return "", err
	}
	return pKey, nil
}

func (s orderService) Update(ctx context.Context, order models.UpdateOrder) (string, error) {
	id, err := s.storage.Order().Update(ctx, order)
	if err != nil {
		s.logger.Error("failed to update order", logger.Error(err))
		return "", err
	}
	return id, nil
}

func (s orderService) GetByID(ctx context.Context, id string) (models.GetOrderResponse, error) {
	order, err := s.storage.Order().GetByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to get order by ID", logger.Error(err))
		return models.GetOrderResponse{}, err
	}
	return order, nil
}

func (s orderService) GetAll(ctx context.Context, req models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error) {

	orders, err := s.storage.Order().GetAll(ctx, req)
	if err != nil {
		s.logger.Error("failed to get all orders", logger.Error(err))
		return models.GetAllOrdersResponse{}, err
	}

	return orders, nil
}

func (s orderService) Delete(ctx context.Context, id string) error {

	err := s.storage.Order().Delete(ctx, id)
	if err != nil {
		s.logger.Error("failed to delete order", logger.Error(err))
		return err
	}

	return nil
}
