package memo

import (
	"net/http"

	"github.com/koki120/go-rest-api/api/middleware"
	"github.com/koki120/go-rest-api/domain/entity"
	"github.com/koki120/go-rest-api/interface/i_memo"
	"github.com/koki120/go-rest-api/log"
	"github.com/koki120/go-rest-api/util"
)

type Handler struct {
	MemoUC i_memo.UseCase
}

func NewHandler(memoUC i_memo.UseCase) *Handler {
	return &Handler{
		MemoUC: memoUC,
	}
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger()

	memoID := util.GetLastPathParameter(*r)
	user, err := middleware.GetUserFromContext(r.Context())
	if err != nil {
		logger.Error("Failed to get user in context", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := h.MemoUC.FindByID(user.UserID, memoID)
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

	user, err := middleware.GetUserFromContext(r.Context())
	if err != nil {
		logger.Error("Failed to get user in context", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := h.MemoUC.Crate(user.UserID, entity.MemoCreate{})
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

	user, err := middleware.GetUserFromContext(r.Context())
	if err != nil {
		logger.Error("Failed to get user in context", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := h.MemoUC.Update(user.UserID, entity.Memo{})
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
	user, err := middleware.GetUserFromContext(r.Context())
	if err != nil {
		logger.Error("Failed to get user in context", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.MemoUC.Delete(user.UserID, memoID)
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
