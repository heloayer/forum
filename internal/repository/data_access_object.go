package repository

import "database/sql"

type DAO interface { 
	NewUserQuery() UserQuery     
	NewSessionQuery() SessionQuery   
	NewPostQuery() PostQuery    
}

type dao struct {
	db *sql.DB       
}

func NewDao(db *sql.DB) DAO { 
	return &dao{ 
		db: db,
	}
}

func (d dao) NewUserQuery() UserQuery {
	return &userQuery{                 
		db: d.db,        
	}
}

func (d dao) NewSessionQuery() SessionQuery {
	return &sessionQuery{           
		db: d.db,           
	}
}

func (d dao) NewPostQuery() PostQuery {
	return &postQuery{   
		db: d.db,           
	}
}
