package database

import "github.com/hukurou-s/user-auth-api-with-jwt/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	user := domain.User{
		Snum:     u.Snum,
		Name:     u.Name,
		Password: u.Password,
	}
	if result := repo.Create(&user); result.Error != nil {
		err = result.Error
		return
	}
	id = int(user.ID)
	return
}

func (repo *UserRepository) FindByID(id int) (user domain.User, err error) {
	if result := repo.Where(&user, id); result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (repo *UserRepository) FindBySnum(snum string) (user domain.User, err error) {
	if result := repo.Where("snum = ?", snum).First(&user); result.Error != nil {
		err = result.Error
		return
	}
	return
}
