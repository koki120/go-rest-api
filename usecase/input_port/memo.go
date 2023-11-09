package input_port

import "github.com/koki120/go-rest-api/domain/entity"

type MemoCrate struct {
	Title string
	Body  string
}

type IMemoUseCase interface {
	FindByID(string) (entity.Memo, error)
	Crate(MemoCrate) (entity.Memo, error)
	Update(entity.Memo) (entity.Memo, error)
	Delete(string) (entity.Memo, error)
}
