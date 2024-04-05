package main

import (
	"backend_course/rent_car/api"
	"backend_course/rent_car/config"
	"backend_course/rent_car/pkg/logger"
	"backend_course/rent_car/service"
	"backend_course/rent_car/storage/postgres"
	"context"
	"fmt"

	_ "github.com/joho/godotenv"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.ServiceName)

	store, err := postgres.New(context.Background(), cfg, log)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	services := service.New(store, log)
	server := api.New(services, log)

	fmt.Println("programm is running on localhost:8080...")
	server.Run(":8080")

}
