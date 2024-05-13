package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int    `gorm:"primary_key"`
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string
	Token     string         `gorm:"-"`
	CreatedAt time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt gorm.DeletedAt `gorm:"type:DATETIME(0)"`
}

type UserRepositoryInterface interface {
	Register(user *User) error
	Login(user *User) error
	GetAll() ([]User, error)
	GetByID(id int) (User, error)
	Delete(id int) (User, error)
	ChangePassword(id int, newPassword string) (User, error)
	ResetPassword(id int) (User, error)
}

type UserUseCaseInterface interface {
	Register(user *User) (User, error)
	Login(user *User) (User, error)
	GetAll() ([]User, error)
	GetByID(id int) (User, error)
	Delete(id int) (User, error)
	ChangePassword(id int, newPassword string) (User, error)
	ResetPassword(id int) (User, error)
}
