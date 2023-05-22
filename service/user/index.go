package service

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/pkg/errors"
)

type UserAuthorizedService struct {
	db *sql.DB
}

type UserRes struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserAuthorizedService(db *sql.DB) *UserAuthorizedService {
	return &UserAuthorizedService{
		db: db,
	}
}

func (uas *UserAuthorizedService) ChangeName(userID int, name string) (*UserRes, error) {
	row := uas.db.QueryRow(fmt.Sprintf("UPDATE users SET name='%s' WHERE id=%d RETURNING *", name, userID))

	var user UserRes
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &user, nil
}

func (uas *UserAuthorizedService) ChangePassword(userID int, password string) (*UserRes, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate hashed password")
	}

	row := uas.db.QueryRow(fmt.Sprintf("UPDATE users SET password='%s' WHERE id=%d RETURNING *", hashedPassword, userID))

	var user UserRes
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &user, nil
}

func (uas *UserAuthorizedService) ChangeEmail(userID int, email string) (*UserRes, error) {
	row := uas.db.QueryRow(fmt.Sprintf("UPDATE users SET email='%s' WHERE id=%d RETURNING *", email, userID))

	var user UserRes
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &user, nil
}
