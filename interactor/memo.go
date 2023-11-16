package interactor

import (
	"github.com/koki120/go-rest-api/domain/entity"
	"github.com/koki120/go-rest-api/interface/i_memo"
)

type MemoUseCase struct {
	memoRepository i_memo.Repository
}

func NewMemoUseCase() i_memo.UseCase {
	return &MemoUseCase{}
}

func (m *MemoUseCase) FindByID(userID string, memoID string) (entity.Memo, error) {
	return m.memoRepository.FindByID(userID, memoID)
}

func (m *MemoUseCase) Create(userID string, memo entity.MemoCreate) (entity.Memo, error) {
	if err := m.memoRepository.Create(userID, entity.MemoCreate{}); err != nil {
		return entity.Memo{}, nil
	}
	return m.FindByID(userID, memo.MemoID)
}

func (m *MemoUseCase) Update(userID string, memo entity.Memo) (entity.Memo, error) {
	if _, err := m.FindByID(userID, memo.MemoID); err != nil {
		return entity.Memo{}, err
	}
	if err := m.memoRepository.Update(memo); err != nil {
		return entity.Memo{}, err
	}
	return m.FindByID(userID, memo.MemoID)
}

func (m *MemoUseCase) Delete(userID string, memoID string) error {
	if _, err := m.FindByID(userID, memoID); err != nil {
		return err
	}
	return m.memoRepository.Delete(memoID)
}

func (m *MemoUseCase) Search() ([]entity.Memo, int, error) {
	return []entity.Memo{}, 0, nil
}
