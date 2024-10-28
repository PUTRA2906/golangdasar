package usecase

import (
	"myapp/entity"
	"myapp/repository"
)

type UserUsecase interface {
	GetUsers() ([]entity.User, error)
	GetUser(id int) (*entity.User, error)
	CreateUser(user *entity.User) error
	DeleteUser(id int) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) GetUsers() ([]entity.User, error) {
	return u.userRepo.GetAll()
}

func (u *userUsecase) GetUser(id int) (*entity.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userUsecase) CreateUser(user *entity.User) error {
	return u.userRepo.Create(user)
}

func (u *userUsecase) DeleteUser(id int) error {
	return u.userRepo.Delete(id)
}
