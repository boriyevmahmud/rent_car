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

func TestCreateCustomer(t *testing.T) {
	customerRepo := NewCustomerRepo(db, log)

	reqCustomer := models.CreateCustomer{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Phone:     faker.Phonenumber(),
		Password:  faker.Password(),
		Address:   gofakeit.RandomString([]string{"Tashkent", "Samarkand", "Fergana"}),
	}

	id, err := customerRepo.Create(context.Background(), reqCustomer)
	assert.NoError(t, err)

	createdCustomer, err := customerRepo.GetByID(context.Background(), id)
	assert.NoError(t, err)

	assert.Equal(t, reqCustomer.FirstName, createdCustomer.FirstName)
	assert.Equal(t, reqCustomer.LastName, createdCustomer.LastName)
	assert.Equal(t, reqCustomer.Email, createdCustomer.Email)
	assert.Equal(t, reqCustomer.Phone, createdCustomer.Phone)
	assert.Equal(t, reqCustomer.Address, createdCustomer.Address)
}

func TestUpdateCustomer(t *testing.T) {
	customerRepo := NewCustomerRepo(db, log)

	customerID, err := customerRepo.Create(context.Background(), models.CreateCustomer{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Phone:     faker.Phonenumber(),
		Address:   gofakeit.RandomString([]string{"Tashkent", "Samarkand", "Fergana"}),
	})
	assert.NoError(t, err)

	updateCustomer := models.UpdateCustomer{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Phone:     faker.Phonenumber(),
		Address:   gofakeit.RandomString([]string{"Tashkent", "Samarkand", "Fergana"}),
	}

	updatedID, err := customerRepo.Update(context.Background(), updateCustomer, customerID)
	assert.NoError(t, err)
	assert.Equal(t, customerID, updatedID)

	updatedCustomer, err := customerRepo.GetByID(context.Background(), customerID)
	assert.NoError(t, err)

	assert.Equal(t, updateCustomer.FirstName, updatedCustomer.FirstName)
	assert.Equal(t, updateCustomer.LastName, updatedCustomer.LastName)
	assert.Equal(t, updateCustomer.Email, updatedCustomer.Email)
	assert.Equal(t, updateCustomer.Phone, updatedCustomer.Phone)
	assert.Equal(t, updateCustomer.Address, updatedCustomer.Address)
}

func TestGetByIDCustomer(t *testing.T) {
	customerRepo := NewCustomerRepo(db, log)

	expectedCustomer := models.CreateCustomer{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Phone:     faker.Phonenumber(),
		Address:   gofakeit.RandomString([]string{"Tashkent", "Samarkand", "Fergana"}),
	}

	customerID, err := customerRepo.Create(context.Background(), expectedCustomer)
	assert.NoError(t, err)

	customer, err := customerRepo.GetByID(context.Background(), customerID)
	assert.NoError(t, err)

	assert.Equal(t, expectedCustomer.FirstName, customer.FirstName)
	assert.Equal(t, expectedCustomer.LastName, customer.LastName)
	assert.Equal(t, expectedCustomer.Email, customer.Email)
	assert.Equal(t, expectedCustomer.Phone, customer.Phone)
	assert.Equal(t, expectedCustomer.Address, customer.Address)
}

func TestGetAllCustomer(t *testing.T) {
	customerRepo := NewCustomerRepo(db, log)

	for i := 0; i < 3; i++ {
		reqCustomer := models.CreateCustomer{
			FirstName: "Yulduz",
			LastName:  faker.LastName(),
			Email:     faker.Email(),
			Phone:     faker.Phonenumber(),
			Address:   gofakeit.RandomString([]string{"Tashkent", "Samarkand", "Fergana"}),
		}
		_, err := customerRepo.Create(context.Background(), reqCustomer)
		assert.NoError(t, err)
	}

	testCases := []struct {
		name     string
		req      models.GetAllCustomersRequest
		expected int
	}{
		{"Get first page with limit 3", models.GetAllCustomersRequest{Page: 1, Limit: 2}, 2},
		{"Get second page with limit 3", models.GetAllCustomersRequest{Page: 2, Limit: 1}, 1},
		{"Search for customers with first name 'Yulduz'", models.GetAllCustomersRequest{Search: "Yulduz", Page: 1, Limit: 3}, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			customers, err := customerRepo.GetAll(context.Background(), tc.req)
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, len(customers.Customers))
			assert.NoError(t, err)
		})
	}
}

func TestGetCustomerCars(t *testing.T) {
	orderRepo := NewOrderRepo(db, log)
	carRepo := NewCarRepo(db, log)
	customerRepo := NewCustomerRepo(db, log)

	reqCustomer := models.CreateCustomer{
		FirstName: "Bilmasam",
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Phone:     faker.Phonenumber(),
		Address:   gofakeit.RandomString([]string{"Tashkent", "Samarkand", "Fergana"}),
	}

	Customerid, err := customerRepo.Create(context.Background(), reqCustomer)
	assert.NoError(t, err)

	reqCar := models.CreateCarRequest{
		Name:       "Nimadirde",
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
		ToDate:     time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
		Status:     "active",
		Paid:       true,
	}

	orderId, err := orderRepo.Create(context.Background(), reqOrder)
	assert.NoError(t, err)

	testCases := []struct {
		name     string
		carName  string
		id       string
		boolean  bool
		expected int
	}{
		{
			name:     "Test case 1: Get customer cars by customer ID",
			carName:  "",
			id:       Customerid,
			boolean:  true,
			expected: 1,
		},
		{
			name:     "Test case 2: Get customer cars by car name",
			carName:  "Nimadirde",
			id:       Customerid,
			boolean:  false,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			customerCars, err := customerRepo.GetCustomerCars(context.Background(), tc.carName, tc.id, tc.boolean)
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, len(customerCars.CustomerCars))
			assert.NoError(t, err)
		})
	}

	err = orderRepo.DeleteHard(context.Background(), orderId)
	assert.NoError(t, err)
	err = carRepo.DeleteHard(context.Background(), Carid)
	assert.NoError(t, err)
	err = customerRepo.DeleteHard(context.Background(), Customerid)
	assert.NoError(t, err)
}

func TestDeleteCustomer(t *testing.T) {
	customerRepo := NewCustomerRepo(db, log)

	customerID, err := customerRepo.Create(context.Background(), models.CreateCustomer{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Phone:     faker.Phonenumber(),
		Address:   gofakeit.RandomString([]string{"Tashkent", "Samarkand", "Fergana"}),
	})
	assert.NoError(t, err)

	err = customerRepo.Delete(context.Background(), customerID)
	assert.NoError(t, err)

	_, err = customerRepo.GetByID(context.Background(), customerID)
	assert.Error(t, err)
}
