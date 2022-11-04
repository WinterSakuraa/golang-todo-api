package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/wintersakura/golang-vue-todo/pkg/model"
	"strings"
)

type TodoPostgres struct {
	db *sqlx.DB
}

func NewTodoPostgres(db *sqlx.DB) *TodoPostgres {
	return &TodoPostgres{db: db}
}

func (r *TodoPostgres) Create(todo model.Todo) (int, error) {
	var todoId int
	query := fmt.Sprintf(`INSERT INTO %s (task, done) VALUES ($1, $2) RETURNING id;`, todosTable)

	err := r.db.QueryRow(query, todo.Task, todo.Done).Scan(&todoId)
	if err != nil {
		return 0, err
	}

	return todoId, nil
}

func (r *TodoPostgres) GetAll() ([]model.Todo, error) {
	var todos []model.Todo

	query := fmt.Sprintf(`SELECT * FROM %s;`, todosTable)
	if err := r.db.Select(&todos, query); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *TodoPostgres) Delete(todoId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1;`, todosTable)
	_, err := r.db.Exec(query, todoId)

	return err
}

func (r *TodoPostgres) GetById(todoId int) (model.Todo, error) {
	var todo model.Todo

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1;`, todosTable)
	err := r.db.Get(&todo, query, todoId)

	return todo, err
}

func (r *TodoPostgres) Update(todoId int, todo model.UpdateTodoInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if todo.Task != nil {
		setValues = append(setValues, fmt.Sprintf("task = $%d", argId))
		args = append(args, *todo.Task)
		argId++
	}

	if todo.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done = $%d", argId))
		args = append(args, *todo.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d;`, todosTable, setQuery, argId)
	args = append(args, todoId)

	_, err := r.db.Exec(query, args...)
	return err
}
