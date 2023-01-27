package repository

import "database/sql"

type DAO interface { // предост. операции с данными, не раскрывая детали БД. Имеет нужные методы для удобного доступа к данным
	NewUserQuery() UserQuery      // запрос нового юзера
	NewSessionQuery() SessionQuery    // запрос новой сессии
	NewPostQuery() PostQuery     // запрос нового поста
}

type dao struct { // dao перенимает данные с БД
	db *sql.DB       
}

func NewDao(db *sql.DB) DAO { // возвращаем указатель на структуру dao с вложенным db, теперь мы можем получать опред. данные с пом. методов интерфейса DAO
	return &dao{ // возвращаем указатель на структуру
		db: db,
	}
}

func (d dao) NewUserQuery() UserQuery {
	return &userQuery{                   // repository/user
		db: d.db,         // assign data from dao to user Query struct
	}
}

func (d dao) NewSessionQuery() SessionQuery {
	return &sessionQuery{            // repository/session
		db: d.db,           
	}
}

func (d dao) NewPostQuery() PostQuery {
	return &postQuery{              // repository/post
		db: d.db,             // assign data from dao to SessionQuery struct
	}
}
