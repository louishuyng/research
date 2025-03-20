package ports

import "context"

type AuthUser struct {
	Email string
}

type AuthData struct {
	AccessToken  string
	RefreshToken string
	User         AuthUser
}

type AuthAPIPort interface {
	SignUp(ctx context.Context, email, password string) (AuthData, error)
	SignIn(ctx context.Context, email, password string) (AuthData, error)
}
