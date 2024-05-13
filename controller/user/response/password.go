package response

import "campsite_reservation/model"

type Password struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func PasswordFromEntitiesToResponse(user *model.User) *Password {
	return &Password{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}
