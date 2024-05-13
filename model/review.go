package model

import (
	"time"

	"gorm.io/gorm"
)

type Riview struct {
	gorm.Model
	ID          int       `json:"id" from:"id" gorm:"primary_key"`
	CampsiteID  int       `json:"campsite_id" from:"campsite_id" gorm:"foreign_key"`
	UserID      int       `json:"user_id" from:"user_id" gorm:"foreign_key"`
	Rating      int       `json:"rating" from:"rating"`
	Comment     string    `json:"comment" from:"comment"`
	Review_date string    `json:"review_date" from:"review_date"`
	CreatedAt   time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt   time.Time `gorm:"type:DATETIME(0)"`
}
