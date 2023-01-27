package repository

import (
	"database/sql"
	"log"

	"forum/internal/model"
)

type PostQuery interface {
	CreatePost(post model.Post) (int64, error)
	GetAllPosts() ([]model.Post, error)
	GetPostById(postId int64) (model.Post, error)
	CreateCategory(category model.Category) error
	GetCategory() ([]model.Category, error)
	LikePost(postId int64, username string) error
	DislikePost(postId int64, username string) error
	LikeComment(commentId int64, username string) error
	DislikeComment(commentId int64, username string) error
	GetLikedPostIdByUser(user model.User) ([]int64, error)
	CommentPost(comment model.Comment) error
	GetAllCommentByPostID(postId int64) ([]model.Comment, error)
	GetCommentByCommentID(commentId int64) (model.Comment, error)
}

type postQuery struct {       // used in data access object
	db *sql.DB             // assign data from dao to SessionQuery struct
}

func (p *postQuery) CreatePost(post model.Post) (int64, error) {
	stmt, err := p.db.Prepare("INSERT INTO posts (title, message, email, username, category, like, dislike, born) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(post.Title, post.Content, post.Author.Email, post.Author.Username, post.Category, 0, 0, post.CreateTime)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	post.ID = id
	return id, nil
}

func (p *postQuery) GetAllPosts() ([]model.Post, error) {
	stmt, err := p.db.Prepare("SELECT * FROM posts")
	if err != nil {
		return []model.Post{}, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return []model.Post{}, err
	}
	defer rows.Close()
	var all []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author.Email, &post.Author.Username, &post.Like, &post.Dislike, &post.Category, &post.CreateTime); err != nil {
			return []model.Post{}, err
		}
		all = append(all, post)
	}
	return all, nil
}

func (p *postQuery) GetPostById(postId int64) (model.Post, error) {
	stmt, err := p.db.Prepare("SELECT post_id, title,  message, email, username, like, dislike, category, born FROM posts WHERE post_id = ? ")
	if err != nil {
		log.Println(err)
	}
	row := stmt.QueryRow(postId)

	var post model.Post
	if err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Author.Email, &post.Author.Username, &post.Like, &post.Dislike, &post.Category, &post.CreateTime); err != nil {
		return model.Post{}, err
	}
	return post, nil
}

func (p *postQuery) CreateCategory(category model.Category) error {
	query := `INSERT INTO categories(category, post_id) VALUES(?,?)`
	_, err := p.db.Exec(query, category.CategoryName, category.PostId)
	if err != nil {
		return err
	}
	return nil
}

func (p *postQuery) GetCategory() ([]model.Category, error) {
	query := `SELECT * FROM categories`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	var result []model.Category
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.CategoryName, &category.PostId); err != nil {
			return nil, err
		}
		result = append(result, category)
	}
	return result, nil
}
