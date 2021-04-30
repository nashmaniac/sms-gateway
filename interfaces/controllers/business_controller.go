package controllers

import "github.com/gin-gonic/gin"

type BusinessController interface {
	CreateBusinessEntity(c *gin.Context)
}
