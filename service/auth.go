package service

import (
	"backend_course/rent_car/api/models"
	"backend_course/rent_car/config"
	"backend_course/rent_car/pkg"
	"backend_course/rent_car/pkg/jwt"
	"backend_course/rent_car/pkg/logger"
	"backend_course/rent_car/pkg/password"
	"backend_course/rent_car/pkg/smtp"
	"backend_course/rent_car/storage"
	"context"
	"fmt"
	"time"
)

type authService struct {
	storage storage.IStorage
	log     logger.ILogger
	redis   storage.IRedisStorage
}

func NewAuthService(storage storage.IStorage, log logger.ILogger, redis storage.IRedisStorage) authService {
	return authService{
		storage: storage,
		log:     log,
		redis:   redis,
	}
}

func (a authService) CustomerLogin(ctx context.Context, loginRequest models.CustomerLoginRequest) (models.CustomerLoginResponse, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.Login)
	customer, err := a.storage.Customer().GetByLogin(ctx, loginRequest.Login)
	if err != nil {
		a.log.Error("error while getting customer credentials by login", logger.Error(err))
		return models.CustomerLoginResponse{}, err
	}

	if err = password.CompareHashAndPassword(customer.Password, loginRequest.Password); err != nil {
		a.log.Error("error while comparing password", logger.Error(err))
		return models.CustomerLoginResponse{}, err
	}

	m := make(map[interface{}]interface{})

	m["user_id"] = customer.ID
	m["user_role"] = config.CUSTOMER_ROLE

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for customer login", logger.Error(err))
		return models.CustomerLoginResponse{}, err
	}

	return models.CustomerLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a authService) CustomerRegister(ctx context.Context, loginRequest models.CustomerRegisterRequest) error {
	fmt.Println(" loginRequest.Login: ", loginRequest.Mail)

	otpCode := pkg.GenerateOTP()

	msg := fmt.Sprintf("Your otp code is: %v, for registering RENT_CAR. Don't give it to anyone", otpCode)

	err := a.redis.SetX(ctx, loginRequest.Mail, otpCode, time.Minute*2)
	if err != nil {
		a.log.Error("error while setting otpCode to redis customer register", logger.Error(err))
		return err
	}

	err = smtp.SendMail(loginRequest.Mail, msg)
	if err != nil {
		a.log.Error("error while sending otp code to customer register", logger.Error(err))
		return err
	}
	return nil
}
