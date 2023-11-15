package i_memo

import "github.com/koki120/go-rest-api/domain/entity"

type IUseCase interface {
	FindByID(string) (entity.Memo, error)
	Crate(entity.MemoCrate) (entity.Memo, error)
	Update(entity.Memo) (entity.Memo, error)
	Delete(string) error
	Search() ([]entity.Memo, int, error)
}
