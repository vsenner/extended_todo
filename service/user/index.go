package service

import (
	db "extended_todo/routing"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserRes struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ChangeName(userID int, name string) (*UserRes, error) {
	row := db.DB.QueryRow(fmt.Sprintf("UPDATE users SET name='%s' WHERE id=%d RETURNING *", name, userID))

	var user UserRes
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &user, nil
}

func ChangePassword(userID int, password string) (*UserRes, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate hashed password")
	}

	row := db.DB.QueryRow(fmt.Sprintf("UPDATE users SET password='%s' WHERE id=%d RETURNING *", hashedPassword, userID))

	var user UserRes
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &user, nil
}

func ChangeEmail(userID int, email string) (*UserRes, error) {
	row := db.DB.QueryRow(fmt.Sprintf("UPDATE users SET email='%s' WHERE id=%d RETURNING *", email, userID))

	var user UserRes
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &user, nil
}
