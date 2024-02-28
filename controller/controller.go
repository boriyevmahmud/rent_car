package controller

import (
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
