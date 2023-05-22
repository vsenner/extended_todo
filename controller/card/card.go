package card_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type card struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Admin_ID int    `json:"admin_id"`
}

var cards = []card{
	{ID: 1, Name: "My day", Admin_ID: 1},
}

func GetAllCards(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cards)
}

func GetCardByID(id int) (*card, error) {
	for i, c := range cards {
		if c.ID == id {
			return &cards[i], nil
		}
	}

	return nil, errors.New("Card not found.")
}

func CreateCard(c *gin.Context) {
	var newCard card

	if err := c.BindJSON(&newCard); err != nil {
		return
	}

	cards = append(cards, newCard)
	c.IndentedJSON(http.StatusCreated, newCard)
}

func GetOneCard(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	card, err := GetCardByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Card not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, card)
}

func RenameCard(c *gin.Context) {
	var newCard card

	if err := c.BindJSON(&newCard); err != nil {
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	card, err := GetCardByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Card not found"})
		return
	}

	card.Name = newCard.Name
	c.IndentedJSON(http.StatusOK, card)
}
