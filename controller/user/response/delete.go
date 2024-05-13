package response

import "campsite_reservation/model"

type Delete struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func DeleteFromEntitiesToResponse(user *model.User) *Delete {
	return &Delete{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
