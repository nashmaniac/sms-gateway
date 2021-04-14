package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Message struct {
	Id string `json:"id" gorm:"primary_key"`
	BusinessEntityId uuid.UUID `json:"business_entity_id" gorm:"index:,not null"`
	FromNumber string `json:"from_number" gorm:"not null"`
	ToNumber string `json:"to_number" gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
	RequestedDate time.Time `json:"requested_date" gorm:"index:,not null"`
	RequestedTime time.Time `json:"requested_time" gorm:"not null"`
	SendingTime time.Time `json:"sending_time"`
	ResponseId string `json:"response_id"`
	Reason string `json:"reason"`
	CreationTime time.Time `json:"creation_time"`
	UpdateTime time.Time `json:"update_time"`
	IsSent bool `json:"is_sent" gorm:"default:false"`
	IsSuccessful bool `json:"is_successful" gorm:"default:false"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) (err error)  {
	m.Id = strconv.FormatUint(uint64(time.Now().UnixNano()), 10)
	m.RequestedTime = time.Now()
	m.RequestedDate = time.Now()
	m.CreationTime = time.Now()
	m.UpdateTime = time.Now()
	return
}

func (m *Message) BeforeSave(tx *gorm.DB) (err error)  {
	m.UpdateTime = time.Now()
	return
}

