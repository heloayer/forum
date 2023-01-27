package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"forum/internal/app"
	"forum/internal/repository"
	"forum/internal/service"
)

func main() {
	db, err := repository.NewDB() // вызываем функцию для создания и заполнения db
	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Ping() // проверяет соединение с базой данных
	if err != nil {
		log.Fatalf("cannot ping db: %v", err)
	}
	dao := repository.NewDao(db)               // dao - интерфейс с базой данных имеющая нужные данные для удобного доступа к БД
	authService := service.NewAuthService(dao)   // service/auth
	// интерфейсы с базой данных для удобного доступа к БД, реализуются через интерфейс dao   
	sessionService := service.NewSessionService(dao)
	postService := service.NewPostService(dao)
	userService := service.NewUserService(dao)
	app := app.NewAppService(authService, sessionService, postService, userService) // app - структура с полями интерфейсов // вызывается из сервис
	server := app.Run()

	go func() {
		log.Printf("server started at http://localhost%v\n", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Shutting down servers..")
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown: %s\n", err)
	}
	log.Println("Server gracefully stoped")
}
