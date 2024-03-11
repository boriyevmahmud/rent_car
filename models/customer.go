package models

type Customer struct {
	Id         string     `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Gmail      string     `json:"gmail"`
	Phone      string     `json:"phone"`
	Is_Blocked bool       `json:"isblocked"`
	CreatedAt  string     `json:"createdAt,omitempty"`
	UpdatedAt  string     `json:"updatedAt"`
	Orders     []GetOrder `json:"orders"`
}

type GetAllCustomersResponse struct {
	Customers []Customer `json:"customers"`
	Count     int16      `json:"count"`
}
