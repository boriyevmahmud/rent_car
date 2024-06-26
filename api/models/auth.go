package models

type CustomerLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CustomerLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthInfo struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
}

type CustomerRegisterRequest struct {
	Mail string `json:"mail"`
}

type CustomerRegisterConfRequest struct {
	Mail     string         `json:"mail"`
	Otp      string         `json:"otp"`
	Customer CreateCustomer `json:"customer"`
}
