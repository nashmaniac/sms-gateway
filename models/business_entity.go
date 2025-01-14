package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lab-smart/sms-gateway/dto/business"
	"github.com/lab-smart/sms-gateway/utils"
	"gorm.io/gorm"
)

type BusinessEntity struct {
	Id           uuid.UUID `json:"id" gorm:"primary_key"`
	Name         string    `json:"name"`
	ApiKey       string    `json:"api_key"`
	CreationTime time.Time `json:"creation_time"`
	UpdateTime   time.Time `json:"update_time"`
	Message      []Message
}

func (b *BusinessEntity) BeforeCreate(tx *gorm.DB) (err error) {
	b.Id = uuid.New()
	b.CreationTime = time.Now()
	b.UpdateTime = time.Now()
	b.ApiKey = utils.GenerateRandomString(20)
	return
}

func (b *BusinessEntity) BeforeSave(tx *gorm.DB) (err error) {
	b.UpdateTime = time.Now()
	return
}

func (b BusinessEntity) ToBusinessResponse() business.BusinessEntityResponse {
	return business.BusinessEntityResponse{
		Id:           b.Id,
		Name:         b.Name,
		ApiKey:       b.ApiKey,
		CreationTime: b.CreationTime,
		UpdateTime:   b.UpdateTime,
	}
}
