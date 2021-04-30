package sms

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sms-gateway/dto/errors"
	"sms-gateway/interfaces/services"
)

func ApiKeyValidator(service services.SmsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey, exists := c.GetQuery("apiKey")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, errors.ErrorResponse{
				Message: "Api Key not present",
				Code:    http.StatusForbidden,
			})
			c.Abort()
			return
		}
		b, err := service.GetBusinessEntityByApiKey(apiKey)
		if err != nil || b == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrorResponse{
				Message: "Business entity not found",
				Code:    http.StatusNotFound,
			})
			c.Abort()
			return
		}
		log.Println(fmt.Sprintln("Request received from business entity ", b.Name))
		c.Next()
	}
}