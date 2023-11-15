package i_user

import "github.com/koki120/go-rest-api/domain/entity"

type IUseCase interface {
	FindByID(string) (entity.User, error)
	Crate(entity.UserCrate) (entity.User, error)
	Update(entity.User) (entity.User, error)
	Delete(string) (entity.User, error)
	Search() ([]entity.User, int, error)
}
