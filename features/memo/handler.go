package memo

import (
	"encoding/json"
	"net/http"

	i_memo "github.com/koki120/go-rest-api/interface/imemo"
	"github.com/koki120/go-rest-api/log"
)

type Handler struct {
	MemoUC i_memo.IUseCase
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
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
