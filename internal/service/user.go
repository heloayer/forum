package service

import (
	"forum/internal/model"
	"forum/internal/repository"
)

type UserService interface {
	GetUserByEmail(email string) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
}

type userService struct {
	repository.UserQuery
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{
		dao.NewUserQuery(),
	}
}

func (u *userService) GetUserByEmail(email string) (model.User, error) {
	return u.UserQuery.GetUserByEmail(email)
}

func (u *userService) GetUserByUsername(username string) (model.User, error) {
	return u.UserQuery.GetUserByUsername(username)
}
