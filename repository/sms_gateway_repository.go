package repository

import (
	"gorm.io/gorm"
	"sms-gateway/db"
	"sms-gateway/interfaces"
	"sms-gateway/models"
)

type smsRepository struct {
	db *gorm.DB
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

func NewSmsRepository() interfaces.SmsRepository {
	database := db.GetPostgresConnection()
	return &smsRepository{
		db: database,
	}
}
