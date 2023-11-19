package interactor

import (
	"github.com/koki120/go-rest-api/interface/i_health"
)

type HealthUseCase struct {
	healthRepository i_health.Repository
}

func NewHealthUseCase(healthRepository i_health.Repository) i_health.UseCase {
	return &HealthUseCase{
		healthRepository: healthRepository,
	}
}

func (h *HealthUseCase) Health() error {
	return h.healthRepository.Health()
}
