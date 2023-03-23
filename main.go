package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Deadline    string `json:"deadline"`
	Completed   bool   `json:"compteted"`
	Description string `json:"description"`
}

var tasks = []task{
	{ID: "1", Name: "Finish GO task", Deadline: "22.03.2023 15:00", Completed: false, Description: "Finish creating a GO program"},
	{ID: "2", Name: "Make dinner", Deadline: "22.03.2023 18:30", Completed: false, Description: "Make a dinner"},
	{ID: "3", Name: "Prepare to test", Deadline: "21.03.2023 12:00", Completed: true, Description: "Prepare to a tomorrow test"},
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskById(id string) (*task, error) {
	for i, t := range tasks {
		if t.ID == id {
			return &tasks[i], nil
		}
	}

	return nil, errors.New("Task not found.")
}

func taskById(c *gin.Context) {
	id := c.Param("id")
	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func createTask(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)

}

func updateStatus(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Completed = !task.Completed

	c.IndentedJSON(http.StatusOK, task)
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", taskById)
	router.POST("/tasks", createTask)
	router.PATCH("/updatestatus", updateStatus)
	router.Run("localhost:8080")
}
