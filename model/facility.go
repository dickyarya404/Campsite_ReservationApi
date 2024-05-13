package model

import (
	"time"

	"gorm.io/gorm"
)

type Facility struct {
	*gorm.Model
	ID          int       `json:"id" from:"id" gorm:"primary_key"`
	CampsiteID  int       `json:"campsite_id" from:"campsite_id" gorm:"primary_key"`
	Name        string    `json:"name" from:"name"`
	Description string    `json:"description" from:"description"`
	CreatedAt   time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt   time.Time `gorm:"type:DATETIME(0)"`
}

type FacilityRepositoryInterface interface {
}
