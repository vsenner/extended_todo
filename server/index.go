package server

import (
	card_controller "extended_todo/controller/card"
	task_controller "extended_todo/controller/task"
	"extended_todo/controller/test"
	user_unauthorized_controller "extended_todo/controller/user-unauthorized"
	token_service "extended_todo/service/token"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Server() {
	router := gin.Default()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	router.Use(CORSMiddleware())

	router.POST("/user/registration", user_unauthorized_controller.Registration)
	router.POST("/user/login", user_unauthorized_controller.Login)

	authGroup := router.Group("/auth", Authenticate)

	authGroup.GET("/test", test.Test)
	authGroup.GET("/cards", card_controller.GetAllCards)
	authGroup.GET("/cards/:id", card_controller.GetOneCard)
	authGroup.POST("/cards", card_controller.CreateCard)
	authGroup.PATCH("/cards/rename/:id", card_controller.RenameCard)
	authGroup.DELETE("/cards/remove/:id", card_controller.RemoveCard)

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

	router.Run("0.0.0.0" + ":" + port)
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
