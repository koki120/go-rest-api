package user

import (
	"net/http"

	"github.com/koki120/go-rest-api/domain/entity"
	"github.com/koki120/go-rest-api/interface/i_user"
	"github.com/koki120/go-rest-api/log"
	"github.com/koki120/go-rest-api/util"
)

type Handler struct {
	UserUC i_user.IUseCase
}

func NewHandler(userUC i_user.IUseCase) *Handler {
	return &Handler{
		UserUC: userUC,
	}
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger()

	userID := util.GetLastPathParameter(*r)
	res, err := h.UserUC.FindByID(userID)
	if err != nil {
		logger.Error("Failed to find user", err)
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

	res, err := h.UserUC.Crate(entity.UserCrate{})
	if err != nil {
		logger.Error("Failed to crete user", err)
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

	res, err := h.UserUC.Update(entity.User{})
	if err != nil {
		logger.Error("Failed to update user", err)
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

	userID := util.GetLastPathParameter(*r)
	err := h.UserUC.Delete(userID)
	if err != nil {
		logger.Error("Failed to delete user", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.WriteStatusText(w, http.StatusNoContent)
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger()
	logger.Warn("Search not implement")
}
