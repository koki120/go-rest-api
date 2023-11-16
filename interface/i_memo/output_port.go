package i_memo

import "github.com/koki120/go-rest-api/domain/entity"

type Repository interface {
	FindByID(userID string, memoID string) (entity.Memo, error)
	Crate(userID string, memo entity.Memo) (entity.Memo, error)
	Update(memo entity.Memo) (entity.Memo, error)
	Delete(memoID string) error
	Search() ([]entity.Memo, int, error)
}
