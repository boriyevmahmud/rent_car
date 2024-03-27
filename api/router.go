package api

import (
	"rent-car/api/handler"
	"rent-car/service"
	"rent-car/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func New(services service.IServiceManager, store storage.IStorage) *gin.Engine {
	h := handler.NewStrg(store, services)

	r := gin.Default()

	r.POST("/car", h.CreateCar)
	r.GET("/car", h.GetAllCars)
	r.PUT("/car/:id", h.UpdateCar)
	r.DELETE("/car/:id", h.DeleteCar)
	// r.PATCH("/car/:id", h.UpdateUserPassword)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
