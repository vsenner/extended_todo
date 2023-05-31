package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "TEST AUTH"})
}
