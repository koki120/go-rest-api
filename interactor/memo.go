package interactor

import (
	"github.com/koki120/go-rest-api/domain/entity"
	"github.com/koki120/go-rest-api/interface/i_memo"
)

type MemoUseCase struct{}

func NewMemoUseCase() i_memo.IUseCase {
	return &MemoUseCase{}
}

func (m *MemoUseCase) FindByID(string) (entity.Memo, error) {
	return entity.Memo{}, nil
}
func (m *MemoUseCase) Crate(entity.MemoCrate) (entity.Memo, error) {
	return entity.Memo{}, nil
}
func (m *MemoUseCase) Update(entity.Memo) (entity.Memo, error) {
	return entity.Memo{}, nil
}
func (m *MemoUseCase) Delete(string) (entity.Memo, error) {
	return entity.Memo{}, nil
}
func (m *MemoUseCase) Search() ([]entity.Memo, int, error) {
	return []entity.Memo{}, 0, nil
}
