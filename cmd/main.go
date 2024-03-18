package main

import (
	"fmt"
	"rent-car/api"
	"rent-car/config"
	"rent-car/storage/postgres"
)

func main() {
	cfg := config.Load()
	store, err := postgres.New(cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	c := api.New(store)

	fmt.Println("programm is running on localhost:8008...")
	c.Run(":8080")
}
