package app

import (
	"net/http"
	"time"

	"forum/pkg"
)

var (
	WelcomeCookieOnPaths = make(map[string]struct{}, 10) // карта с ключами путей unauth. users со значением пустой структуры
	HomeCookieOnPaths    = make(map[string]struct{}, 10) // карта с ключами путей auth. users со значением пустой структуры
)

func AddWelcomeCookieCheckOnPaths(paths ...string) {
	for _, path := range paths {
		WelcomeCookieOnPaths[path] = struct{}{} // struct{} просто тип, чтобы указать что он пусть нужно {}{}
	}
}

func AddHomeCookieCheckOnPaths(paths ...string) {
	for _, path := range paths {
		HomeCookieOnPaths[path] = struct{}{}
	}
}

func (app *App) WelcomeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, ok := WelcomeCookieOnPaths[r.URL.Path]; !ok { // проверяем имеется ли путь в карте, ; ok??
			pkg.ErrorHandler(w, http.StatusNotFound)
			return
		} else {        // если путь действителен
			c, err := r.Cookie("session_token") // присваиваем токен (session token), указанный в запросе через метод Cookie
			if err == http.ErrNoCookie {        // если токена нет
				next.ServeHTTP(w, r) // то возвращаем реализуем request, response обработчика и выходим из функции
				return
			}
			sessionFromDb, err := app.sessionService.GetSessionByToken(c.Value)  // операция ля поля интерфейса app (repository/session)
			             // вызываем GetSessionByToken для знач. Value ("session token") структуры Cookie
			if err != nil {
				next.ServeHTTP(w, r)   // если все ок инициализируем оригинальный обработчик
				return
			}
			if sessionFromDb.Expiry.Before(time.Now()) {
				next.ServeHTTP(w, r)
				return
			}
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}
}

func (app *App) HomeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, ok := HomeCookieOnPaths[r.URL.Path]; !ok {
			pkg.ErrorHandler(w, http.StatusNotFound)
			return
		} else {
			c, err := r.Cookie("session_token")
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/sign-in-form", http.StatusFound)
				return
			}
			sessionFromDb, err := app.sessionService.GetSessionByToken(c.Value)
			if err != nil {
				http.Redirect(w, r, "/sign-in-form", http.StatusFound)
				return
			}
			if sessionFromDb.Expiry.Before(time.Now()) {
				http.Redirect(w, r, "/sign-in-form", http.StatusFound)
				return
			}
			next.ServeHTTP(w, r)
		}
	}
}
