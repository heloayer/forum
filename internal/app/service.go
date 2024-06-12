package app

import "forum/internal/service"

type App struct { 
	authService    service.AuthService
	sessionService service.SessionService   
	postService    service.PostService
	userService    service.UserService
}

func NewAppService( 
	authService service.AuthService,
	sessionService service.SessionService,    
	postService service.PostService,
	userService service.UserService,
) App {
	return App{
		authService:    authService,
		sessionService: sessionService,
		postService:    postService,
		userService:    userService,
	}
}
