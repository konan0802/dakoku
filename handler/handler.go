package handler

import (
	"github.com/konan0802/dakoku/usecase"

	"github.com/spf13/cobra"
)

type Handler interface {
	Tasks(cmd *cobra.Command, args []string) (string, error)
}

type handler struct {
	usecase usecase.Usecase
}

// NewHandler handlerを生成
func NewHandler(uc usecase.Usecase) Handler {
	return &handler{uc}
}

func (hdr handler) Tasks(cmd *cobra.Command, args []string) (string, error) {
	ret := "called tasks YEAH"
	return ret, nil
}
