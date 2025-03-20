package ports

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Email          string
	HashedPassword string
}

type UserDbPort interface {
	CreateUser(ctx context.Context, email, hashedPassword string) (User, error)
	CreateRefreshToken(ctx context.Context, userId uuid.UUID, token string, expiresAt time.Time) error
	GetUserByEmail(ctx context.Context, email string) (User, error)
}
