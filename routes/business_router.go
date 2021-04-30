package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	controller_implementation "sms-gateway/controllers"
	"sms-gateway/interfaces/controllers"
	"sms-gateway/repository"
	"sms-gateway/services"
)

func ConfigureBusinessController(db *gorm.DB) controllers.BusinessController {
	smsRepo := repository.NewSmsRepository(db)
	smsService := services.NewSmsService(smsRepo)
	businessController := controller_implementation.NewBusinessController(smsService)
	return businessController
}

func configureBusinessVersionOne(group *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	v1 := group.Group("/v1")
	controller := ConfigureBusinessController(db)
	v1.POST("/create", controller.CreateBusinessEntity)
	return v1
}

func ConfigureBusinessRouter(r *gin.Engine, db *gorm.DB) *gin.Engine{
	businessGroup := r.Group("/business")
	configureBusinessVersionOne(businessGroup, db)
	return r
}
