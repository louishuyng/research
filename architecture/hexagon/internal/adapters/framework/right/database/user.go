package database

import (
	"context"
	"time"

	"coffee/internal/adapters/framework/right/database/repository"
	rightPort "coffee/internal/ports/right/database"

	"github.com/google/uuid"
)

type UserDbAdapter struct {
	q *repository.Queries
}

func NewUserDbAdapter(q *repository.Queries) *UserDbAdapter {
	return &UserDbAdapter{
		q: q,
	}
}

func (a *UserDbAdapter) CreateUser(ctx context.Context, email, hashedPassword string) (rightPort.User, error) {
	user, err := a.q.CreateUser(ctx, repository.CreateUserParams{
		Email:          email,
		HashedPassword: hashedPassword,
	})

	if err != nil {
		return rightPort.User{}, err
	}

	return rightPort.User{
		ID:             user.ID,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
	}, nil
}

func (a *UserDbAdapter) GetUserByEmail(ctx context.Context, email string) (rightPort.User, error) {
	user, err := a.q.GetUserByEmail(ctx, email)

	if err != nil {
		return rightPort.User{}, err
	}

	return rightPort.User{
		ID:             user.ID,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
	}, nil
}

func (a *UserDbAdapter) CreateRefreshToken(ctx context.Context, userId uuid.UUID, token string, expiresAt time.Time) error {
	_, err := a.q.CreateRefreshToken(ctx, repository.CreateRefreshTokenParams{
		UserID: userId,
		Token:  token,
	})

	if err != nil {
		return err
	}

	return nil
}
