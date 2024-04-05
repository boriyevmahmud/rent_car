package models

type Car struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Year       int64   `json:"year"`
	Brand      string  `json:"brand"`
	Model      string  `json:"model"`
	HorsePower int64   `json:"horse_power"`
	Colour     string  `json:"colour"`
	EngineCap  float32 `json:"engine_cap"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type GetCar struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

type CreateCarRequest struct {
	Name       string  `json:"name"`
	Year       int64   `json:"year"`
	Brand      string  `json:"brand"`
	Model      string  `json:"model"`
	HorsePower int64   `json:"horse_power"`
	Colour     string  `json:"colour"`
	EngineCap  float32 `json:"engine_cap"`
}

type UpdateCarRequest struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Year       int64   `json:"year"`
	Brand      string  `json:"brand"`
	Model      string  `json:"model"`
	HorsePower int64   `json:"horse_power"`
	Colour     string  `json:"colour"`
	EngineCap  float32 `json:"engine_cap"`
}

type GetCarByIDResponse struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Year       int64   `json:"year"`
	Brand      string  `json:"brand"`
	Model      string  `json:"model"`
	HorsePower int64   `json:"horse_power"`
	Colour     string  `json:"colour"`
	EngineCap  float32 `json:"engine_cap"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	Orders     []Car   `json:"orders"`
}

type GetAllCarsRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllCarsResponse struct {
	Cars  []Car `json:"cars"`
	Count int64 `json:"count"`
}

type GetAvailableCarsRequest struct {
	Search string `json:"search"`
	Page   uint64    `json:"page"`
	Limit  uint64    `json:"limit"`
}

type GetAvailableCarsResponse struct {
	Cars  []Car `json:"cars"`
	Count uint64   `json:"count"`
}
