package app

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"forum/pkg"
)

func (app *App) LikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	path := r.URL.Query().Get("path")
	if _, ok := HomeCookieOnPaths[r.URL.Path]; !ok {
		path = "/"
	}
	post, err := app.postService.GetPostById(int64(id))
	if err != nil {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	c, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/sign-in-form", http.StatusFound)
		return
	}
	session, err := app.sessionService.GetSessionByToken(c.Value)
	if err != nil {
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = app.postService.LikePost(post.ID, session.Username)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, path, http.StatusFound)
}

func (app *App) DislikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	path := r.URL.Query().Get("path")
	if _, ok := HomeCookieOnPaths[r.URL.Path]; !ok {
		path = "/"
	}
	post, err := app.postService.GetPostById(int64(id))
	if err != nil {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	c, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/sign-in-form", http.StatusFound)
		return
	}
	session, err := app.sessionService.GetSessionByToken(c.Value)
	if err != nil {
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = app.postService.DislikePost(post.ID, session.Username)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, path, http.StatusFound)
}

func (app *App) LikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	comment, err := app.postService.GetCommentByCommentID(int64(id))
	if err != nil {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	c, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/sign-in-form", http.StatusFound)
		return
	}
	session, err := app.sessionService.GetSessionByToken(c.Value)
	if err != nil {
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = app.postService.LikeComment(comment.ID, session.Username)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, ("/comment-view?id=")+strconv.Itoa(int(comment.PostId)), http.StatusFound)
}

func (app *App) DislikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	comment, err := app.postService.GetCommentByCommentID(int64(id))
	if err != nil {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	c, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/sign-in-form", http.StatusFound)
		return
	}
	session, err := app.sessionService.GetSessionByToken(c.Value)
	if err != nil {
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = app.postService.DislikeComment(comment.ID, session.Username)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, ("/comment-view?id=")+strconv.Itoa(int(comment.PostId)), http.StatusFound)
}
