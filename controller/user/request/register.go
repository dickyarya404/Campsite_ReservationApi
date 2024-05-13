package request

import "campsite_reservation/model"

type Register struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *Register) ToEntities() *model.User {
	return &model.User{
		Email:    r.Email,
		Username: r.Username,
		Password: r.Password,
	}
}
