package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sms-gateway/utils"
	"time"
)

type BusinessEntity struct {
	Id uuid.UUID `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	ApiKey string `json:"api_key"`
	CreationTime time.Time `json:"creation_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (b *BusinessEntity) BeforeCreate(tx *gorm.DB) (err error)  {
	b.Id = uuid.New()
	b.CreationTime = time.Now()
	b.UpdateTime = time.Now()
	b.ApiKey = utils.GenerateRandomString(20)
	return
}

func (b *BusinessEntity) BeforeSave(tx *gorm.DB) (err error)  {
	b.UpdateTime = time.Now()
	return
}
