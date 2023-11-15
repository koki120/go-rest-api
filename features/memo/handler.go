package memo

import (
	"net/http"

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
	pathParm := util.GetLastPathParameter(*r)
	logger.Warn("FindByID not implement" + pathParm)

	// var memoID string
	// // TODO:bind query

	// res, err := h.MemoUC.FindByID(memoID)
	// if err != nil {
	// 	logger.Error("Failed to find memo", err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// if err = json.NewEncoder(w).Encode(res); err != nil {
	// 	logger.Error("Failed to encode response to JSON", err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger()
	pathParm := util.GetLastPathParameter(*r)
	logger.Warn("Create not implement " + pathParm)

}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger()
	pathParm := util.GetLastPathParameter(*r)
	logger.Warn("Update not implement" + pathParm)

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger()
	pathParm := util.GetLastPathParameter(*r)
	logger.Warn("Delete not implement" + pathParm)
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	logger := log.NewLogger()
	logger.Warn("Search not implement")
}
