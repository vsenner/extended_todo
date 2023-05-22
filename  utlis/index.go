package _utlis

import (
	"errors"
	"extended_todo/server"
)

func CheckUserCandidate(parameter interface{}) (bool, error) {
	var query string
	var args []interface{}

	switch parameter := parameter.(type) {
	case string:
		query = "SELECT COUNT(*) FROM users WHERE email = $1"
		args = []interface{}{parameter}
	case int:
		query = "SELECT COUNT(*) FROM users WHERE id = $1"
		args = []interface{}{parameter}
	default:
		return false, errors.New("Invalid parameter type")
	}

	var count int
	err := server.DB.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return false, err
	}

	return count == 1, nil
}
