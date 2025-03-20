package auth

import (
	core "coffee/internal/application/core/user"
	apiPort "coffee/internal/ports/left/http"
	dbPort "coffee/internal/ports/right/database"
	"context"
	"fmt"
	"time"
)

type AuthENV struct {
	JWT_SECRET string
}

type AuthApplication struct {
	db       dbPort.UserDbPort
	userCore core.IUser
	env      AuthENV
}

func NewAuthApplication(db dbPort.UserDbPort, userCore core.IUser, env AuthENV) *AuthApplication {
	return &AuthApplication{
		db:       db,
		userCore: userCore,
		env:      env,
	}
}

func (authApp *AuthApplication) SignUp(
	ctx context.Context,
	email,
	password string,
) (apiPort.AuthData, error) {
	userData, err := authApp.userCore.NewUser(email, password)

	if err != nil {
		return apiPort.AuthData{}, err
	}

	savedUser, err := authApp.db.CreateUser(ctx, userData.Email, userData.HashedPassword)

	if err != nil {
		return apiPort.AuthData{}, err
	}

	token, err := authApp.generateToken(
		ctx,
		savedUser,
		time.Hour,
		time.Hour*24*7,
	)

	if err != nil {
		return apiPort.AuthData{}, fmt.Errorf("could not generate token")
	}

	return apiPort.AuthData{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		User: apiPort.AuthUser{
			Email: savedUser.Email,
		},
	}, nil
}

func (authApp *AuthApplication) SignIn(
	ctx context.Context,
	email,
	password string,
) (apiPort.AuthData, error) {
	user, err := authApp.db.GetUserByEmail(ctx, email)

	if err != nil {
		return apiPort.AuthData{}, err
	}

	err = authApp.userCore.ComparePassword(password, user.HashedPassword)

	if err != nil {
		return apiPort.AuthData{}, fmt.Errorf("invalid email or password")
	}

	token, err := authApp.generateToken(
		ctx,
		user,
		time.Hour,
		time.Hour*24*7,
	)

	if err != nil {
		return apiPort.AuthData{}, fmt.Errorf("could not generate token")
	}

	return apiPort.AuthData{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		User: apiPort.AuthUser{
			Email: user.Email,
		},
	}, nil
}
