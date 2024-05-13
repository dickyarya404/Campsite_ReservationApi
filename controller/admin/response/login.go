package response

import "campsite_reservation/model"

type Login struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func LoginFromEntitiesToResponse(admin *model.Admin) *Login {
	return &Login{
		ID:       admin.ID,
		Username: admin.Username,
		Token:    admin.Token,
	}
}
