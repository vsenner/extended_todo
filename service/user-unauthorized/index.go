package user_authorized_service
import (
	utils "extended_todo/ utlis"
	//token_service "extended_todo/service/token"
	db "extended_todo/routing"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func  Registration(email string, password string, name string) (*dto.UserDto, *TokenPair, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate hashed password")
	}

	row := db.DB.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *", name, email, string(hashedPassword))

	var user
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to scan row")
	}

	userDto := dto.NewUserDto(&user)

	tokens, err := us.tokenService.GenerateTokens(userDto)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate tokens")
	}

	return userDto, tokens, nil
}

func  Login(email string, password string) (*dto.UserDto, *TokenPair, error) {
	candidate, err := utils.CheckUserCandidate(email)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to check user-unauthorized-authorized candidate")
	}

	if !candidate {
		return nil, nil, exceptions.NewErrorAPI("Invalid credentials or password", exceptions.Unauthorized)
	}

	row := db.DB.QueryRow("SELECT * FROM public.users WHERE email = $1", email)

	var user controller.UserRes
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to scan row")
	}

	passwordCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if passwordCompare != nil {
		return nil, nil, exceptions.NewErrorAPI("Invalid credentials or password", exceptions.Unauthorized)
	}

	userDto := dto.NewUserDto(&user)

	tokens, err := us.tokenService.GenerateTokens(userDto)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate tokens")
	}

	err = us.tokenService.SaveToken(tokens.RefreshToken, user.ID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to save token")
	}

	return userDto, tokens, nil
}

func  Refresh(refreshToken string) (*dto.UserDto, *TokenPair, error) {
	validate, err := us.tokenService.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, nil, exceptions.NewErrorAPI("Invalid refresh token", exceptions.Unauthorized)
	}

	row := db.DB.QueryRow("SELECT * FROM users WHERE id = $1", validate.ID)

	var user controller.UserRes
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to scan row")
	}

	userDto := dto.NewUserDto(&user)

	tokens, err := us.tokenService.GenerateTokens(userDto)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate tokens")
	}

	err = us.tokenService.SaveToken(tokens.RefreshToken, user.ID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to save token")
	}

	return userDto, tokens, nil
}
