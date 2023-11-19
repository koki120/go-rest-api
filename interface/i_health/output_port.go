package i_health

type Repository interface {
	Health() error
}
