package request

import "campsite_reservation/model"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *Login) ToEntities() *model.User {
	return &model.User{
		Username: r.Username,
		Password: r.Password,
	}
}
