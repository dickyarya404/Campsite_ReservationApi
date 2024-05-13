package response

import "campsite_reservation/model"

type Delete struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	CampsiteID string `json:"capsite_id"`
	BookingID  string `json:"booking_id"`
	FacilityID string `json:"facility_id"`
	Email      string `json:"email"`
}

func DeleteFromEntitiesToResponse(admin *model.Admin) *Delete {
	return &Delete{
		ID:         admin.ID,
		Username:   admin.Username,
		CampsiteID: admin.CampsiteID,
		BookingID:  admin.BookingID,
		FacilityID: admin.FacilityID,
		Email:      admin.Email,
	}
}
