package response

import "campsite_reservation/model"

type CreateAccount struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func CreateAccountFromEntitiesToResponse(admin *model.Admin) *CreateAccount {
	return &CreateAccount{
		ID:       admin.ID,
		Email:    admin.Email,
		Username: admin.Username,
	}
}
