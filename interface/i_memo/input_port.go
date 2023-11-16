package i_memo

import "github.com/koki120/go-rest-api/domain/entity"

type UseCase interface {
	FindByID(userID string, memoID string) (entity.Memo, error)
	Crate(userID string, memo entity.MemoCreate) (entity.Memo, error)
	Update(userID string, memo entity.Memo) (entity.Memo, error)
	Delete(userID string, memoID string) error
	Search() ([]entity.Memo, int, error)
}
