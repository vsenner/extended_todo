package user_authorized_service

import (
	db "extended_todo/routing"
)

type UserFull struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Registration(email string, password string, name string) UserFull {
	u := UserFull{}
	if err := db.DB.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *", name, email, password).Scan(&u.Id, &u.Name, &u.Email, &u.Password); err != nil {
	}

	return u
}

//
//func Refresh(refreshToken string) (*dto.UserDto, *TokenPair, error) {
//	validate, err := us.tokenService.ValidateRefreshToken(refreshToken)
//	if err != nil {
//		return nil, nil, exceptions.NewErrorAPI("Invalid refresh token", exceptions.Unauthorized)
//	}
//
//	row := db.DB.QueryRow("SELECT * FROM users WHERE id = $1", validate.ID)
//
//	var user controller.UserRes
//	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
//	if err != nil {
//		return nil, nil, errors.Wrap(err, "failed to scan row")
//	}
//
//	userDto := dto.NewUserDto(&user)
//
//	tokens, err := us.tokenService.GenerateTokens(userDto)
//	if err != nil {
//		return nil, nil, errors.Wrap(err, "failed to generate tokens")
//	}
//
//	err = us.tokenService.SaveToken(tokens.RefreshToken, user.ID)
//	if err != nil {
//		return nil, nil, errors.Wrap(err, "failed to save token")
//	}
//
//	return userDto, tokens, nil
//}
