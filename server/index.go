package server

import (
	card_controller "extended_todo/controller/card"
	task_controller "extended_todo/controller/task"
	"extended_todo/controller/test"
	user_unauthorized_controller "extended_todo/controller/user-unauthorized"
	token_service "extended_todo/service/token"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Server() {
	router := gin.Default()

	router.POST("/user/registration", user_unauthorized_controller.Registration)
	router.POST("/user/login", user_unauthorized_controller.Login)

	authGroup := router.Group("/auth", Authenticate)

	authGroup.GET("/test", test.Test)
	authGroup.GET("/cards", card_controller.GetAllCards)
	authGroup.GET("/cards/:id", card_controller.GetOneCard)
	authGroup.POST("/cards", card_controller.CreateCard)
	authGroup.PATCH("/cards/rename/:id", card_controller.RenameCard)
	authGroup.DELETE("/cards/:id", card_controller.RemoveCard)

	authGroup.GET("/tasks", task_controller.GetAllTasks)
	authGroup.GET("/tasks/:id", task_controller.GetOneTask)
	authGroup.POST("/tasks", task_controller.CreateTask)
	authGroup.PATCH("/tasks/change_card/:id", task_controller.ChangeTaskCard)
	authGroup.PATCH("/tasks/change_completed/:id", task_controller.ChangeTaskCompleted)
	authGroup.PATCH("/tasks/change_deadline/:id", task_controller.ChangeTaskDeadline)
	authGroup.PATCH("/tasks/change_description/:id", task_controller.ChangeTaskDescription)
	authGroup.PATCH("/tasks/change_percent/:id", task_controller.ChangeTaskPercent)
	authGroup.PATCH("/tasks/change_start/:id", task_controller.ChangeTaskStart)
	authGroup.PATCH("/tasks/change_title/:id", task_controller.ChangeTaskTitle)
	authGroup.DELETE("/tasks/:id", task_controller.RemoveTask)

	router.Run("localhost:8080")
}

func Authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	tokenString = tokenString[7:]

	fmt.Print("Token - " + tokenString + "\n")

	token, err := jwt.ParseWithClaims(tokenString, &token_service.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(*token_service.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Set("username", claims.Email)
	c.Next()
}
