package handler

import (
	"backend_course/rent_car/api/models"
	"backend_course/rent_car/pkg/check"
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateCar godoc
// @Security ApiKeyAuth
// @Router		/car [POST]
// @Summary		create a car
// @Description This api creates a new car and returns its id
// @Tags		car
// @Accept		json
// @Produce		json
// @Param		car body models.CreateCarRequest true "car"
// @Success		201  {object}  string
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) CreateCar(c *gin.Context) {
	var carReq models.CreateCarRequest

	if err := c.ShouldBindJSON(&carReq); err != nil {
		handleResponseLog(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateCarYear(int(carReq.Year)); err != nil {
		handleResponseLog(c, h.Log, "error while validating car year, year: "+strconv.Itoa(int(carReq.Year)), http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.Services.Car().Create(c.Request.Context(), carReq)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating car", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Car was successfully created", http.StatusCreated, id)
}

// UpdateCar godoc
// @Security ApiKeyAuth
// @Router		/car/{id} [PUT]
// @Summary		update a car
// @Description This api updates a car by its id and returns its id
// @Tags		car
// @Accept		json
// @Produce		json
// @Param		car body models.UpdateCarRequest true "car"
// @Success		200  {object}  string
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) UpdateCar(c *gin.Context) {
	var carReq models.UpdateCarRequest

	if err := c.ShouldBindJSON(&carReq); err != nil {
		handleResponseLog(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := uuid.Validate(carReq.ID); err != nil {
		handleResponseLog(c, h.Log, "error while validating car ID", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.Car().Update(context.Background(), carReq)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating car", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Car was successfully updated", http.StatusOK, id)
}

// GetCarByID godoc
// @Security ApiKeyAuth
// @Router		/car/{id} [GET]
// @Summary		get a car by its id
// @Description This api gets a car by its id and returns its info
// @Tags		car
// @Accept		json
// @Produce		json
// @Param		id path string true "car"
// @Success		200  {object}  models.GetCarByIDResponse
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) GetCarByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		handleResponseLog(c, h.Log, "missing car ID", http.StatusBadRequest, id)
		return
	}

	car, err := h.Services.Car().GetByID(context.Background(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			handleResponseLog(c, h.Log, "car not found", http.StatusNotFound, err.Error())
			return
		}
		handleResponseLog(c, h.Log, "error while getting car by ID", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Car was successfully gotten by ID", http.StatusOK, car)
}

// GetAllCars godoc
// @Security ApiKeyAuth
// @Router		/car [GET]
// @Summary		get all cars
// @Description This api gets all cars and returns their info
// @Tags		car
// @Accept		json
// @Produce		json
// @Param		car query string true "cars"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.GetAllCarsResponse
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) GetAllCars(c *gin.Context) {
	var (
		req models.GetAllCarsRequest
	)

	req.Search = c.Query("search")

	page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}

	limit, err := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 64)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	req.Page = page
	req.Limit = limit

	cars, err := h.Services.Car().GetAll(context.Background(), req)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting cars", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Cars were successfully gotten", http.StatusOK, cars)
}

// GetAvailableCars godoc
// @Security ApiKeyAuth
// @Router		/car/available/ [GET]
// @Summary		get available cars
// @Description This api gets available cars and returns their info
// @Tags		car
// @Accept		json
// @Produce		json
// @Param		car query string true "cars"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.GetAvailableCarsResponse
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) GetAvailableCars(c *gin.Context) {
	var (
		req models.GetAvailableCarsRequest
	)

	req.Search = c.Query("search")

	page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}

	limit, err := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 64)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	req.Page = page
	req.Limit = limit

	cars, err := h.Services.Car().GetAvailable(context.Background(), req)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting available cars", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Available Cars were successfully gotten", http.StatusOK, cars)
}

// DeleteCar godoc
// @Security ApiKeyAuth
// @Router		/car/{id} [DELETE]
// @Summary		delete a car by its id
// @Description This api deletes a car by its id and returns error or nil
// @Tags		car
// @Accept		json
// @Produce		json
// @Param		id path string true "car"
// @Success		200  {object}  nil
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *Handler) DeleteCar(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		handleResponseLog(c, h.Log, "missing car ID", http.StatusBadRequest, id)
		return
	}

	if err := uuid.Validate(id); err != nil {
		handleResponseLog(c, h.Log, "error while validating car ID", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Services.Car().Delete(context.Background(), id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while deleting car", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Car was successfully deleted by ID", http.StatusOK, id)
}
