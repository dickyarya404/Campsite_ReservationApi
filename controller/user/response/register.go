package response

import "campsite_reservation/model"

type Register struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func RegisterFromEntitiesToResponse(user *model.User) *Register {
	return &Register{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}
