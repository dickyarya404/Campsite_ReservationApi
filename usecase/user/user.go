package user

import (
	"campsite_reservation/constant"
	"campsite_reservation/middleware"
	"campsite_reservation/model"
	"strings"
)

type UserUseCase struct {
	repository model.UserRepositoryInterface
}

func NewUserUseCase(repository model.UserRepositoryInterface) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (u *UserUseCase) Register(user *model.User) (model.User, error) {
	if user.Email == "" || user.Password == "" || user.Username == "" {
		return model.User{}, constant.ErrAllFieldsMustBeFilled
	}

	err := u.repository.Register(user)

	if err != nil {
		if strings.HasPrefix(err.Error(), "Error 1062") {
			if strings.HasSuffix(err.Error(), "email'") {
				return model.User{}, constant.ErrEmailAlreadyExist
			} else if strings.HasSuffix(err.Error(), "username'") {
				return model.User{}, constant.ErrUsernameAlreadyExist
			}
		} else {
			return model.User{}, constant.ErrInternalServerError
		}
	}

	return *user, nil
}

func (u *UserUseCase) Login(user *model.User) (model.User, error) {
	if user.Username == "" || user.Password == "" {
		return model.User{}, constant.ErrAllFieldsMustBeFilled
	}

	err := u.repository.Login(user)

	(*user).Token = middleware.GenerateTokenJWT(user.ID, user.Username, "user")

	if err != nil {
		return model.User{}, constant.ErrInvalidUsernameOrPassword
	}

	return *user, nil
}

func (u *UserUseCase) GetAll() ([]model.User, error) {
	users, err := u.repository.GetAll()

	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (u *UserUseCase) GetByID(id int) (model.User, error) {
	user, err := u.repository.GetByID(id)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *UserUseCase) Delete(id int) (model.User, error) {
	user, err := u.repository.Delete(id)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *UserUseCase) ChangePassword(id int, newPassword string) (model.User, error) {
	if newPassword == "" {
		return model.User{}, constant.ErrAllFieldsMustBeFilled
	}

	user, err := u.repository.ChangePassword(id, newPassword)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *UserUseCase) ResetPassword(id int) (model.User, error) {
	user, err := u.repository.ResetPassword(id)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
