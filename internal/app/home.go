package app

import (
	"log"
	"net/http"

	"forum/internal/model"
	"forum/pkg"
)

var Messages model.Data

func (app *App) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		pkg.ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	all, err := app.postService.GetAllPosts()
	if err != nil {
		log.Println(err)
	}
	c, err := r.Cookie("session_token")
	if err != nil {
		log.Println(err)
	}
	session, err := app.sessionService.GetSessionByToken(c.Value)
	if err != nil {
		log.Println(err)
	}
	user, err := app.userService.GetUserByEmail(session.Email)
	if err != nil {
		log.Println(err)
		pkg.ErrorHandler(w, 500)
		return
	}
	user.Password = ""
	data := model.Data{
		Posts: all,
		User:  user,
		Genre: "/",
	}
	pkg.RenderTemplate(w, "index.html", data)
}

func (app *App) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	all, err := app.postService.GetAllPosts()
	if err != nil {
		log.Println(err)
	}
	data := model.Data{
		Posts: all,
	}
	pkg.RenderTemplate(w, "unauth.html", data)
}
