package user_authorized_service

import (
	db "extended_todo/routing"
	"fmt"
)

type UserFull struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Registration(email string, password string, name string) UserFull {

	//rows, err := db.DB.Query("SELECT * FROM users")
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()

	// An album slice to hold data from returned rows.
	//var users = make([]UserFull, 0)

	// Loop through rows, using Scan to assign column data to struct fields.
	//for rows.Next() {
	//	var user UserFull
	//	if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
	//		return users, err
	//	}
	//	users = append(users, user)
	//}
	//if err = rows.Err(); err != nil {
	//	return users, err
	//}
	//
	//var finalUsers []UserFull
	//for i := 0; i < len(users); i++ {
	//	finalUsers = append(finalUsers, UserFull{
	//		Id:       users[i].Id,
	//		Name:     users[i].Name,
	//		Email:    users[i].Email,
	//		Password: users[i].Password,
	//	})
	//
	//}
	//
	//return finalUsers, nil

	u := UserFull{}
	// Query for a value based on a single row.
	if err := db.DB.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *", name, email, password).Scan(&u.Id, &u.Email, &u.Name, &u.Password); err != nil {
	}

	fmt.Print("GOOD")

	return u

	//
	//var user
	//
	//err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	//if err != nil {
	//	return nil, nil, errors.Wrap(err, "failed to scan row")
	//}
	//
	//userDto := dto.NewUserDto(&user)
	//
	//tokens, err := us.tokenService.GenerateTokens(userDto)
	//if err != nil {
	//	return nil, nil, errors.Wrap(err, "failed to generate tokens")
	//}
	//
	//return userDto, tokens, nil
}

//func Login(email string, password string) (*dto.UserDto, *TokenPair, error) {
//	candidate, err := utils.CheckUserCandidate(email)
//	if err != nil {
//		return nil, nil, errors.Wrap(err, "failed to check user-unauthorized-authorized candidate")
//	}
//
//	if !candidate {
//		return nil, nil, exceptions.NewErrorAPI("Invalid credentials or password", exceptions.Unauthorized)
//	}
//
//	row := db.DB.QueryRow("SELECT * FROM public.users WHERE email = $1", email)
//
//	var user controller.UserRes
//	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
//	if err != nil {
//		return nil, nil, errors.Wrap(err, "failed to scan row")
//	}
//
//	passwordCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
//	if passwordCompare != nil {
//		return nil, nil, exceptions.NewErrorAPI("Invalid credentials or password", exceptions.Unauthorized)
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
