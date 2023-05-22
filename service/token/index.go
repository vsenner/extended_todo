package token_service

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type TokenService struct {
	refreshTokenSecret string
	accessTokenSecret  string
	db                 *sql.DB
}

type UserRes struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewTokenService(refreshTokenSecret string, accessTokenSecret string, db *sql.DB) *TokenService {
	return &TokenService{
		refreshTokenSecret: refreshTokenSecret,
		accessTokenSecret:  accessTokenSecret,
		db:                 db,
	}
}

func (ts *TokenService) GenerateTokens(payload *UserRes) (string, string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["id"] = payload.ID
	refreshClaims["name"] = payload.Name
	refreshClaims["email"] = payload.Email
	refreshClaims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	refreshTokenString, err := refreshToken.SignedString([]byte(ts.refreshTokenSecret))
	if err != nil {
		return "", "", errors.Wrap(err, "failed to generate refresh token")
	}

	accessToken := jwt.New(jwt.SigningMethodHS256)
	accessClaims := accessToken.Claims.(jwt.MapClaims)
	accessClaims["id"] = payload.ID
	accessClaims["name"] = payload.Name
	accessClaims["email"] = payload.Email
	accessClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	accessTokenString, err := accessToken.SignedString([]byte(ts.accessTokenSecret))
	if err != nil {
		return "", "", errors.Wrap(err, "failed to generate access token")
	}

	return accessTokenString, refreshTokenString, nil
}

func (ts *TokenService) SaveToken(refreshToken string, userID int) error {
	row := ts.db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM token WHERE user_id = %d", userID))

	var count int
	err := row.Scan(&count)
	if err != nil {
		return errors.Wrap(err, "failed to check token existence")
	}

	if count == 1 {
		_, err = ts.db.Exec(fmt.Sprintf("UPDATE token SET refresh_token = '%s' WHERE user_id = %d", refreshToken, userID))
		if err != nil {
			return errors.Wrap(err, "failed to update token")
		}
	} else {
		_, err = ts.db.Exec(fmt.Sprintf("INSERT INTO token (refresh_token, user_id) VALUES ('%s', %d)", refreshToken, userID))
		if err != nil {
			return errors.Wrap(err, "failed to insert token")
		}
	}

	return nil
}

func (ts *TokenService) ValidateAccessToken(accessToken string) (*UserRes, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(ts.accessTokenSecret), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse access token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid access token")
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return nil, errors.New("invalid user-authorized-authorized ID in access token")
	}

	name, ok := claims["name"].(string)
	if !ok {
		return nil, errors.New("invalid name in access token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("invalid email in access token")
	}

	return &UserRes{
		ID:    int(userID),
		Name:  name,
		Email: email,
	}, nil
}

func (ts *TokenService) ValidateRefreshToken(refreshToken string) (*UserRes, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(ts.refreshTokenSecret), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return nil, errors.New("invalid user-authorized-authorized ID in refresh token")
	}

	name, ok := claims["name"].(string)
	if !ok {
		return nil, errors.New("invalid name in refresh token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("invalid email in refresh token")
	}

	return &UserRes{
		ID:    int(userID),
		Name:  name,
		Email: email,
	}, nil
}
