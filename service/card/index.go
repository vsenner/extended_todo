package card_service

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type CardService struct {
	db *sql.DB
}

type Card struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	AdminID int    `json:"admin_id"`
}

func NewCardService(db *sql.DB) *CardService {
	return &CardService{
		db: db,
	}
}

func (cs *CardService) GetAll(userID int) ([]Card, error) {
	rows, err := cs.db.Query(fmt.Sprintf("SELECT * FROM card WHERE admin_id=%d", userID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}
	defer rows.Close()

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

func (cs *CardService) GetOne(cardID int) (*Card, error) {
	row := cs.db.QueryRow(fmt.Sprintf("SELECT * FROM card WHERE id=%d", cardID))

	var card Card
	err := row.Scan(&card.ID, &card.Name, &card.AdminID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("card not found")
		}
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &card, nil
}

func (cs *CardService) Add(userID int, name string) (*Card, error) {
	row := cs.db.QueryRow(fmt.Sprintf("INSERT INTO card (name, admin_id) VALUES ('%s', %d) RETURNING *", name, userID))

	var card Card
	err := row.Scan(&card.ID, &card.Name, &card.AdminID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &card, nil
}

func (cs *CardService) Remove(cardID int) (bool, error) {
	result, err := cs.db.Exec(fmt.Sprintf("DELETE FROM card WHERE id=%d", cardID))
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

func (cs *CardService) Rename(cardID int, name string) (*Card, error) {
	row := cs.db.QueryRow(fmt.Sprintf("UPDATE card SET name='%s' WHERE id=%d RETURNING *", name, cardID))

	var card Card
	err := row.Scan(&card.ID, &card.Name, &card.AdminID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &card, nil
}
