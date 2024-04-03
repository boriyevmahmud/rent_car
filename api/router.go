package api

import (
	"errors"
	"net/http"
	"rent-car/api/handler"
	"rent-car/pkg/logger"
	"rent-car/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(services service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.NewStrg(services, log)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(authMiddleware)

	r.POST("/car", h.CreateCar)
	r.GET("/car", h.GetAllCars)
	r.PUT("/car/:id", h.UpdateCar)
	r.DELETE("/car/:id", h.DeleteCar)
	r.GET("/car/:id", h.GetCar)
	// r.PATCH("/car/:id", h.UpdateUserPassword)

	return r
}

func authMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
	}
	c.Next()
}
