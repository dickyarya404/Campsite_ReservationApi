package response

import "campsite_reservation/model"

type Campsite struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func FromUseToResponse(campsite *model.Campsite) *Campsite {
	return &Campsite{
		ID:   campsite.ID,
		Name: campsite.Name,
	}
}
