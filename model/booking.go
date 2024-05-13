package model

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID              int       `json:"id" from:"id"`
	CampsiteID      int       `json:"campsite_id" from:"campsite_id"`
	UserID          int       `json:"user_id" from:"user_id"`
	CheckIn         string    `json:"check_in" from:"check_in"`
	CheckOut        string    `json:"check_out" from:"check_out"`
	NumberOfPeaople int       `json:"number_of_peaople" from:"number_of_peaople"`
	TotalPrice      int       `json:"total_price" from:"total_price"`
	Status          string    `gorm:"type:enum('pending','confirmed','cancelled','expired');default:'pending'" json:"status" from:"status"`
	CreatedAt       time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt       time.Time `gorm:"type:DATETIME(0)"`
}

type BookingRepository interface {
	GetAll(camspiteID int) ([]Booking, error)
	GetByID(id int) (Booking, error)
}
