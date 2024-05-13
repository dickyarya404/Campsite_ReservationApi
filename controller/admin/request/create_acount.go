package request

import "campsite_reservation/model"

type CreateAccount struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	CampsiteID string `json:"campsite_id"`
	BookingID  string `json:"booking_id"`
	FacilityID string `json:"facility_id"`
}

func (c *CreateAccount) ToEntities() *model.Admin {
	return &model.Admin{
		Username:   c.Username,
		Email:      c.Email,
		Password:   c.Password,
		CampsiteID: c.CampsiteID,
		BookingID:  c.BookingID,
		FacilityID: c.FacilityID,
	}
}
