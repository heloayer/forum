package model

type Post struct {
	ID         int64
	Title      string
	Content    string
	Category   string
	Comment    []Comment
	Author     User
	Like       int64
	Dislike    int64
	CreateTime string
}
