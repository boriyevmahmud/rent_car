package storage

import (
	"backend_course/rent_car/api/models"
	"context"
	"time"
)

type IStorage interface {
	CloseDB()
	Car() ICarStorage
	Customer() ICustomerStorage
	Order() IOrderStorage
	Redis() IRedisStorage
}

type ICarStorage interface {
	Create(ctx context.Context, car models.CreateCarRequest) (string, error)
	Update(ctx context.Context, car models.UpdateCarRequest) (string, error)
	GetByID(ctx context.Context, id string) (models.GetCarByIDResponse, error)
	GetAll(ctx context.Context, req models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
	GetAvailable(ctx context.Context, req models.GetAvailableCarsRequest) (models.GetAvailableCarsResponse, error)
	Delete(ctx context.Context, id string) error
}

type ICustomerStorage interface {
	Create(ctx context.Context, customer models.CreateCustomer) (string, error)
	Update(ctx context.Context, customer models.UpdateCustomer, id string) (string, error)
	Login(ctx context.Context, pass models.LoginCustomer) (string, error)
	ChangePassword(ctx context.Context, pass models.ChangePassword) (string, error)
	GetByID(ctx context.Context, id string) (models.Customer, error)
	GetAll(ctx context.Context, req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error)
	GetCustomerCars(ctx context.Context, name string, id string, boolean bool) (models.GetCustomerCarsResponse, error)
	Delete(ctx context.Context, id string) error
	GetPassword(ctx context.Context, phone string) (string, error)
	GetByLogin(context.Context, string) (models.Customer, error)
}

type IOrderStorage interface {
	Create(ctx context.Context, order models.CreateOrder) (string, error)
	Update(ctx context.Context, order models.UpdateOrder) (string, error)
	GetByID(ctx context.Context, id string) (models.GetOrderResponse, error)
	GetAll(ctx context.Context, req models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error)
	Delete(ctx context.Context, id string) error
	DeleteHard(ctx context.Context, id string) error
}

type IRedisStorage interface {
	SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string) interface{}
	Del(ctx context.Context, key string) error
}
