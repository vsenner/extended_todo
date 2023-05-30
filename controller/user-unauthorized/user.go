package user_controller

import (
	"database/sql"
	db "extended_todo/routing"
	token_service "extended_todo/service/token"
	user_unauthorized_service "extended_todo/service/user-unauthorized"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type UserBodyLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Registration(c *gin.Context) {
	var newUser UserBody

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	var candidate user_unauthorized_service.UserFull
	err := db.DB.QueryRow("SELECT * from users WHERE email=$1;", newUser.Email).Scan(&candidate.Id, &candidate.Name, &candidate.Email, &candidate.Password)

	if err != nil && err == sql.ErrNoRows {

	}

	if candidate.Email == newUser.Email {
		c.JSON(http.StatusBadRequest, gin.H{"message": "USER EXIST"})
		return
	}

	user := user_unauthorized_service.Registration(newUser.Email, newUser.Password, newUser.Name)

	token, _ := token_service.CreateToken(user.Email)

	is_generated := token_service.SaveToken(token, user.Id)

	if !is_generated {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Server Token Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": UserDto{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, "token": token})
}

func Login(c *gin.Context) {
	var body UserBody

	if err := c.BindJSON(&body); err != nil {
		return
	}

	var user user_unauthorized_service.UserFull
	err := db.DB.
		QueryRow("SELECT * from users WHERE email=$1 and password=$2;", body.Email, body.Password).
		Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	if err != nil && err == sql.ErrNoRows {

	}

	if user.Email != body.Email {
		c.JSON(http.StatusBadRequest, gin.H{"message": "INCORRECT DATA"})
		return
	}

	token, _ := token_service.CreateToken(user.Email)

	is_updated := token_service.UpdateToken(token, user.Id)

	if !is_updated {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Server Token update Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": UserDto{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, "token": token})
}
