package models

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type MessageTemplate struct {
	Id uuid.UUID `json:"id" gorm:"primary_key"`
	Template string `json:"template"`
	Count int64 `json:"count"`
	Category string `json:"category"`
	CreationTime  time.Time `json:"creation_time"`
	UpdateTime  time.Time `json:"update_time"`
}

func (template *MessageTemplate) BeforeCreate(tx *gorm.DB) (err error)  {
	template.Id = uuid.New()
	template.CreationTime = time.Now()
	template.UpdateTime = time.Now()
	return
}
func (template *MessageTemplate) BeforeSave(tx *gorm.DB) (err error)  {
	template.UpdateTime = time.Now()
	return
}

func (template MessageTemplate) OutputFormattedMessage(code string) string {
	return fmt.Sprintf(template.Template, code)
}


