package token_service

import (
	db "extended_todo/routing"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func SaveToken(token string, userID int64) bool {
	row := db.DB.QueryRow("INSERT INTO token (user_id, refresh_token) VALUES ($1, $2) RETURNING *", userID, token)

	if row == nil {
		return false
	}

	return true
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func UpdateToken(token string, user_id int64) bool {
	row := db.DB.QueryRow("UPDATE token SET refresh_token = $1 WHERE user_id = $2", token, user_id)
	if row == nil {
		return false
	}

	return true
}
