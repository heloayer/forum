package app

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"forum/internal/model"
	"forum/pkg"
)

func (app *App) CommentWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println(err)
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	comments, err := app.postService.GetAllCommentByPostID(int64(id))
	if err != nil {
		log.Println(err)
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	post, err := app.postService.GetPostById(int64(id))
	if err != nil {
		log.Println(err)
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	data := model.Data{
		Comment:     comments,
		InitialPost: post,
	}
	pkg.RenderTemplate(w, "commentunauth.html", data)
}

func (app *App) CommentViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println(err)
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	comments, err := app.postService.GetAllCommentByPostID(int64(id))
	if err != nil {
		log.Println(err)
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	post, err := app.postService.GetPostById(int64(id))
	if err != nil {
		log.Println(err)
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	data := model.Data{
		Comment:     comments,
		InitialPost: post,
	}
	pkg.RenderTemplate(w, "commentview.html", data)
}

func (app *App) CommentPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	path := "/comment-view?id=" + (r.URL.Query().Get("id"))
	c, err := r.Cookie("session_token")
	if err != nil {
		log.Println(err)
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	session, err := app.sessionService.GetSessionByToken(c.Value)
	if err != nil {
		log.Println(err)
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	comment := r.FormValue("comment")
	b := app.CheckSpace(comment)
	if len(comment) == 0 || !b {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	user, err := app.userService.GetUserByEmail(session.Email)
	if err != nil {
		log.Println(err)
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	com := model.Comment{
		PostId:   int64(id),
		Message:  comment,
		Username: user.Username,
		Born:     time.Now().Format(time.RFC822),
	}
	if err := app.postService.CommentPost(com); err != nil {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, path, http.StatusFound)
}
