package usecase

import "github.com/hukurou-s/user-auth-api-with-jwt/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	Update(domain.User) (int, error)
	FindByID(int) (domain.User, error)
	FindBySnum(string) (domain.User, error)
}
