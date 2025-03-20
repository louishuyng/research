package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	dbPort "coffee/internal/ports/right/database"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

type Token struct {
	AccessToken  string
	RefreshToken string
}

const (
	// TokenTypeAccess -
	TokenTypeAccess TokenType = "coffee-access"
)

func GenerateToken(
	savedUser dbPort.User,
	tokenSecret string,
	expiresIn time.Duration,
) (string, error) {
	signingKey := []byte(tokenSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    string(TokenTypeAccess),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
		Subject:   savedUser.ID.String(),
	})

	return token.SignedString(signingKey)
}

// MakeRefreshToken makes a random 256 bit token
// encoded in hex
func MakeRefreshToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}

func (authApp *AuthApplication) generateToken(
	ctx context.Context,
	user dbPort.User,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
) (Token, error) {
	token, err := GenerateToken(user, authApp.env.JWT_SECRET, accessTokenDuration)

	if err != nil {
		return Token{}, err
	}

	refreshToken, err := MakeRefreshToken()

	if err != nil {
		return Token{}, err
	}

	err = authApp.db.CreateRefreshToken(
		ctx,
		user.ID,
		refreshToken,
		time.Now().Add(refreshTokenDuration),
	)

	if err != nil {
		return Token{}, err
	}

	return Token{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}, nil
}
