package models

type GetOrder struct {
	Id        string   `json:"id"`
	Car       Car      `json:"car"`
	Customer  Customer `json:"customer"`
	FromDate  string   `json:"from_date"`
	ToDate    string   `json:"to_date"`
	Status    string   `json:"status"`
	Paid      bool     `json:"payment_status"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type CreateOrder struct {
	Id         string `json:"id"`
	CarId      string `json:"car_id"`
	CustomerId string `json:"customer_id"`
	FromDate   string `json:"from_date"`
	ToDate     string `json:"to_date"`
	Status     string `json:"status"`
	Paid       bool   `json:"payment_status"`
}

type GetAllOrders struct {
	Orders []GetOrder `json:"orders"`
	Count  int        `json:"count"`
}
