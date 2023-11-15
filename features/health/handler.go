package health

import (
	"database/sql"
	"net/http"

	"github.com/koki120/go-rest-api/util"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	err := h.DB.Ping()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.WriteStatusText(w, http.StatusOK)
}
