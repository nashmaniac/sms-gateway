package business

import (
	"github.com/google/uuid"
	"time"
)

type CreateBusinessParams struct {
	Name string `json:"name"`
}

type BusinessEntityResponse struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	ApiKey string `json:"api_key"`
	CreationTime time.Time `json:"creation_time"`
	UpdateTime time.Time `json:"update_time"`
}