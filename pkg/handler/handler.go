package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wintersakura/golang-vue-todo/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())

	api := router.Group("/api/v1")
	{
		todos := api.Group("/todos")
		{
			todos.POST("/", h.createTodo)
			todos.GET("/", h.getAllTodos)
			todos.GET("/:id", h.getTodoById)
			todos.PUT("/:id", h.updateTodo)
			todos.DELETE("/:id", h.deleteTodo)
		}
	}

	return router
}
