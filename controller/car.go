package controller

import (
	"fmt"
	"rent-car/models"
	"time"
)

func (c *Controller) CreateCar() {
	car := getCarInfo()
	if car.Year <= 0 || car.Year > time.Now().Year()+1 {
		fmt.Println("year intput is not correct")
		return
	}
	id, err := c.Store.Car.Create(car)
	if err != nil {
		fmt.Println("error while creating car, err: ", err)
		return
	}
	fmt.Printf("Car created successfully with ID: %v\n", id)

}

func getCarInfo() models.Car {
	car := models.Car{}
	fmt.Println(`Enter the car datas 
	Name
	Year
	Brand
	Model
	HoursePower
	Colour
	EngineCap`)
	// var (
	// 	Name, Brand, Model, Colour string
	// 	Year, HoursePower          int
	// 	EngineCap                  float32
	// )
	fmt.Scan(&car.Name, &car.Year, &car.Brand, &car.Model, &car.HoursePower, &car.Colour, &car.EngineCap)

	return car
}
