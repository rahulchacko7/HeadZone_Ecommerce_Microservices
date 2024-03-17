package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	interfaces "rahulchacko7/GO-GRPC-API-GATEWAY/pkg/client/interface"
	"rahulchacko7/GO-GRPC-API-GATEWAY/pkg/utils/models"
	"rahulchacko7/GO-GRPC-API-GATEWAY/pkg/utils/response"
)

type AdminHandler struct {
	GRPC_Client interfaces.AdminClient
}

func NewAdminHandler(adminClient interfaces.AdminClient) *AdminHandler {
	return &AdminHandler{
		GRPC_Client: adminClient,
	}
}

func (ad *AdminHandler) LoginHandler(c *gin.Context) {
	var adminDetails models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	admin, err := ad.GRPC_Client.AdminLogin(adminDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, success)
}

func (ad *AdminHandler) AdminSignUp(c *gin.Context) {
	var adminDetails models.AdminSignUp
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	admin, err := ad.GRPC_Client.AdminSignUp(adminDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Admin created successfully", admin, nil)
	c.JSON(http.StatusOK, success)
}
