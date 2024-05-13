package request

import "campsite_reservation/model"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Login) ToEntities() *model.Admin {
	return &model.Admin{
		Username: c.Username,
		Password: c.Password,
	}
}
