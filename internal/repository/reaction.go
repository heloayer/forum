package repository

import (
	"log"

	"forum/internal/model"
)

func (p *postQuery) LikePost(postId int64, username string) error {
	post, err := p.GetPostById(postId)
	if err != nil {
		log.Println(err)
		return err
	}
	// проверяем есть ли лайк
	query := `SELECT status FROM likes WHERE post_id = ? AND username = ?`
	var like int
	_ = p.db.QueryRow(query, postId, username).Scan(&like)

	// проверка на дизлайк
	query = `SELECT status FROM dislikes WHERE post_id = ? AND username = ?`
	var dislike int
	_ = p.db.QueryRow(query, postId, username).Scan(&dislike)

	// если лайка нет то ставим лайк
	if like == 0 && dislike == 0 {
		query = `INSERT INTO likes(post_id, username, status) VALUES(?,?,?)`
		_, err := p.db.Exec(query, postId, username, 1)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `UPDATE posts SET like = ? WHERE post_id = ?`
		post.Like++
		_, err = p.db.Exec(query, post.Like, post.ID)
		if err != nil {
			log.Println(err)
			return err
		}
		// если лайка есть то удаляем лайк
	} else if like == 0 && dislike == 1 {
		query = `DELETE FROM dislikes WHERE post_id = ? AND username = ?`
		_, err := p.db.Exec(query, postId, username)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `INSERT INTO likes(post_id, username, status) VALUES(?,?,?)`
		_, err = p.db.Exec(query, postId, username, 1)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `UPDATE posts SET like = ?, dislike = ? WHERE post_id = ? `
		post.Like++
		post.Dislike--
		_, err = p.db.Exec(query, post.Like, post.Dislike, post.ID)
		if err != nil {
			log.Println(err)
			return err
		}
		// если срабатывает else то значит пользователь дважды нажал на кнопку и соответственно мы стираем лайк
	} else {
		query = `DELETE FROM likes WHERE post_id = ? AND username = ?`
		_, err := p.db.Exec(query, postId, username)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `UPDATE posts SET like = ? WHERE post_id = ?`
		post.Like--
		_, err = p.db.Exec(query, post.Like, post.ID)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (p *postQuery) DislikePost(postId int64, username string) error {
	post, err := p.GetPostById(postId)
	if err != nil {
		log.Println(err)
		return err
	}
	query := `SELECT status FROM dislikes WHERE post_id = ? AND username = ?`
	var dislike int
	_ = p.db.QueryRow(query, postId, username).Scan(&dislike)

	query = `SELECT status FROM likes WHERE post_id = ? AND username = ?`
	var like int
	_ = p.db.QueryRow(query, postId, username).Scan(&like)

	if dislike == 0 && like == 0 {
		query = `INSERT INTO dislikes(post_id, username, status) VALUES(?,?,?)`
		_, err := p.db.Exec(query, postId, username, 1)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `UPDATE posts SET dislike = ? WHERE post_id = ?`
		post.Dislike++
		_, err = p.db.Exec(query, post.Dislike, post.ID)
		if err != nil {
			log.Println(err)
			return err
		}

	} else if dislike == 0 && like == 1 {
		query = `DELETE FROM likes WHERE post_id = ? AND username = ?`
		_, err := p.db.Exec(query, postId, username)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `INSERT INTO dislikes(post_id, username, status) VALUES(?,?,?)`
		_, err = p.db.Exec(query, postId, username, 1)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `UPDATE posts SET like = ?, dislike = ? WHERE post_id = ? `
		post.Like--
		post.Dislike++
		_, err = p.db.Exec(query, post.Like, post.Dislike, post.ID)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		query = `DELETE FROM dislikes WHERE post_id = ? AND username = ?`
		_, err := p.db.Exec(query, postId, username)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `UPDATE posts SET dislike = ? WHERE post_id = ?`
		post.Dislike--
		_, err = p.db.Exec(query, post.Like, post.ID)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (p *postQuery) LikeComment(commentId int64, username string) error {
	comment, err := p.GetCommentByCommentID(commentId)
	if err != nil {
		log.Println(err)
		return err
	}
	query := `SELECT status FROM comment_likes WHERE comment_id = ? AND username = ?`
	var like int
	_ = p.db.QueryRow(query, commentId, username).Scan(&like)

	query = `SELECT status FROM comment_dislikes WHERE comment_id = ? AND username = ?`
	var dislike int
	_ = p.db.QueryRow(query, commentId, username).Scan(&dislike)

	if like == 0 && dislike == 0 {
		query = `INSERT INTO comment_likes(comment_id, username, status) VALUES(?,?,?)`
		_, err := p.db.Exec(query, comment.ID, username, 1)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `UPDATE comments SET like = ? WHERE comment_id = ?`
		comment.Like++
		_, err = p.db.Exec(query, comment.Like, comment.ID)
		if err != nil {
			log.Println(err)
			return err
		}
	} else if like == 0 && dislike == 1 {
		query = `DELETE FROM comment_dislikes WHERE comment_id = ? AND username = ?`
		_, err := p.db.Exec(query, commentId, username)
		if err != nil {
			log.Println(err)
			log.Println(err)
			return err
		}
		query = `INSERT INTO comment_likes(comment_id, username, status) VALUES(?,?,?)`
		_, err = p.db.Exec(query, commentId, username, 1)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `UPDATE comments SET like = ?, dislike = ? WHERE comment_id = ?`
		comment.Like++
		comment.Dislike--
		_, err = p.db.Exec(query, comment.Like, comment.Dislike, comment.ID)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		query = `DELETE FROM comment_likes WHERE comment_id = ? AND username = ?`
		_, err := p.db.Exec(query, commentId, username)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `UPDATE comments SET like = ? WHERE comment_id = ?`
		comment.Like--
		_, err = p.db.Exec(query, comment.Like, comment.ID)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (p *postQuery) DislikeComment(commentId int64, username string) error {
	comment, err := p.GetCommentByCommentID(commentId)
	if err != nil {
		log.Println(err)
		return err
	}
	query := `SELECT status FROM comment_likes WHERE comment_id = ? AND username = ?`
	var like int
	_ = p.db.QueryRow(query, commentId, username).Scan(&like)

	query = `SELECT status FROM comment_dislikes WHERE comment_id = ? AND username = ?`
	var dislike int
	_ = p.db.QueryRow(query, commentId, username).Scan(&dislike)

	if like == 0 && dislike == 0 {
		query = `INSERT INTO comment_dislikes(comment_id, username, status) VALUES(?,?,?)`
		_, err := p.db.Exec(query, comment.ID, username, 1)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `UPDATE comments SET dislike = ? WHERE comment_id = ?`
		comment.Dislike++
		_, err = p.db.Exec(query, comment.Dislike, comment.ID)
		if err != nil {
			log.Println(err)
			return err
		}
	} else if like == 1 && dislike == 0 {
		query = `DELETE FROM comment_likes WHERE comment_id = ? AND username = ?`
		_, err := p.db.Exec(query, commentId, username)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `INSERT INTO comment_dislikes(comment_id, username, status) VALUES(?,?,?)`
		_, err = p.db.Exec(query, commentId, username, 1)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `UPDATE comments SET dislike = ?, like = ? WHERE comment_id = ?`
		comment.Like--
		comment.Dislike++
		_, err = p.db.Exec(query, comment.Dislike, comment.Like, comment.ID)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		query = `DELETE FROM comment_dislikes WHERE comment_id = ? AND username = ?`
		_, err := p.db.Exec(query, commentId, username)
		if err != nil {
			log.Println(err)
			return err
		}
		query = `UPDATE comments SET dislike = ? WHERE comment_id = ?`
		comment.Dislike--
		_, err = p.db.Exec(query, comment.Dislike, comment.ID)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (p *postQuery) GetLikedPostIdByUser(user model.User) ([]int64, error) {
	var postId []int64
	query := `SELECT post_id FROM likes WHERE username = ?`
	rows, err := p.db.Query(query, user.Username)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		postId = append(postId, id)
	}
	return postId, nil
}
