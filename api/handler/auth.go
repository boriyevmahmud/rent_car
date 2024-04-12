package handler

import (
	"backend_course/rent_car/api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CustomerLogin godoc
// @Router       /customer/login [POST]
// @Summary      Customer login
// @Description  Customer login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.CustomerLoginRequest true "login"
// @Success      201  {object}  models.CustomerLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) CustomerLogin(c *gin.Context) {
	loginReq := models.CustomerLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponseLog(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate login & password

	loginResp, err := h.Services.Auth().CustomerLogin(c.Request.Context(), loginReq)
	if err != nil {
		handleResponseLog(c, h.Log, "unauthorized", http.StatusUnauthorized, err)
		return
	}

	handleResponseLog(c, h.Log, "Succes", http.StatusOK, loginResp)

}

// CustomerRegister godoc
// @Router       /customer/register [POST]
// @Summary      Customer register
// @Description  Customer register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        register body models.CustomerRegisterRequest true "register"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) CustomerRegister(c *gin.Context) {
	loginReq := models.CustomerRegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponseLog(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate for (gmail.com or mail.ru) & check if email is not exists

	err := h.Services.Auth().CustomerRegister(c.Request.Context(), loginReq)
	if err != nil {
		handleResponseLog(c, h.Log, "", http.StatusInternalServerError, err)
		return
	}

	handleResponseLog(c, h.Log, "Otp sent successfull", http.StatusOK, "")
}

// CustomerLogin godoc
// @Router       /customer/login [POST]
// @Summary      Customer login
// @Description  Customer login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.CustomerLoginRequest true "login"
// @Success      201  {object}  models.CustomerLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) CustomerRegisterc(c *gin.Context) {
	loginReq := models.CustomerLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponseLog(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate login & password

	loginResp, err := h.Services.Auth().CustomerLogin(c.Request.Context(), loginReq)
	if err != nil {
		handleResponseLog(c, h.Log, "unauthorized", http.StatusUnauthorized, err)
		return
	}

	handleResponseLog(c, h.Log, "Succes", http.StatusOK, loginResp)

}
