package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rent-car/models"
	"rent-car/pkg/check"

	"github.com/google/uuid"
)

func (c Controller) Car(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateCar(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllCars(w, r)
		} else {
			// get one
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.UpdateCar(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteCar(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (c Controller) CreateCar(w http.ResponseWriter, r *http.Request) {
	car := models.Car{}

	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}
	if err := check.ValidateCarYear(car.Year); err != nil {
		fmt.Println("error while validating year: ", car.Year)
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	id, err := c.Store.Car.Create(car)
	if err != nil {
		fmt.Println("error while creating car, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) UpdateCar(w http.ResponseWriter, r *http.Request) {
	car := models.Car{}

	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}

	if err := check.ValidateCarYear(car.Year); err != nil {
		fmt.Println("error while validating year: ", car.Year)
		handleResponse(w, http.StatusBadRequest, err)
		return
	}
	car.Id = r.URL.Query().Get("id")

	err := uuid.Validate(car.Id)
	if err != nil {
		fmt.Println("error while validating, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := c.Store.Car.Update(car)
	if err != nil {
		fmt.Println("error while creating car, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetAllCars(w http.ResponseWriter, r *http.Request) {
	var (
		values = r.URL.Query()
		search string
	)
	if _, ok := values["search"]; ok {
		search = values["search"][0]
	}

	cars, err := c.Store.Car.GetAllCars(search)
	if err != nil {
		fmt.Println("error while getting cars, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, cars)
}

func (c Controller) DeleteCar(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = c.Store.Car.Delete(id)
	if err != nil {
		fmt.Println("error while deleting car, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}
