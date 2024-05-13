package admin

import "campsite_reservation/model"

type AdminUseCase struct {
	repository model.AdminRepositoryInterface
}

func NewAdminController(repository model.AdminRepositoryInterface) *AdminUseCase {
	return &AdminUseCase{
		repository: repository,
	}
}
