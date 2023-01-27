package service

import (
	"strings"

	"forum/internal/model"
	"forum/internal/repository"
)

type PostService interface {
	CreatePost(post model.Post) (int64, error)
	CreateCategory(category model.Category) error
	GetAllPosts() ([]model.Post, error)
	CommentPost(comment model.Comment) error
	GetAllCommentByPostID(postId int64) ([]model.Comment, error)
	GetCommentByCommentID(commentId int64) (model.Comment, error)
	GetPostById(postId int64) (model.Post, error)
	LikePost(postId int64, username string) error
	DislikePost(postId int64, username string) error
	LikeComment(commentId int64, username string) error
	DislikeComment(commentId int64, username string) error
	GetLikedPostIdByUser(user model.User) ([]int64, error)
	GetCategory() ([]model.Category, error)
	GetFilterPosts(category string) ([]model.Post, error)
}

type postService struct {
	repository.PostQuery
}

func NewPostService(dao repository.DAO) PostService {
	return &postService{
		dao.NewPostQuery(),
	}
}

func (p *postService) CreatePost(post model.Post) (int64, error) {
	return p.PostQuery.CreatePost(post)
}

func (p *postService) GetAllPosts() ([]model.Post, error) {
	posts, err := p.PostQuery.GetAllPosts()
	if err != nil {
		return nil, err
	}
	result := []model.Post{}
	for i := len(posts) - 1; i >= 0; i-- {
		result = append(result, posts[i])
	}
	return result, nil
}

func (p *postService) CommentPost(comment model.Comment) error {
	return p.PostQuery.CommentPost(comment)
}

func (p *postService) GetAllCommentByPostID(postId int64) ([]model.Comment, error) {
	comments, err := p.PostQuery.GetAllCommentByPostID(postId)
	if err != nil {
		return nil, err
	}
	result := []model.Comment{}
	for i := len(comments) - 1; i >= 0; i-- {
		result = append(result, comments[i])
	}
	return result, nil
}

func (p *postService) GetPostById(postId int64) (model.Post, error) {
	return p.PostQuery.GetPostById(postId)
}

func (p *postService) LikePost(postId int64, username string) error {
	return p.PostQuery.LikePost(postId, username)
}

func (p *postService) DislikePost(postId int64, username string) error {
	return p.PostQuery.DislikePost(postId, username)
}

func (p *postService) LikeComment(commentId int64, username string) error {
	return p.PostQuery.LikeComment(commentId, username)
}

func (p *postService) DislikeComment(commentId int64, username string) error {
	return p.PostQuery.DislikeComment(commentId, username)
}

func (p *postService) GetCommentByCommentID(commentId int64) (model.Comment, error) {
	return p.PostQuery.GetCommentByCommentID(commentId)
}

func (p *postService) CreateCategory(category model.Category) error {
	return p.PostQuery.CreateCategory(category)
}

func (p *postService) GetLikedPostIdByUser(user model.User) ([]int64, error) {
	return p.PostQuery.GetLikedPostIdByUser(user)
}

func (p *postService) GetCategory() ([]model.Category, error) {
	return p.PostQuery.GetCategory()
}

func (p *postService) GetFilterPosts(genre string) ([]model.Post, error) {
	switch genre {
	case "most-like":
		sort, err := p.PostQuery.GetAllPosts()
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(sort); i++ {
			for j := 0; j < len(sort)-i-1; j++ {
				if sort[j+1].Like > sort[j].Like {
					sort[j], sort[j+1] = sort[j+1], sort[j]
				}
			}
		}
		return sort, nil

	case "most-dislike":
		sort, err := p.PostQuery.GetAllPosts()
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(sort); i++ {
			for j := 0; j < len(sort)-i-1; j++ {
				if sort[j+1].Dislike > sort[j].Dislike {
					sort[j], sort[j+1] = sort[j+1], sort[j]
				}
			}
		}
		return sort, nil

	default:
		categories, err := p.PostQuery.GetCategory()
		if err != nil {
			return nil, err
		}
		var postId []int64
		for _, v := range categories {
			category := strings.Fields(v.CategoryName)
			genreName := strings.Fields(genre)
			for _, k := range category {
				for _, j := range genreName {
					if k == j {
						postId = append(postId, v.PostId)
					}
				}
			}
		}
		var allPost []model.Post
		for _, v := range postId {
			post, err := p.PostQuery.GetPostById(v)
			if err != nil {
				return nil, err
			}
			allPost = append(allPost, post)
		}
		return allPost, nil
	}
}
