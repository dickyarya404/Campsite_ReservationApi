package response

import (
	campsite_response "campsite_reservation/controller/campsite/response"
	"campsite_reservation/model"
)

type Get struct {
	ID       int                        `json:"id"`
	Username string                     `json:"username"`
	Campsite campsite_response.Campsite `json:"campsite"`
	Booking  booking_response.Booking   `json:"booking"`
	Email    string                     `json:"email"`
}

func GetFromEntitiesToResponse(admin *model.Admin) *Get {
	return &Get{
		ID:       admin.ID,
		Username: admin.Username,
		Email:    admin.Email,
	}
}
