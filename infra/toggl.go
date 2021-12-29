package infra

import (
	"github.com/konan0802/dakoku/domain/repository"
)

type togglRepository struct {
}

// NewRepository handlerを生成
func NewTogglRepository() repository.TogglRepository {
	return &togglRepository{}
}
