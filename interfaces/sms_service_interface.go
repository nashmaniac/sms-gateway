package interfaces

import "sms-gateway/models"

type SmsService interface {
	CreateMessageTemplate(message string, category string) *models.MessageTemplate
	FindLeastUsedMessageTemplate() *models.MessageTemplate
}