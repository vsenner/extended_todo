package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []user{
	{ID: 1, Name: "Nick", Email: "nick@gmail.com", Password: "password"},
}

//User's funcs

func GetAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func GetUserByID(id int) (*user, error) {
	for i, u := range users {
		if u.ID == id {
			return &users[i], nil
		}
	}

	return nil, errors.New("User not found.")
}

func CreateUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func GetOneUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	user, err := GetUserByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func ChangeUserName(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	user, err := GetUserByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	user.Name = newUser.Name
	c.IndentedJSON(http.StatusOK, user)
}

func ChangeUserEmail(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	user, err := GetUserByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	user.Email = newUser.Email
	c.IndentedJSON(http.StatusOK, user)
}

func ChangeUserPassword(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	user, err := GetUserByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	user.Password = newUser.Password
	c.IndentedJSON(http.StatusOK, user)
}
