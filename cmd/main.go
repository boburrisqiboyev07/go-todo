package main

import (
	"log"

	"github.com/boburrisqiboyev07/go-todo.git"
	"github.com/boburrisqiboyev07/go-todo.git/package/handler"

)

func main() {
	handler := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
