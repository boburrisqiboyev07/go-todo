package main

import (
	"log"

	"github.com/boburrisqiboyev07/go-todo.git"
	"github.com/boburrisqiboyev07/go-todo.git/package/handler"
	"github.com/boburrisqiboyev07/go-todo.git/package/repository"
	"github.com/boburrisqiboyev07/go-todo.git/package/service"

)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
