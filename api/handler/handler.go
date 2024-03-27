package handler

import (
	"fmt"
	"rent-car/api/models"
	"rent-car/config"
	"rent-car/service"
	"rent-car/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Store    storage.IStorage
	Services service.IServiceManager
}

func NewStrg(store storage.IStorage, services service.IServiceManager) Handler {
	return Handler{
		Store:    store,
		Services: services,
	}
}

func handleResponse(c *gin.Context, msg string, statusCode int, data interface{}) {
	resp := models.Response{}

	if statusCode >= 100 && statusCode <= 199 {
		resp.Description = config.ERR_INFORMATION
	} else if statusCode >= 200 && statusCode <= 299 {
		resp.Description = config.SUCCESS
	} else if statusCode >= 300 && statusCode <= 399 {
		resp.Description = config.ERR_REDIRECTION
	} else if statusCode >= 400 && statusCode <= 499 {
		resp.Description = config.ERR_BADREQUEST
		fmt.Println("BAD REQUEST: "+msg, "reason: ", data)
	} else {
		resp.Description = config.ERR_INTERNAL_SERVER
		fmt.Println("INTERNAL SERVER ERROR: "+msg, "reason: ", data)
	}

	resp.StatusCode = statusCode
	resp.Data = data

	c.JSON(resp.StatusCode, resp)
}

func ParsePageQueryParam(c *gin.Context) (uint64, error) {
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.ParseUint(pageStr, 10, 30)
	if err != nil {
		return 0, err
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

func ParseLimitQueryParam(c *gin.Context) (uint64, error) {
	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.ParseUint(limitStr, 10, 30)
	if err != nil {
		return 0, err
	}

	if limit == 0 {
		return 10, nil
	}
	return limit, nil
}
