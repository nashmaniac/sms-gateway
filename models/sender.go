package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Sender struct {
	Id uuid.UUID `json:"id" gorm:"primary_key"`
	Msisdn string `json:"msisdn"`
	Count int64 `json:"count" gorm:"default:0;"`
	IsActive bool `json:"is_active" gorm:"default:true;"`
	CreationTime time.Time `json:"creation_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (s *Sender) BeforeCreate(tx *gorm.DB) (err error) {
	s.Id = uuid.New()
	s.CreationTime = time.Now()
	s.UpdateTime = time.Now()
	return
}

func (s *Sender) BeforeSave(tx *gorm.DB) (err error)  {
	s.UpdateTime = time.Now()
	return
}
