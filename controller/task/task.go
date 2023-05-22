package task_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type task struct {
	ID          int    `json:"id"`
	Card_ID     int    `json:"card_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Start       string `json:"start"`
	Percent     int    `json:"percent"`
	Deadline    string `json:"deadline"`
	Completed   bool   `json:"completed"`
}

var tasks = []task{
	{ID: 1, Card_ID: 1, Title: "task1", Description: "task1", Start: "", Percent: 0, Deadline: "", Completed: false},
}

func GetAllTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskByID(id int) (*task, error) {
	for i, t := range tasks {
		if t.ID == id {
			return &tasks[i], nil
		}
	}

	return nil, errors.New("Task not found.")
}

func GetOneTask(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := GetTaskByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)

}

func ChangeTaskCard(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := GetTaskByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Card_ID = newTask.Card_ID
	c.IndentedJSON(http.StatusOK, task)
}

func ChangeTaskTitle(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := GetTaskByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Title = newTask.Title
	c.IndentedJSON(http.StatusOK, task)
}

func ChangeTaskDescription(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := GetTaskByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Description = newTask.Description
	c.IndentedJSON(http.StatusOK, task)
}

func ChangeTaskStart(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := GetTaskByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Start = newTask.Start
	c.IndentedJSON(http.StatusOK, task)
}

func ChangeTaskPercent(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := GetTaskByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Percent = newTask.Percent
	c.IndentedJSON(http.StatusOK, task)
}

func ChangeTaskDeadline(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := GetTaskByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Deadline = newTask.Deadline
	c.IndentedJSON(http.StatusOK, task)
}

func ChangeTaskCompleted(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := GetTaskByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Completed = !task.Completed
	c.IndentedJSON(http.StatusOK, task)
}
