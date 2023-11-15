package health

import (
	"database/sql"
	"net/http"
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
