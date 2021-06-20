package repository

import (
	"github.com/lab-smart/sms-gateway/interfaces/repository"
	"github.com/lab-smart/sms-gateway/models"
	"gorm.io/gorm"
)

type smsRepository struct {
	db *gorm.DB
}

func (s *smsRepository) GetDB() *gorm.DB {
	return s.db
}

func (s *smsRepository) CreateMessage(message models.Message) *models.Message {
	s.db.Create(&message)
	return &message
}

func (s *smsRepository) FindLeastUsedSender() *models.Sender {
	var model models.Sender
	s.db.Find(&models.Sender{}).Order("count").First(&model)
	return &model
}

func (s *smsRepository) FindBusinessEntityByApiKey(apiKey string) *models.BusinessEntity {
	var model models.BusinessEntity
	err := s.db.Where("api_key = ?", apiKey).First(&model).Error
	if err != nil {
		return nil
	}
	return &model
}

func (s *smsRepository) CreateBusinessEntity(entity models.BusinessEntity) *models.BusinessEntity {
	s.db.Create(&entity)
	return &entity
}

func (s *smsRepository) CreateSender(sender models.Sender) *models.Sender {
	s.db.Create(&sender)
	return &sender
}

func (s *smsRepository) FindLeastUsedMessageTemplate() *models.MessageTemplate {
	var messageTemplate models.MessageTemplate
	s.db.Find(&models.MessageTemplate{}).Order("count").First(&messageTemplate)
	return &messageTemplate
}

func (s *smsRepository) CreateMessageTemplate(m models.MessageTemplate) *models.MessageTemplate {
	s.db.Create(&m)
	return &m
}

func NewSmsRepository(db *gorm.DB) repository.SmsRepository {
	return &smsRepository{
		db: db,
	}
}
