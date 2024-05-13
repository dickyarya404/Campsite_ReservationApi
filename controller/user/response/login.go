package response

import "campsite_reservation/model"

type Login struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func LoginFromEntitiesToResponse(user *model.User) *Login {
	return &Login{
		ID:       user.ID,
		Username: user.Username,
		Token:    user.Token,
	}
}
