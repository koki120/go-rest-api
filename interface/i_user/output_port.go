package i_user

import "github.com/koki120/go-rest-api/domain/entity"

type Repository interface {
	FindByID(userID string) (entity.User, error)
	Create(user entity.UserCrate) error
	Update(user entity.User) error
	Delete(userID string) error
	Search() ([]entity.User, int, error)
}
