package app

import (
	"net/http"
	"os"
	"time"
)

func (app *App) Run() http.Server {
	welcomePaths := []string{ 
		"/welcome/filter", 
		"/sign-in-form",
		"/sign-up-form",
		"/sign-in",
		"/sign-up",
		"/welcome",
		"/welcome/comment-view",
	}
	homePaths := []string{
		"/", 
		"/create-post-form",
		"/comment-post",
		"/like-post",
		"/dislike-post",
		"/like-comment",
		"/dislike-comment",
		"/comment-view",
		"/filter",
		"/create-post",
		"/logout",
	}
	AddWelcomeCookieCheckOnPaths(welcomePaths...)
	AddHomeCookieCheckOnPaths(homePaths...)     

	mux := http.NewServeMux()                               
	mux.HandleFunc("/", app.HomeMiddleware(app.HomeHandler)) 
	mux.HandleFunc("/welcome", app.WelcomeMiddleware(app.WelcomeHandler))
	mux.HandleFunc("/sign-in-form", app.WelcomeMiddleware(app.SignInPageHandler))
	mux.HandleFunc("/sign-up-form", app.WelcomeMiddleware(app.SignUpPageHandler))
	mux.HandleFunc("/create-post-form", app.HomeMiddleware(app.CreatePostPageHandler))
	mux.HandleFunc("/comment-post", app.HomeMiddleware(app.CommentPostHandler))
	mux.HandleFunc("/like-post", app.HomeMiddleware(app.LikePostHandler))
	mux.HandleFunc("/dislike-post", app.HomeMiddleware(app.DislikePostHandler))
	mux.HandleFunc("/like-comment", app.HomeMiddleware(app.LikeCommentHandler))
	mux.HandleFunc("/dislike-comment", app.HomeMiddleware(app.DislikeCommentHandler))
	mux.HandleFunc("/comment-view", app.HomeMiddleware(app.CommentViewHandler))
	mux.HandleFunc("/welcome/comment-view", app.WelcomeMiddleware(app.CommentWelcomeHandler))
	mux.HandleFunc("/filter", app.HomeMiddleware(app.FilterHandler))
	mux.HandleFunc("/logout", app.HomeMiddleware(app.LogoutHandler))
	mux.HandleFunc("/create-post", app.HomeMiddleware(app.CreatePostHandler))
	mux.HandleFunc("/sign-in", app.WelcomeMiddleware(app.SignInHandler))
	mux.HandleFunc("/sign-up", app.WelcomeMiddleware(app.SignUpHandler))
	mux.HandleFunc("/welcome/filter", app.WelcomeMiddleware(app.FilterWelcomeHandler))

	fs := http.FileServer(http.Dir("./templates/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	port, ok := os.LookupEnv("FORUM_PORT")
	if !ok {
		port = ":8080"
	}
	server := http.Server{
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 30,
		Addr:         port,
		Handler:      mux,
	}

	return server
}
