package handler

import (
	"backend_course/rent_car/api/models"
	"backend_course/rent_car/config"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateOrder godoc
// @Security ApiKeyAuth
// @Router		/order [POST]
// @Summary		create an order
// @Description This api creates a new order and returns its id
// @Tags		order
// @Accept		json
// @Produce		json
// @Param		order body models.CreateOrder true "order"
// @Success		200  {string}  string
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateOrder(c *gin.Context) {
	var order models.CreateOrder

	data, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting auth", http.StatusUnauthorized, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	//TODO: need validate body

	order.Status = config.STATUS_NEW
	order.CustomerId = data.UserID

	id, err := h.Services.Order().Create(c.Request.Context(), order)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating order", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Order was successfully created", http.StatusOK, id)
}

// UpdateOrder godoc
// @Security ApiKeyAuth
// @Router		/order/{id} [PUT]
// @Summary		update an order
// @Description This api updates a order by its id and returns its id
// @Tags		order
// @Accept		json
// @Produce		json
// @Param		id path string true "order id"
// @Param		order body models.UpdateOrder true "order"
// @Success		200  {string}  string
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateOrder(c *gin.Context) {
	var order models.UpdateOrder

	if err := c.ShouldBindJSON(&order); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")
	order.Id = id

	if err := uuid.Validate(order.Id); err != nil {
		handleResponseLog(c, h.Log, "error while validating order ID", http.StatusBadRequest, err.Error())
		return
	}

	if _, err := h.Services.Order().Update(context.Background(), order); err != nil {
		handleResponseLog(c, h.Log, "error while updating order", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Order was successfully updated", http.StatusOK, id)
}

// GetOrderByID godoc
// @Security ApiKeyAuth
// @Router		/order/{id} [GET]
// @Summary		get an order by its id
// @Description This api gets a order by its id and returns its info
// @Tags		order
// @Accept		json
// @Produce		json
// @Param		id path string true "order"
// @Success		200  {object}  models.GetOrderResponse
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetOrderByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		handleResponseLog(c, h.Log, "missing order ID", http.StatusBadRequest, "")
		return
	}

	order, err := h.Services.Order().GetByID(context.Background(), id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting order by ID", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Order was successfully gotten by Id", http.StatusOK, order)
}

// GetAllOrders godoc
// @Security ApiKeyAuth
// @Router		/order [GET]
// @Summary		get all orders
// @Description This api gets all orders and returns their info
// @Tags		order
// @Accept		json
// @Produce		json
// @Param		order query string true "orders"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.GetAllOrdersResponse
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllOrders(c *gin.Context) {
	var (
		req = models.GetAllOrdersRequest{}
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

	orders, err := h.Services.Order().GetAll(context.Background(), req)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting all orders", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Orders were gotten successfully", http.StatusOK, orders)
}

// DeleteOrder godoc
// @Security ApiKeyAuth
// @Router		/order/{id} [DELETE]
// @Summary		delete an order by its id
// @Description This api deletes a order by its id
// @Tags		order
// @Accept		json
// @Produce		json
// @Param		id path string true "order id"
// @Success		200  {string}  nil
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		handleResponseLog(c, h.Log, "missing order ID", http.StatusBadRequest, id)
		return
	}

	if err := uuid.Validate(id); err != nil {
		handleResponseLog(c, h.Log, "error while validating order ID", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Services.Order().Delete(context.Background(), id); err != nil {
		handleResponseLog(c, h.Log, "error while deleting order", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "Order successfully deleted", http.StatusOK, "Order successfully deleted")
}
