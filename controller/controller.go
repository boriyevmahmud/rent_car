package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rent-car/config"
	"rent-car/models"
	"rent-car/storage"
	"strconv"
)

type Controller struct {
	Store storage.IStorage
}

func NewController(store storage.IStorage) Controller {
	return Controller{
		Store: store,
	}
}

func handleResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	resp := models.Response{}

	if statusCode >= 100 && statusCode <= 199 {
		resp.Description = config.ERR_INFORMATION
	} else if statusCode >= 200 && statusCode <= 299 {
		resp.Description = config.SUCCESS
	} else if statusCode >= 300 && statusCode <= 399 {
		resp.Description = config.ERR_REDIRECTION
	} else if statusCode >= 400 && statusCode <= 499 {
		resp.Description = config.ERR_BADREQUEST
	} else {
		resp.Description = config.ERR_INTERNAL_SERVER
	}
	resp.StatusCode = statusCode
	resp.Data = data

	js, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("error while marshalling: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(js)
}

func ParsePageQueryParam(r *http.Request) (uint64, error) {
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.ParseUint(pageStr, 10, 30)
	if err != nil {
		return 0, err
	}
	//offset: page - 1 * limit = 0
	//limit: limit = 10 
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

func ParseLimitQueryParam(r *http.Request) (uint64, error) {
	limitStr := r.URL.Query().Get("limit")
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
