package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var tasks = []string{"Task 1", "Task 2", "Task 3"}

func main() {
	r := gin.Default()
	r.Use(AuthMiddleware)

	v1 := r.Group("/api/v1/tasks")
	{
		v1.GET("/", GetTasks)
		v1.POST("/", CreateTask)
		v1.PUT("/:id", UpdateTask)
		v1.DELETE("/:id", DeleteTask)
	}

	r.Run(":8080")
}

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task string
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks = append(tasks, task)
	c.JSON(http.StatusCreated, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var newTask string
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	i := getTaskIndex(id)
	if i == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	tasks[i] = newTask
	c.JSON(http.StatusOK, newTask)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	i := getTaskIndex(id)
	if i == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	tasks = append(tasks[:i], tasks[i+1:]...)
	c.Status(http.StatusNoContent)
}

func getTaskIndex(id string) int {
	for i, task := range tasks {
		if task == id {
			return i
		}
	}
	return -1
}

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "my-secret-token" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}
