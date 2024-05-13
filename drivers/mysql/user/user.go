package user

import (
	"campsite_reservation/constant"
	"campsite_reservation/model"
	"campsite_reservation/utils"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Register(user *model.User) error {
	hash, _ := utils.HashPassword(user.Password)
	(*user).Password = hash

	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Login(user *model.User) error {
	var userDB model.User

	if err := r.DB.Where("username = ?", user.Username).First(&userDB).Error; err != nil {
		return errors.New("username or password is incorrect")
	}

	if !utils.CheckPasswordHash(user.Password, userDB.Password) {
		return errors.New("username or password is incorrect")
	}

	(*user).ID = userDB.ID
	(*user).Username = userDB.Username

	return nil
}

func (r *UserRepo) GetAll() ([]model.User, error) {
	var users []model.User

	if err := r.DB.Find(&users).Error; err != nil {
		return []model.User{}, constant.ErrInternalServerError
	}

	return users, nil
}

func (r *UserRepo) GetByID(id int) (model.User, error) {
	var user model.User

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return model.User{}, constant.ErrUserNotFound
	}

	return user, nil
}

func (r *UserRepo) Delete(id int) (model.User, error) {
	var user model.User

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return model.User{}, constant.ErrUserNotFound
	}

	user.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if err := r.DB.Save(&user).Error; err != nil {
		return model.User{}, constant.ErrInternalServerError
	}

	return user, nil
}

func (r *UserRepo) ResetPassword(id int) (model.User, error) {
	var user model.User

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return model.User{}, constant.ErrUserNotFound
	}

	hash, _ := utils.HashPassword("user")
	user.Password = hash

	if err := r.DB.Save(&user).Error; err != nil {
		return model.User{}, constant.ErrInternalServerError
	}

	return user, nil
}

func (r *UserRepo) ChangePassword(id int, newPassword string) (model.User, error) {
	var user model.User

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return model.User{}, constant.ErrUserNotFound
	}

	hash, _ := utils.HashPassword(newPassword)
	user.Password = hash

	if err := r.DB.Save(&user).Error; err != nil {
		return model.User{}, constant.ErrInternalServerError
	}

	return user, nil
}
