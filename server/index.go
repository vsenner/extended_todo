package server

import (
	"extended_todo/controller/test"
	user_unauthorized_controller "extended_todo/controller/user-unauthorized"
	token_service "extended_todo/service/token"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

func Server() {
	router := gin.Default()

	//cardGroup := router.Group("/cards")
	//cardGroup.Use(Authenticate)
	//{
	//	cardGroup.GET("", card_controller.GetAllCards)
	//	cardGroup.GET("/:id", card_controller.GetOneCard)
	//}

	router.POST("/user/registration", user_unauthorized_controller.Registration)
	router.POST("/user/login", user_unauthorized_controller.Login)

	authGroup := router.Group("/auth", Authenticate)

	authGroup.GET("/test", test.Test)

	//taskGroup := router.Group("/tasks")
	//taskGroup.Use(Authenticate)
	//{
	//	taskGroup.GET("", task.getAllTasks)
	//	taskGroup.GET("/:id", task.getOneTask)
	//}

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
