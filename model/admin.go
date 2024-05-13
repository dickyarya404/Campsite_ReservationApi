package model

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	*gorm.Model
	ID         int       `gorm:"primary_key"`
	Username   string    `gorm:"unique"`
	CampsiteID string    `gorm:"type:varchar(255)" json:"campsite_id" from:"campsite_id"`
	BookingID  string    `gorm:"type:varchar(255)" json:"booking_id" from:"booking_id"`
	FacilityID string    `gorm:"type:varchar(255)" json:"facility_id" from:"facility_id"`
	Email      string    `gorm:"unique"`
	Password   string    `json:"password" from:"password"`
	Token      string    `gorm:"-"`
	CreatedAt  time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt  time.Time `gorm:"type:DATETIME(0)"`
}

type AdminRepositoryInterface interface {
	CreateAccount(admin *Admin) error
	Login(admin *Admin) error
	GetById(id int) (Admin, error)
	GetAll() ([]Admin, error)
	DeleteAccount(id int) (Admin, error)
	ResetAccount(id int) (Admin, error)
	ChangePassword(id int, newPassword string) (Admin, error)
}

type AdminUseCaseInterface interface {
	CreateAccount(admin *Admin) (Admin, error)
	Login(admin *Admin) (Admin, error)
	GetByID(id int) (Admin, error)
	GetAll() ([]Admin, error)
	DeleteAccount(id int) (Admin, error)
	ResetPassword(id int) (Admin, error)
	ChangePassword(id int, newPassword string) (Admin, error)
}
