package interactor

import (
	"github.com/koki120/go-rest-api/domain/entity"
	"github.com/koki120/go-rest-api/interface/i_user"
)

type UserUseCase struct {
	userRepository i_user.Repository
}

func (u *UserUseCase) FindByID(userID string) (entity.User, error) {
	return u.userRepository.FindByID(userID)
}

func (u *UserUseCase) Create(user entity.UserCrate) (entity.User, error) {
	if err := u.userRepository.Create(user); err != nil {
		return entity.User{}, nil
	}
	return u.userRepository.FindByID(user.UserID)
}

func (u *UserUseCase) Update(user entity.User) (entity.User, error) {
	if err := u.userRepository.Update(user); err != nil {
		return entity.User{}, nil
	}
	return u.userRepository.FindByID(user.UserID)
}

func (u *UserUseCase) Delete(userID string) error {
	return u.userRepository.Delete(userID)
}

func (u *UserUseCase) Search() ([]entity.User, int, error) {
	panic("")
}
