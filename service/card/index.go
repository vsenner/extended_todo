package card_service

import (
	//"database/sql"
	db "extended_todo/routing"
	"fmt"

	"github.com/pkg/errors"
)

type Card struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	AdminID int    `json:"admin_id"`
}

func GetAll(userID int) ([]Card, error) {

	query := fmt.Sprintf("select * from card where admin_id=%d", userID)

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	var cards []Card

	for rows.Next() {
		var card Card
		err := rows.Scan(&card.ID, &card.Name, &card.AdminID)

		if err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}
		cards = append(cards, card)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error iterating over rows")
	}

	return cards, nil
}

func GetOne(cardID int) Card {

	c := Card{}

	if err := db.DB.QueryRow("SELECT * FROM card WHERE id=$1", cardID).Scan(&c.ID, &c.Name, &c.AdminID); err != nil {
	}

	return c
}

func Add(userID int, name string) Card {

	c := Card{}

	if err := db.DB.QueryRow("INSERT INTO card (name, admin_id) VALUES ($1, $2) RETURNING *", name, userID).Scan(&c.ID, &c.Name, &c.AdminID); err != nil {
	}

	return c
}

func Remove(cardID int) (bool, error) {
	result, err := db.DB.Exec("DELETE FROM card WHERE id=$1", cardID)
	if err != nil {
		return false, errors.Wrap(err, "failed to execute query")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, errors.Wrap(err, "failed to get rows affected")
	}

	if rowsAffected == 0 {
		return false, errors.New("card not found")
	}

	return true, nil
}

func Rename(cardID int, name string) Card {

	c := Card{}

	if err := db.DB.QueryRow("UPDATE card SET name=$1 WHERE id=$2 RETURNING *", name, cardID).Scan(&c.ID, &c.Name, &c.AdminID); err != nil {
	}

	return c
}
