package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wintersakura/golang-vue-todo/pkg/model"
	"net/http"
	"strconv"
)

func (h *Handler) createTodo(c *gin.Context) {
	var todo model.Todo

	if err := c.BindJSON(&todo); err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	todoId, err := h.service.Todo.Create(todo)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": todoId,
	})
}

func (h *Handler) getAllTodos(c *gin.Context) {
	todos, err := h.service.Todo.GetAll()
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (h *Handler) getTodoById(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := h.service.Todo.GetById(todoId)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *Handler) updateTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	var todoInput model.UpdateTodoInput
	if err := c.BindJSON(&todoInput); err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Todo.Update(todoId, todoInput); err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Todo.Delete(todoId); err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
