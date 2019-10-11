package usecase

import "github.com/hukurou-s/user-auth-api-with-jwt/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) UserByID(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByID(id)
	return
}

func (interactor *UserInteractor) UserBySnum(snum string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindBySnum(snum)
	return
}
