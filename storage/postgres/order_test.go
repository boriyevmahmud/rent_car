package postgres

import (
	"backend_course/rent_car/api/models"
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	orderRepo := NewOrderRepo(db, log)

	reqOrder := models.CreateOrder{
		CarId:      CariD,
		CustomerId: CustomeriD,
		FromDate:   time.Now().Format(time.RFC3339),
		ToDate:     time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
		Status:     "active",
		Paid:       false,
	}

	id, err := orderRepo.Create(context.Background(), reqOrder)
	assert.NoError(t, err)

	createdOrder, err := orderRepo.GetByID(context.Background(), id)
	assert.NoError(t, err)

	assert.Equal(t, reqOrder.CarId, createdOrder.Car.ID)
	assert.Equal(t, reqOrder.CustomerId, createdOrder.Customer.ID)
	assert.Equal(t, reqOrder.Status, createdOrder.Status)
	assert.Equal(t, reqOrder.Paid, createdOrder.Paid)
}

func TestUpdateOrder(t *testing.T) {
	orderRepo := NewOrderRepo(db, log)

	updateOrder := models.UpdateOrder{
		Id:         OrderiD,
		CarId:      CariD,
		CustomerId: CustomeriD,
		FromDate:   time.Now().AddDate(0, 0, 2).Format(time.RFC3339),
		ToDate:     time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Status:     "active",
		Paid:       true,
	}

	updatedID, err := orderRepo.Update(context.Background(), updateOrder)
	assert.NoError(t, err)
	assert.Equal(t, updateOrder.Id, updatedID)

	updatedOrder, err := orderRepo.GetByID(context.Background(), updateOrder.Id)
	assert.NoError(t, err)

	assert.Equal(t, updateOrder.Id, updatedOrder.Id)
	assert.Equal(t, updateOrder.Status, updatedOrder.Status)
	assert.Equal(t, updateOrder.Paid, updatedOrder.Paid)
}

func TestGetByIDOrder(t *testing.T) {
	orderRepo := NewOrderRepo(db, log)

	order, err := orderRepo.GetByID(context.Background(), OrderiD)
	assert.NoError(t, err)

	assert.Equal(t, OrderiD, order.Id)
	assert.Equal(t, CariD, order.Car.ID)
	assert.Equal(t, CustomeriD, order.Customer.ID)
}

func TestGetAllOrder(t *testing.T) {
	orderRepo := NewOrderRepo(db, log)
	customerRepo := NewCustomerRepo(db, log)
	carRepo := NewCarRepo(db, log)

	for i := 0; i < 5; i++ {
		reqCustomer := models.CreateCustomer{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Email:     faker.Email(),
			Phone:     faker.Phonenumber(),
			Address:   gofakeit.RandomString([]string{"Tashkent", "Samarkand", "Fergana"}),
		}

		Customerid, err := customerRepo.Create(context.Background(), reqCustomer)
		assert.NoError(t, err)

		reqCar := models.CreateCarRequest{
			Name:       "Chenger",
			Year:       2010,
			Brand:      faker.Word(),
			Model:      faker.Word(),
			HorsePower: 200,
			Colour:     "Blue",
			EngineCap:  2.0,
		}

		Carid, err := carRepo.Create(context.Background(), reqCar)
		assert.NoError(t, err)

		reqOrder := models.CreateOrder{
			CarId:      Carid,
			CustomerId: Customerid,
			FromDate:   time.Now().Format(time.RFC3339),
			ToDate:     time.Now().AddDate(0, 1, 5).Format(time.RFC3339),
			Status:     "active",
			Paid:       false,
		}

		_, err = orderRepo.Create(context.Background(), reqOrder)
		assert.NoError(t, err)
	}

	testCases := []struct {
		name     string
		req      models.GetAllOrdersRequest
		expected int
	}{
		{"Get first page with limit 5", models.GetAllOrdersRequest{Search: "", Page: 1, Limit: 5}, 5},
		{"Get second page with limit 3", models.GetAllOrdersRequest{Search: "", Page: 2, Limit: 3}, 3},
		{"Search for orders with 'Chenger' status", models.GetAllOrdersRequest{Search: "Chenger", Page: 1, Limit: 1}, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			orders, err := orderRepo.GetAll(context.Background(), tc.req)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, len(orders.Orders))
		})
	}
}

func TestDeleteOrder(t *testing.T) {
	orderRepo := NewOrderRepo(db, log)

	orderID, err := orderRepo.Create(context.Background(), models.CreateOrder{
		CarId:      CariD,
		CustomerId: CustomeriD,
		FromDate:   time.Now().Format(time.RFC3339),
		ToDate:     time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
		Status:     "active",
		Paid:       false,
	})
	assert.NoError(t, err)

	err = orderRepo.Delete(context.Background(), orderID)
	assert.NoError(t, err)

	_, err = orderRepo.GetByID(context.Background(), orderID)
	assert.Error(t, err)
}
