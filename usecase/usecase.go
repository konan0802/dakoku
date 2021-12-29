package usecase

import (
	"github.com/konan0802/dakoku/domain/repository"
)

type Usecase interface {
}

type usecase struct {
	arp repository.AsanaRepository
	trp repository.TogglRepository
}

func NewUsecase(arp repository.AsanaRepository, trp repository.TogglRepository) Usecase {
	return &usecase{
		arp: arp,
		trp: trp,
	}
}
