package app

import (
	"log"
	"net/http"
	"strings"
	"time"

	"forum/internal/model"
	"forum/pkg"
)

func (app *App) CreatePostPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	pkg.RenderTemplate(w, "createpost.html", model.Data{})
}

func (app *App) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	title := r.FormValue("title")
	message := r.FormValue("message")
	genre := r.Form["category"]
	if len(genre) == 0 {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	t := app.CheckSpace(title)
	ct := app.CheckSpace(message)
	for _, v := range genre {
		if v != "romance" && v != "adventure" && v != "comedy" && v != "drama" && v != "fantazy" {
			pkg.ErrorHandler(w, http.StatusBadRequest)
			return
		}
	}
	if !t || !ct {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	category := strings.Join(genre, " ")
	c, err := r.Cookie("session_token")
	if err != nil {
		log.Println(err)
	}
	session, err := app.sessionService.GetSessionByToken(c.Value)
	if err != nil {
		log.Println(err)
	}
	post := model.Post{
		Title:    title,
		Content:  message,
		Category: category,
		Author: model.User{
			Email:    session.Email,
			Username: session.Username,
		},
		CreateTime: time.Now().Format(time.RFC822),
	}
	id, err := app.postService.CreatePost(post)
	if err != nil {
		log.Println(err)
	}
	categories := model.Category{
		CategoryName: category,
		PostId:       id,
	}
	err = app.postService.CreateCategory(categories)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
