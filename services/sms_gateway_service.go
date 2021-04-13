package services

import (
	"sms-gateway/interfaces"
	"sms-gateway/models"
)

type smsService struct {
	smsRepo interfaces.SmsRepository
}

func (s *smsService) FindLeastUsedMessageTemplate() *models.MessageTemplate {
	return s.smsRepo.FindLeastUsedMessageTemplate()
}

func (s *smsService) CreateMessageTemplate(message string, category string) *models.MessageTemplate {
	model := models.MessageTemplate{
		Template: message,
		Count:    0,
		Category: category,
	}
	savedModel := s.smsRepo.CreateMessageTemplate(model)
	return savedModel
}

func NewSmsService(repository interfaces.SmsRepository) interfaces.SmsService {
	return &smsService{
		smsRepo: repository,
	}
}