package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lab-smart/sms-gateway/dto/errors"
	"github.com/lab-smart/sms-gateway/interfaces/controllers"
	"github.com/lab-smart/sms-gateway/interfaces/services"
	"github.com/lab-smart/sms-gateway/utils"
)

type smsController struct {
	smsService services.SmsService
}

func (s *smsController) SendSMS(c *gin.Context) {
	apiKey, err := c.GetQuery("apiKey")
	if !err {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrorResponse{
			Message: "Api key is mssing",
			Code:    http.StatusBadRequest,
		})
		c.Abort()
		return
	}
	toNum, err := c.GetQuery("to")
	if !err {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrorResponse{
			Message: "To number is mssing",
			Code:    http.StatusBadRequest,
		})
		c.Abort()
		return
	}
	pin, err := c.GetQuery("pin")
	if !err {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrorResponse{
			Message: "Pin is mssing",
			Code:    http.StatusBadRequest,
		})
		c.Abort()
		return
	}
	source, err := c.GetQuery("source")
	if !err {
		source = "en"
	}
	destinaton := utils.DetectDestinationBasedOnNumber(toNum)

	msg, er := s.smsService.SendTextMessage(apiKey, pin, toNum, source, *destinaton, true)
	if er != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrorResponse{
			Message: er.Error(),
			Code:    http.StatusBadRequest,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, msg)
}

func NewSMSController(service services.SmsService) controllers.SMSController {
	return &smsController{
		smsService: service,
	}
}
