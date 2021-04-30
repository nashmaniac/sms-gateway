package repository

import (
	"gorm.io/gorm"
	"sms-gateway/models"
)

type SmsRepository interface {
	CreateMessageTemplate(m models.MessageTemplate) *models.MessageTemplate
	FindLeastUsedMessageTemplate() *models.MessageTemplate
	CreateSender(sender models.Sender) *models.Sender
	CreateBusinessEntity(entity models.BusinessEntity) *models.BusinessEntity
	FindBusinessEntityByApiKey(apiKey string) *models.BusinessEntity
	FindLeastUsedSender() *models.Sender
	CreateMessage(message models.Message) *models.Message
	GetDB() *gorm.DB
}
