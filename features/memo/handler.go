package memo

import (
	"net/http"

	"github.com/koki120/go-rest-api/domain/entity"
	"github.com/koki120/go-rest-api/interface/i_memo"
	"github.com/koki120/go-rest-api/log"
	"github.com/koki120/go-rest-api/util"
)

type Handler struct {
	MemoUC i_memo.IUseCase
}

func NewHandler(memoUC i_memo.IUseCase) *Handler {
	return &Handler{
		MemoUC: memoUC,
	}
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger()

	memoID := util.GetLastPathParameter(*r)
	res, err := h.MemoUC.FindByID(memoID)
	if err != nil {
		logger.Error("Failed to find memo", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = util.WriteJSONResponse(w, res, http.StatusOK)
	if err != nil {
		logger.Error("Failed to encode response to JSON", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger()

	res, err := h.MemoUC.Crate(entity.MemoCrate{})
	if err != nil {
		logger.Error("Failed to crete memo", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = util.WriteJSONResponse(w, res, http.StatusOK)
	if err != nil {
		logger.Error("Failed to encode response to JSON", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger()

	res, err := h.MemoUC.Update(entity.Memo{})
	if err != nil {
		logger.Error("Failed to update memo", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = util.WriteJSONResponse(w, res, http.StatusOK)
	if err != nil {
		logger.Error("Failed to encode response to JSON", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger()

	memoID := util.GetLastPathParameter(*r)
	err := h.MemoUC.Delete(memoID)
	if err != nil {
		logger.Error("Failed to delete memo", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.WriteStatusText(w, http.StatusNoContent)
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger()
	logger.Warn("Search not implement")
}
