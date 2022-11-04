package service

import (
	"github.com/wintersakura/golang-vue-todo/pkg/model"
	"github.com/wintersakura/golang-vue-todo/pkg/repository"
)

type Todo interface {
	Create(todo model.Todo) (int, error)
	GetAll() ([]model.Todo, error)
	Delete(todoId int) error
	GetById(todoId int) (model.Todo, error)
	Update(todoId int, todo model.UpdateTodoInput) error
}

type Service struct {
	Todo
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Todo: NewTodoService(repo.Todo),
	}
}
