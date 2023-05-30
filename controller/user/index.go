package controller

import (
	"net/http"

	userServicePackage "extended_todo/service/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	userService userServicePackage.UserRes
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func Registration(c *gin.Context) {
	var userReq service.UserReq
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userService.Registration(userReq.Email, userReq.Password, userReq.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("refreshToken", user.Tokens.RefreshToken, 30*24*60*60, "/", "", true, true)
	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var userReq service.UserReq
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userService.Login(userReq.Email, userReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("refreshToken", user.Tokens.RefreshToken, 30*24*60*60, "/", "", true, true)
	c.JSON(http.StatusOK, user)
}

func Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userService.Refresh(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("refreshToken", user.Tokens.RefreshToken, 30*24*60*60, "/", "", true, true)
	c.JSON(http.StatusOK, user)
}
