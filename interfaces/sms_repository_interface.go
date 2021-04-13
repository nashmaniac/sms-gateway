package interfaces

import "sms-gateway/models"

type SmsRepository interface {
	CreateMessageTemplate(m models.MessageTemplate) *models.MessageTemplate
	FindLeastUsedMessageTemplate() *models.MessageTemplate
}
