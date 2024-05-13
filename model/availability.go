package model

import (
	"time"

	"gorm.io/gorm"
)

type Availability struct {
	*gorm.Model
	ID         int       ` gorm:"primary_key" json:"availability_id" from:"availability_id"`
	CampsiteID int       `gorm:"foreign_key" json:"campsite_id" from:"campsite_id" `
	Quantity   int       `json:"quantity" from:"quantity"`
	CreatedAt  time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt  time.Time `gorm:"type:DATETIME(0)"`
}

type AvailabilityRepositoryInterface interface {
	FindAll() []Availability
}

type AvailabilityUseCaseInterface interface {
	FindAll() []Availability
}
