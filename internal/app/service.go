package app

import "forum/internal/service"

type App struct { // структура в кот. будут храниться интерфейсы
	authService    service.AuthService
	sessionService service.SessionService   // интерфейс из service/session
	postService    service.PostService
	userService    service.UserService
}

func NewAppService( // заполняем структуру App полями интерфейсов
	authService service.AuthService,
	sessionService service.SessionService,       // для поля вызывается GetSessionByToken из интерфейса SessionQuery (repository/session)
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
