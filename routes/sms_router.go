package routes

import (
	"github.com/gin-gonic/gin"
	controller_implementation "github.com/lab-smart/sms-gateway/controllers"
	"github.com/lab-smart/sms-gateway/interfaces/controllers"
	services2 "github.com/lab-smart/sms-gateway/interfaces/services"
	"github.com/lab-smart/sms-gateway/middlewares/sms"
	"github.com/lab-smart/sms-gateway/repository"
	"github.com/lab-smart/sms-gateway/services"
	"gorm.io/gorm"
)

func ConfigureSMSService(db *gorm.DB) services2.SmsService {
	smsRepository := repository.NewSmsRepository(db)
	smsService := services.NewSmsService(smsRepository)
	return smsService
}

func ConfigureSMSController(postgresConnection *gorm.DB) controllers.SMSController {
	smsService := ConfigureSMSService(postgresConnection)
	smsController := controller_implementation.NewSMSController(smsService)
	return smsController
}

func configureSMSVersionOne(group *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	v1 := group.Group("/v1")
	smsController := ConfigureSMSController(db)
	smsService := ConfigureSMSService(db)
	v1.Use(sms.ApiKeyValidator(smsService))
	v1.GET("/send", smsController.SendSMS)
	return v1
}

func ConfigureSmSRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {
	smsGroup := r.Group("/sms")
	// introduce the middleware for api key validation
	configureSMSVersionOne(smsGroup, db)
	return r
}
