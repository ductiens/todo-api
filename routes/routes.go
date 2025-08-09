package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ductiens/todo-api/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetTasks)
	r.PUT("/tasks/:id", handlers.UpdateTaskStatus)
	r.DELETE("/tasks/:id", handlers.DeleteTask)
}