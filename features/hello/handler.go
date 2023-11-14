package hello

import "net/http"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
