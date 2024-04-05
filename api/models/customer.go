package models

type GetCustomer struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

type Customer struct {
	ID              string  `json:"id"`
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	Email           string  `json:"email"`
	Phone           string  `json:"phone"`
	Address         string  `json:"address"`
	CreatedAt       string  `json:"created_at,omitempty"`
	UpdatedAt       string  `json:"updated_at"`
	Orders          []Order `json:"orders,omitempty"`
	OrdersCount     int64   `json:"orders_count"`
	UniqueCarsCount int64   `json:"unique_cars_count"`
}

type CreateCustomer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Address   string `json:"address"`
}

type LoginCustomer struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type ChangePassword struct {
	Login       string `json:"login"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UpdateCustomer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

type GetAllCustomersRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllCustomersResponse struct {
	Customers []Customer `json:"customers"`
	Count     int64      `json:"count"`
}

type GetCustomerCarsResponse struct {
	CustomerCars []GetCustomerCars `json:"customer_cars"`
	Count        int64             `json:"count"`
}

type GetCustomerCars struct {
	CarName        string  `json:"car_name"`
	OrderCreatedAt string  `json:"order_created_at"`
	Duration       float64 `json:"duration"`
	Price          float64 `json:"price"`
}
