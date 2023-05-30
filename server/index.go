package server

import (
	card_controller "extended_todo/controller/card"
	user_unauthorized_controller "extended_todo/controller/user-unauthorized"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

func Server() {
	router := gin.Default()

	router.POST("/login", login)

	//userGroup := router.Group("/users")
	//userGroup.Use(Authenticate)
	//{
	//	userGroup.GET("", user_controller.Registration)
	//	userGroup.GET("/:id", user_controller.GetOneUser)
	//}

	cardGroup := router.Group("/cards")
	cardGroup.Use(Authenticate)
	{
		cardGroup.GET("", card_controller.GetAllCards)
		cardGroup.GET("/:id", card_controller.GetOneCard)
	}

	router.POST("/user/registration", user_unauthorized_controller.Registration)

	//taskGroup := router.Group("/tasks")
	//taskGroup.Use(Authenticate)
	//{
	//	taskGroup.GET("", task.getAllTasks)
	//	taskGroup.GET("/:id", task.getOneTask)
	//}

	router.Run("localhost:8080")
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	tokenString = tokenString[7:]

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Set("username", claims.Username)
	c.Next()
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c *gin.Context) {

	var user UserLogin

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	if user.Username != "admin" || user.Password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	tokenString, err := CreateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
