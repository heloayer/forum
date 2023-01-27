package repository

import (
	"database/sql"

	"forum/internal/model"
)

type UserQuery interface {
	CreateUser(user model.User) error
	GetUserByEmail(email string) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
}

type userQuery struct { // used in data access object
	db *sql.DB // assigns data from dao struct
}

func (u *userQuery) CreateUser(user model.User) error {
	stmt, err := u.db.Prepare("INSERT INTO users(email, username, password) values(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Email, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *userQuery) GetUserByEmail(email string) (model.User, error) {
	stmt, err := u.db.Prepare("SELECT user_id,email,password,username FROM users WHERE email = ?")
	if err != nil {
		return model.User{}, err
	}
	row := stmt.QueryRow(email)
	var user model.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Username); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userQuery) GetUserByUsername(username string) (model.User, error) {
	stmt, err := u.db.Prepare("SELECT user_id,email,password,username FROM users WHERE username = ?")
	if err != nil {
		return model.User{}, err
	}
	row := stmt.QueryRow(username)
	var user model.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Username); err != nil {
		return model.User{}, err
	}
	return user, nil
}
