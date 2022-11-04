package service

import (
	"github.com/wintersakura/golang-vue-todo/pkg/model"
	"github.com/wintersakura/golang-vue-todo/pkg/repository"
)

type TodoService struct {
	repo repository.Todo
}

func NewTodoService(repo repository.Todo) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) Create(todo model.Todo) (int, error) {
	return s.repo.Create(todo)
}

func (s *TodoService) GetAll() ([]model.Todo, error) {
	return s.repo.GetAll()
}

func (s *TodoService) Delete(todoId int) error {
	return s.repo.Delete(todoId)
}

func (s *TodoService) GetById(todoId int) (model.Todo, error) {
	return s.repo.GetById(todoId)
}

func (s *TodoService) Update(todoId int, todo model.UpdateTodoInput) error {
	if err := todo.Validate(); err != nil {
		return err
	}

	return s.repo.Update(todoId, todo)
}
