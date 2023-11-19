package health

import (
	"database/sql"

	"github.com/koki120/go-rest-api/interface/i_health"
)

type HealthRepository struct {
	db *sql.DB
}

func NewHealthRepository(db *sql.DB) i_health.Repository {
	return &HealthRepository{
		db: db,
	}
}

func (h *HealthRepository) Health() error {
	return h.db.Ping()
}
