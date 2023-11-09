package handler

import (
	"encoding/json"
	"net/http"

	"github.com/koki120/go-rest-api/log"
	"github.com/koki120/go-rest-api/usecase/input_port"
)

type MemoHandler struct {
	MemoUC input_port.IMemoUseCase
}

func NewMemoHandler() *MemoHandler {
	return &MemoHandler{}
}

func (h *MemoHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger()

	var memoID string
	// TODO:bind query

	res, err := h.MemoUC.FindByID(memoID)
	if err != nil {
		logger.Error("Failed to find memo", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		logger.Error("Failed to encode response to JSON", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
