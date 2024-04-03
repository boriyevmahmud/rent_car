package handler

import (
	"fmt"
	"net/http"
	_ "rent-car/api/docs"
	"rent-car/api/models"
	"rent-car/pkg/check"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security ApiKeyAuth
// @Router 		/car [POST]
// @Summary 	create a car
// @Description This api is creates a new car and returns it's id
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param		car body models.Car true "car"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateCar(c *gin.Context) {
	car := models.Car{}
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&car); err != nil {
		handleResponseLog(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	if err := check.ValidateCarYear(car.Year); err != nil {
		handleResponseLog(c, h.Log, "error while validating car year, year: "+strconv.Itoa(car.Year), http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.Services.Car().Create(ctx, car)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating car", http.StatusBadRequest, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Created successfully", http.StatusOK, id)
}

func (h Handler) UpdateCar(c *gin.Context) {
	car := models.Car{}

	if err := c.ShouldBindJSON(&car); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	if err := check.ValidateCarYear(car.Year); err != nil {
		handleResponse(c, "error while validating car year, year: "+strconv.Itoa(car.Year), http.StatusBadRequest, err.Error())
		return
	}
	car.Id = c.Param("id")

	err := uuid.Validate(car.Id)
	if err != nil {
		handleResponse(c, "error while validating car id,id: "+car.Id, http.StatusBadRequest, err.Error())
		return
	}

	// id, err := h.Store.Car().Update(context.Background(), car)
	// if err != nil {
	// 	handleResponse(c, "error while updating car", http.StatusBadRequest, err.Error())
	// 	return
	// }

	handleResponse(c, "Updated successfully", http.StatusOK, "id")
}

// @Security ApiKeyAuth
// @Router 		/car [GET]
// @Summary 	get a car
// @Description This api is get a car
// @Tags 		car
// @Accept		json
// @Produce		json
// @Success		200  {object}  models.GetAllCarsResponse
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllCars(c *gin.Context) {
	var (
		request = models.GetAllCarsRequest{}
	)

	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit
	// cars, err := h.Services.Car().GetAll(request)
	// if err != nil {
	// 	handleResponse(c, "error while gettign cars", http.StatusBadRequest, err.Error())

	// 	return
	// }

	// handleResponse(c, "", http.StatusOK, cars)
}

func (h Handler) DeleteCar(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	// err = h.Store.Car().Delete(id)
	// if err != nil {
	// 	handleResponse(c, "error while deleting car", http.StatusInternalServerError, err.Error())
	// 	return
	// }

	handleResponse(c, "", http.StatusOK, id)
}

// @Security ApiKeyAuth
// @Router 		/car/{id} [GET]
// @Summary 	get a car
// @Description This api is get a car
// @Tags 		car
// @Accept		json
// @Produce		json
// @Param 		id path string true "id"
// @Success		200  {object}  models.Car
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetCar(c *gin.Context) {

	id := c.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	cars, err := h.Services.Car().Get(c.Request.Context(), id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while gett car", http.StatusBadRequest, err.Error())

		return
	}

	handleResponseLog(c, h.Log, "", http.StatusOK, cars)
}
