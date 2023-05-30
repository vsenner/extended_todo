package user_controller

import (
	user_unauthorized_service "extended_todo/service/user-unauthorized"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//"strconv"
)

type UserBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//User's funcs

func Registration(c *gin.Context) {
	var newUser UserBody

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	//if user.Username != "admin" || user.Password != "admin" {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	//	return
	//}
	//
	//tokenString, err := CreateToken(user.Username)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
	//	return
	//}

	//c.JSON(http.StatusOK, gin.H{"token": tokenString})

	user := user_unauthorized_service.Registration(newUser.Email, newUser.Password, newUser.Name)

	fmt.Print(user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}
