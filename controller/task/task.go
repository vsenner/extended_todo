package task_controller

import (
	task_service "extended_todo/service/task"
	"fmt"

	"github.com/gin-gonic/gin"

	//"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type TaskBody struct {
	Card_ID     int    `json:"card_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Start       string `json:"start"`
	Percent     int    `json:"percent"`
	Deadline    string `json:"deadline"`
	Completed   bool   `json:"completed"`
}

func GetAllTasks(c *gin.Context) {
	cardIdStr := c.Param("cardID")
	cardId, _ := strconv.Atoi(cardIdStr)
	tasks, err := task_service.GetAll(cardId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Tasks not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func GetOneTask(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task := task_service.GetOne(id)

	fmt.Print(task)

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func CreateTask(c *gin.Context) {
	var newTask TaskBody

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	task := task_service.Add(newTask.Card_ID, newTask.Title, newTask.Description, newTask.Start, newTask.Deadline, newTask.Percent)

	fmt.Print(task)

	c.JSON(http.StatusOK, gin.H{"task": task})

}

func ChangeTaskCard(c *gin.Context) {
	var newTask TaskBody

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task := task_service.ChangeCard(id, newTask.Card_ID)

	fmt.Print(task)

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func ChangeTaskTitle(c *gin.Context) {
	var newTask TaskBody

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task := task_service.ChangeTitle(id, newTask.Title)

	fmt.Print(task)

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func ChangeTaskDescription(c *gin.Context) {
	var newTask TaskBody

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task := task_service.ChangeDescription(id, newTask.Description)

	fmt.Print(task)

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func ChangeTaskStart(c *gin.Context) {
	var newTask TaskBody

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task := task_service.ChangeStart(id, newTask.Start)

	fmt.Print(task)

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func ChangeTaskPercent(c *gin.Context) {
	var newTask TaskBody

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task := task_service.ChangePercent(id, newTask.Percent)

	fmt.Print(task)

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func ChangeTaskDeadline(c *gin.Context) {
	var newTask TaskBody

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task := task_service.ChangeDeadline(id, newTask.Deadline)

	fmt.Print(task)

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func ChangeTaskCompleted(c *gin.Context) {
	var newTask TaskBody

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task := task_service.ChangeComplete(id, newTask.Completed)

	fmt.Print(task)

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func RemoveTask(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	result, err := task_service.Remove(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}
