package controllers

import "github.com/gin-gonic/gin"

type SMSController interface {
	SendSMS(c *gin.Context)
}
