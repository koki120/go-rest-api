package health

import (
	"net/http"

	"github.com/koki120/go-rest-api/interface/i_health"
	"github.com/koki120/go-rest-api/util"
)

type Handler struct {
	healthUC i_health.UseCase
}

func NewHandler(healthUC i_health.UseCase) *Handler {
	return &Handler{
		healthUC: healthUC,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if err := h.healthUC.Health(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteStatusText(w, http.StatusOK)
}
