package response

import "campsite_reservation/model"

type Get struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GetFromEntitiesToResponse(user *model.User) *Get {
	return &Get{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
