package main

import (
	"net/http"
	"strconv"

	"errors"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type card struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Admin_ID int    `json:"admin_id"`
}

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

var users = []user{
	{ID: 1, Name: "Nick", Email: "nick@gmail.com", Password: "password"},
}

var cards = []card{
	{ID: 1, Name: "My day", Admin_ID: 1},
}

var tasks = []task{
	{ID: 1, Card_ID: 1, Title: "task1", Description: "task1", Start: "", Percent: 0, Deadline: "", Completed: false},
}

//User's funcs

func getAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUserById(id int) (*user, error) {
	for i, u := range users {
		if u.ID == id {
			return &users[i], nil
		}
	}

	return nil, errors.New("User not found.")
}

func createUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)

}

func getOneUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	user, err := getUserById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func changeUserName(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	user, err := getUserById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	user.Name = newUser.Name
	c.IndentedJSON(http.StatusOK, user)
}

func changeUserEmail(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	user, err := getUserById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	user.Email = newUser.Email
	c.IndentedJSON(http.StatusOK, user)
}

func changeUserPassword(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	user, err := getUserById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	user.Password = newUser.Password
	c.IndentedJSON(http.StatusOK, user)
}

//Cards's funcs

func getAllCards(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cards)
}

func getCardById(id int) (*card, error) {
	for i, c := range cards {
		if c.ID == id {
			return &cards[i], nil
		}
	}

	return nil, errors.New("Card not found.")
}

func createCard(c *gin.Context) {
	var newCard card

	if err := c.BindJSON(&newCard); err != nil {
		return
	}

	cards = append(cards, newCard)
	c.IndentedJSON(http.StatusCreated, newCard)

}

func getOneCard(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	card, err := getCardById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Card not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, card)
}

func renameCard(c *gin.Context) {
	var newCard card

	if err := c.BindJSON(&newCard); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	card, err := getCardById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Card not found"})
		return
	}

	card.Name = newCard.Name
	c.IndentedJSON(http.StatusOK, card)
}

//Task's funcs

func getAllTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskById(id int) (*task, error) {
	for i, t := range tasks {
		if t.ID == id {
			return &tasks[i], nil
		}
	}

	return nil, errors.New("Task not found.")
}

func getOneTask(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
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

func changeTaskCard(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Card_ID = newTask.Card_ID
	c.IndentedJSON(http.StatusOK, task)
}

func changeTaskTitle(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Title = newTask.Title
	c.IndentedJSON(http.StatusOK, task)
}

func changeTaskDescription(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Description = newTask.Description
	c.IndentedJSON(http.StatusOK, task)
}

func changeTaskStart(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Start = newTask.Start
	c.IndentedJSON(http.StatusOK, task)
}

func changeTaskPercent(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Percent = newTask.Percent
	c.IndentedJSON(http.StatusOK, task)
}

func changeTaskDeadline(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Deadline = newTask.Deadline
	c.IndentedJSON(http.StatusOK, task)
}

func changeTaskCompleted(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

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

	//User routes
	router.GET("/users", getAllUsers)
	router.GET("/users/:id", getOneUser)
	router.POST("/users", createUser)
	router.PATCH("/changeUserName/:id", changeUserName)
	router.PATCH("/changeUserEmail/:id", changeUserEmail)
	router.PATCH("/changeUserPassword/:id", changeUserPassword)

	//Card routes
	router.GET("/cards", getAllCards)
	router.GET("/cards/:id", getOneCard)
	router.POST("/cards", createCard)
	router.PATCH("/renameCard/:id", renameCard)

	//Tasks routes
	router.GET("/tasks", getAllTasks)
	router.GET("/tasks/:id", getOneTask)
	router.POST("/tasks", createTask)
	router.PATCH("/changeTaskCard/:id", changeTaskCard)
	router.PATCH("/changeTaskTitle/:id", changeTaskTitle)
	router.PATCH("/changeTaskDescription/:id", changeTaskDescription)
	router.PATCH("/changeTaskStart/:id", changeTaskStart)
	router.PATCH("/changeTaskPercent/:id", changeTaskPercent)
	router.PATCH("/changeTaskDeadline/:id", changeTaskDeadline)
	router.PATCH("/changeTaskComplete/:id", changeTaskCompleted)

	router.Run("localhost:8080")
}
