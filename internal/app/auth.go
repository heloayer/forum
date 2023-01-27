package app

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"forum/internal/model"
	"forum/pkg"
	"github.com/google/uuid"
)

func (app *App) SignInPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	pkg.RenderTemplate(w, "signin.html", Messages)
	Messages.Message = ""
}

func (app *App) SignUpPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	pkg.RenderTemplate(w, "signup.html", Messages)
	Messages.Message = ""
}

func (app *App) SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	newUser, err := app.GetUser(r)
	if err != nil {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	sessionFromDb, err := app.sessionService.GetSessionByEmail(newUser.Email)
	if err != nil {
		log.Println(err)
	}
	if sessionFromDb.Email == newUser.Email {
		err := app.sessionService.DeleteSession(sessionFromDb.Token)
		if err != nil {
			log.Println(err)
		}
	}

	sessionToken := uuid.NewString()
	expiryAt := time.Now().Add(600 * time.Second)
	session := model.Session{
		Email:  newUser.Email,
		Token:  sessionToken,
		Expiry: expiryAt,
	}
	_, err = app.authService.SignIn(newUser, session)
	if err != nil {
		log.Println("new user sign in was failed")
		Messages.Message = "incorrect data"
		http.Redirect(w, r, "/sign-in-form", http.StatusFound)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiryAt,
	})
	log.Println("new user sign in was successfully")
	http.Redirect(w, r, "/", http.StatusFound)
}

func (app *App) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	newUser, err := app.GetUser(r)
	if err != nil {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	uname := app.CheckSpace(newUser.Username)
	uemail := app.CheckSpace(newUser.Email)
	upass := app.CheckSpace(newUser.Password)
	valid := app.CheckValid(newUser)
	if !uemail || !uname || !upass || !valid {
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
	err = app.authService.SignUp(newUser)
	if err != nil {
		log.Println("new user sign up was failed")
		Messages.Message = "user exists"
		http.Redirect(w, r, "/sign-up-form", http.StatusFound)
		return
	} else {
		http.Redirect(w, r, "/sign-in-form", http.StatusFound)
		return
	}
}

func (app *App) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	c, err := r.Cookie("session_token")
	if err != nil {
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = app.authService.Logout(c.Value)
	if err != nil {
		log.Println(err)
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/welcome", http.StatusFound)
}

func (app *App) GetUser(r *http.Request) (model.User, error) {
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	if r.URL.Path == "/sign-in" && (len(email) == 0 || len(password) == 0) {
		return model.User{}, errors.New("invalid form")
	} else if r.URL.Path == "/sign-up" && (len(email) == 0 || len(password) == 0 || len(username) == 0) {
		return model.User{}, errors.New("invalid form")
	}
	return model.User{
		Email:    email,
		Username: username,
		Password: password,
	}, nil
}

func (app *App) CheckSpace(s string) bool {
	str := strings.TrimSpace(s)
	if len(str) == 0 {
		return false
	}
	return true
}

func (app *App) CheckValid(user model.User) bool {
	if len(user.Username) < 6 {
		return false
	}
	for _, v := range user.Username {
		if v < rune(33) {
			return false
		}
	}

	if len(user.Email) < 6 {
		return false
	}

	b := strings.Contains(user.Email, "@")
	if !b {
		return false
	}
	for _, v := range user.Email {
		if v < rune(33) {
			return false
		}
	}

	if len(user.Password) < 8 {
		return false
	}
	for _, v := range user.Password {
		if v < rune(33) {
			return false
		}
	}

	return true
}
