package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}
	err = createTable(db) // создание таблиц
	if err != nil {
		return nil, err
	}
	log.Println("database created")
	return db, nil
}

func createTable(db *sql.DB) error {
	query := []string{} 
	
	users := `
	CREATE TABLE IF NOT EXISTS users(
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)
	`
	posts := `
	CREATE TABLE IF NOT EXISTS posts(
		post_id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		message TEXT NOT NULL,
		email TEXT NOT NULL,
		username TEXT NOT NULL,
		like INTEGER NOT NULL,
		dislike INTEGER NOT NULL,
		category TEXT NOT NULL,
		born TEXT NOT NULL
	);
	`
	session := `
	CREATE TABLE IF NOT EXISTS sessions(
		session_id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		username TEXT NOT NULL,
		token TEXT NOT NULL,
		expiry DATE NOT NULL
	)
	`

	comments := `
	CREATE TABLE IF NOT EXISTS comments(
		comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		username TEXT NOT NULL,
		message TEXT NOT NULL,
		like INTEGER NOT NULL,
		dislike INTEGER NOT NULL,
		born TEXT NOT NULL
	)
	`

	likes := `
	CREATE TABLE IF NOT EXISTS likes(
		username TEXT NOT NULL,
		post_id INTEGER NOT NULL,
		status INTEGER NOT NULL
	)
	`

	dislikes := `
	CREATE TABLE IF NOT EXISTS dislikes(
		username TEXT NOT NULL,
		post_id INTEGER NOT NULL,
		status INTEGER NOT NULL
	)
	`
	commentLikes := `CREATE TABLE IF NOT EXISTS comment_likes(
		username TEXT NOT NULL, 
		comment_id INTEGER NOT NULL,
		status INTEGER NOT NULL
	)
	`
	commentDislikes := `CREATE TABLE IF NOT EXISTS comment_dislikes(
		username TEXT NOT NULL, 
		comment_id INTEGER NOT NULL,
		status INTEGER NOT NULL
	)
	`
	category := `CREATE TABLE IF NOT EXISTS categories(
		category TEXT NOT NULL,
		post_id INTEGER NOT NULL
	)
	`

	query = append(query, users, posts, session, comments, likes, dislikes, commentDislikes, commentLikes, category)
	for _, v := range query {
		_, err := db.Exec(v)
		if err != nil {
			return err
		}
	}
	return nil
}
