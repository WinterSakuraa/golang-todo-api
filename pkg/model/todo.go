package model

import "github.com/pkg/errors"

type Todo struct {
	Id   int    `json:"id" db:"id"`
	Task string `json:"task" db:"task" binding:"required,min=3"`
	Done bool   `json:"done" db:"done"`
}

type UpdateTodoInput struct {
	Task *string `json:"task"`
	Done *bool   `json:"done"`
}

func (i *UpdateTodoInput) Validate() error {
	if i.Task == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
