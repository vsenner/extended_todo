package server

import (
	"database/sql"
	card_controller "extended_todo/controller/card"
	user_controller "extended_todo/controller/user-authorized"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var DB *sql.DB

func Server() {
	router := gin.Default()

	dbConnStr := "postgres://{YOUR_NAME}}:{YOUR_PASSWORD}@localhost:5432/extended_todo?sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router.POST("/login", login)

	userGroup := router.Group("/users")
	userGroup.Use(Authenticate)
	{
		userGroup.GET("", user_controller.GetAllUsers)
		userGroup.GET("/:id", user_controller.GetOneUser)
	}

	cardGroup := router.Group("/cards")
	cardGroup.Use(Authenticate)
	{
		cardGroup.GET("", card_controller.GetAllCards)
		cardGroup.GET("/:id", card_controller.GetOneCard)
	}

	//taskGroup := router.Group("/tasks")
	//taskGroup.Use(Authenticate)
	//{
	//	taskGroup.GET("", task.getAllTasks)
	//	taskGroup.GET("/:id", task.getOneTask)
	//}

	err = router.Run("localhost:8080")
	if err != nil {
		return
	}
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

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username != "admin" || password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	tokenString, err := CreateToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
