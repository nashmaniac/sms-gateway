package services

import (
	"errors"
	"sms-gateway/interfaces/repository"
	"sms-gateway/interfaces/services"
	"sms-gateway/models"
	"sms-gateway/utils"
	"time"
)

type smsService struct {
	smsRepo repository.SmsRepository
}

func (s *smsService) SendTextMessage(apiKey string, pin string, to string, source string, dest string, conversion bool) (*models.Message, error) {
	businessModel := s.smsRepo.FindBusinessEntityByApiKey(apiKey)
	if businessModel == nil {
		return nil, errors.New("business entity with apiKey not present")
	}
	fromObj := s.smsRepo.FindLeastUsedSender()
	messageTemplate := s.smsRepo.FindLeastUsedMessageTemplate()
	var messageToSend string
	if conversion {
		converter := utils.CodeConverter{Code: pin, Source: source, Destination: dest}
		messageToSend = messageTemplate.OutputFormattedMessage(*converter.ConvertMessage())
	} else {
		messageToSend = pin
	}
	message := s.CreateMessage(*businessModel, *fromObj, to, messageToSend)
	// increment from and template count
	fromObj.Count++
	s.smsRepo.GetDB().Save(fromObj)
	messageTemplate.Count++
	s.smsRepo.GetDB().Save(messageTemplate)
	dispatcher := utils.GetMessageDispatcher(message.Id, message.FromNumber, message.ToNumber, message.Content)
	response := dispatcher.Send()
	message.SendingTime = time.Now()
	message.IsSent = true
	if response.IsSuccess {
		message.IsSuccessful = true
		message.ResponseId = response.ResponseId
	} else {
		message.IsSuccessful = false
		message.Reason = response.ErrorText
	}
	s.smsRepo.GetDB().Save(message)
	return message, nil
}

func (s *smsService) CreateMessage(entity models.BusinessEntity, from models.Sender, to string, content string) *models.Message {
	m := models.Message{
		BusinessEntityId: entity.Id,
		FromNumber:       from.Msisdn,
		ToNumber:         to,
		Content:          content,
	}
	model := s.smsRepo.CreateMessage(m)
	return model
}

func (s *smsService) FindLeastUsedSender() *models.Sender {
	return s.smsRepo.FindLeastUsedSender()
}

func (s *smsService) GetBusinessEntityByApiKey(apiKey string) (*models.BusinessEntity, error) {
	model := s.smsRepo.FindBusinessEntityByApiKey(apiKey)
	if model == nil {
		return nil, errors.New("no business entity found with apiKey")
	}
	return model, nil
}

func (s *smsService) CreateBusinessEntity(entity string) *models.BusinessEntity {
	model := models.BusinessEntity{
		Name:         entity,
	}
	return s.smsRepo.CreateBusinessEntity(model)
}

func (s *smsService) CreateSender(sender string) *models.Sender {
	model := models.Sender{
		Msisdn:       sender,
	}
	savedModel := s.smsRepo.CreateSender(model)
	return savedModel
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

func NewSmsService(repository repository.SmsRepository) services.SmsService {
	return &smsService{
		smsRepo: repository,
	}
}