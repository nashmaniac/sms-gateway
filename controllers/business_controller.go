package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sms-gateway/dto/business"
	"sms-gateway/dto/errors"
	"sms-gateway/interfaces/controllers"
	"sms-gateway/interfaces/services"
)

type businessController struct {
	smsService services.SmsService
}

func (b *businessController) CreateBusinessEntity(c *gin.Context) {
	var params business.CreateBusinessParams

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Message: "Invalid json body",
			Code:    http.StatusBadRequest,
		})
		return
	}
	bEntity := b.smsService.CreateBusinessEntity(params.Name)
	c.JSON(http.StatusOK, bEntity.ToBusinessResponse())
}

func NewBusinessController(service services.SmsService) controllers.BusinessController {
	return &businessController{
		smsService: service,
	}
}