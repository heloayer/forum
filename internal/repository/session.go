package repository

import (
	"database/sql"
	"errors"
	"log"

	"forum/internal/model"
)

type SessionQuery interface {
	CreateSession(session model.Session) error
	GetSessionByEmail(email string) (model.Session, error)
	GetSessionByToken(token string) (model.Session, error)
	DeleteSession(token string) error
}

type sessionQuery struct {       // used in data access object
	db *sql.DB             // assign data from dao to sessionQuery struct
}

func (s *sessionQuery) CreateSession(session model.Session) error {
	stmt, err := s.db.Prepare("INSERT INTO sessions(email, username, token, expiry) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(session.Email, session.Username, session.Token, session.Expiry)
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionQuery) GetSessionByEmail(email string) (model.Session, error) {
	stmt, err := s.db.Prepare("SELECT session_id, email, username, token, expiry FROM sessions WHERE email = ?")
	if err != nil {
		log.Println(err)
	}
	row := stmt.QueryRow(email)
	var session model.Session
	if err := row.Scan(&session.ID, &session.Email, &session.Username, &session.Token, &session.Expiry); err != nil {
		return session, err
	}
	return session, nil
}

func (s *sessionQuery) GetSessionByToken(token string) (model.Session, error) {
	stmt, err := s.db.Prepare("SELECT session_id, email, username, token, expiry FROM sessions WHERE token = ?") // ОТСЮДА
	if err != nil {
		log.Println(err)
	}
	row := stmt.QueryRow(token)
	var session model.Session
	if err := row.Scan(&session.ID, &session.Email, &session.Username, &session.Token, &session.Expiry); err != nil {
		return model.Session{}, err
	}
	return session, nil
}

func (s *sessionQuery) DeleteSession(token string) error {
	stmt, err := s.db.Prepare("DELETE FROM sessions WHERE token = ?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(token)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("delete session was failed")
	}
	return nil
}
