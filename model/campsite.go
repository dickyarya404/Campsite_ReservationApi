package model

import (
	"time"

	"gorm.io/gorm"
)

type Campsite struct {
	*gorm.Model
	ID        int       `gorm:"primary_key"`
	Name      string    `json:"name" from:"name"`
	Location  string    `json:"location" from:"location"`
	Latitude  float64   `json:"latitude" from:"latitude"`
	Lontitude float64   `json:"lontitude" from:"lontitude"`
	AreaCamp  int       `json:"luas" from:"luas"`
	Price     int       `json:"price" from:"price"`
	CreatedAt time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time `gorm:"type:DATETIME(0)"`
}

type CampsiteRepositoryInferface interface {
}
