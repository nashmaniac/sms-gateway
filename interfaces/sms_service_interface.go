package interfaces

import (
	"sms-gateway/models"
)

type SmsService interface {
	CreateMessageTemplate(message string, category string) *models.MessageTemplate
	FindLeastUsedMessageTemplate() *models.MessageTemplate
	CreateSender(sender string) *models.Sender
	CreateBusinessEntity(entity string) *models.BusinessEntity
	GetBusinessEntityByApiKey(apiKey string) (*models.BusinessEntity, error)
	FindLeastUsedSender() *models.Sender
	CreateMessage(entity models.BusinessEntity, from models.Sender, to string, content string) *models.Message
	SendTextMessage(apiKey string, pin string, to string, source string, dest string, conversion bool) (*models.Message, error)
}