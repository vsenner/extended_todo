package card_controller

import (
	card_service "extended_todo/service/card"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CardBody struct {
	Name     string `json:"name"`
	Admin_ID int    `json:"admin_id"`
}

func GetAllCards(c *gin.Context) {
	adminIdStr := c.Query("adminID")
	adminId, _ := strconv.Atoi(adminIdStr)

	cards, err := card_service.GetAll(adminId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Cards not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, cards)
}

func CreateCard(c *gin.Context) {
	var newCard CardBody

	if err := c.BindJSON(&newCard); err != nil {
		return
	}

	card := card_service.Add(newCard.Admin_ID, newCard.Name)

	c.JSON(http.StatusOK, gin.H{"card": card})

}

func GetOneCard(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	card := card_service.GetOne(id)

	if (card == card_service.Card{}) {
		c.JSON(http.StatusOK, gin.H{"status": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"card": card})
}

func RenameCard(c *gin.Context) {
	var newCard CardBody

	if err := c.BindJSON(&newCard); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	card := card_service.Rename(id, newCard.Name)

	fmt.Print(card)

	c.JSON(http.StatusOK, gin.H{"card": card})
}

func RemoveCard(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	result, err := card_service.Remove(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Card not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}
