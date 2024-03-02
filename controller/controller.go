package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rent-car/config"
	"rent-car/models"
	"rent-car/storage"
)

type Controller struct {
	Store storage.Store
}

func NewController(store storage.Store) Controller {
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
