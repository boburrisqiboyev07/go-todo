package service

import "github.com/boburrisqiboyev07/go-todo.git/package/repository"

type Authorization interface {
}
type TodoList interface {
}
type TodoItem interface {
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
