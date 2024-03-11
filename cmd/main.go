package main

import (
	"fmt"
	"net/http"
	"rent-car/config"
	"rent-car/controller"
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

	con := controller.NewController(store)

	http.HandleFunc("/car", con.Car)

	fmt.Println("programm is running on localhost:8008...")
	http.ListenAndServe(":8008", nil)

}
