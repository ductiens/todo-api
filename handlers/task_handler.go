package handlers

import (
	"net/http"
	"strconv"

	"github.com/ductiens/todo-api/models"
	"github.com/gin-gonic/gin"
)

var tasks = []models.Task{}
var nextID = 1

// CreateTask - POST /tasks
func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = nextID
	nextID++
	newTask.Completed = false

	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, gin.H{"id": newTask.ID})

}

// GetTasks - GET /tasks
func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

// UpdateTaskStatus - PUT /tasks/:id
func UpdateTaskStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updateData struct {
		Completed bool `json:"completed"`
	}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = updateData.Completed
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// DeleteTask - DELETE /tasks/:id
func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}