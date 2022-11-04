package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/wintersakura/golang-vue-todo/pkg/model"
)

type Todo interface {
	Create(todo model.Todo) (int, error)
	GetAll() ([]model.Todo, error)
	Delete(todoId int) error
	GetById(todoId int) (model.Todo, error)
	Update(todoId int, todo model.UpdateTodoInput) error
}

type Repository struct {
	Todo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Todo: NewTodoPostgres(db),
	}
}
