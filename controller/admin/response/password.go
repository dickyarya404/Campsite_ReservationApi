package response

import "campsite_reservation/model"

type Password struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func PasswordFromEntitiesToResponse(admin *model.Admin) *Password {
	return &Password{
		ID:       admin.ID,
		Username: admin.Username,
		Email:    admin.Email,
		Password: admin.Password,
	}
}
