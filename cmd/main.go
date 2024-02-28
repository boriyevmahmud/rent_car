package main

import (
	"fmt"
	"rent-car/config"
	"rent-car/controller"
	"rent-car/storage"
)

func main() {
	cfg := config.Load()
	store, err := storage.New(cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.DB.Close()

	c := controller.NewController(store)
	c.CreateCar()

}
