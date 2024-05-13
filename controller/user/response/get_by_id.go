package response

import "campsite_reservation/model"

type GetByID struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func GetByIDFromEntitiesToResponse(user *model.User) *GetByID {
	return &GetByID{
		ID:       user.ID,
		Username: user.Username,
	}
}
