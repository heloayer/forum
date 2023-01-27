package service

import (
	"errors"

	"forum/internal/model"
	"forum/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignUp(user model.User) error
	SignIn(user model.User, session model.Session) (model.Session, error)
	Logout(token string) error
}

type authService struct {
	repository.UserQuery          // интерфейс // rep/user
	repository.SessionQuery          // нтерфейс // rep/query
}

func NewAuthService(dao repository.DAO) AuthService {
	return &authService{
		dao.NewUserQuery(),        // rep/data access object
		dao.NewSessionQuery(),     // rep/data access object
	}
}

func (a *authService) SignUp(user model.User) error {
	_, err := a.UserQuery.GetUserByEmail(user.Email)
	_, err1 := a.UserQuery.GetUserByUsername(user.Username)
	if err == nil || err1 == nil {
		return errors.New("user already exists")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0)
	if err != nil {
		return errors.New("password hash was fail")
	}
	newUser := model.User{
		Email:    user.Email,
		Username: user.Username,
		Password: string(passwordHash),
	}
	err = a.UserQuery.CreateUser(newUser)
	if err != nil {
		return err
	}
	return nil
}

func (a *authService) SignIn(user model.User, session model.Session) (model.Session, error) {
	u, err := a.UserQuery.GetUserByEmail(user.Email)
	if err != nil {
		return model.Session{}, errors.New("user is not registered")
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return model.Session{}, errors.New("wrong password")
	}
	session.Username = u.Username
	err = a.SessionQuery.CreateSession(session)
	if err != nil {
		return model.Session{}, err
	}
	return session, nil
}

func (a *authService) Logout(token string) error {
	err := a.SessionQuery.DeleteSession(token)
	if err != nil {
		return err
	}
	return nil
}
