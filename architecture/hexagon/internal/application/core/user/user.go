package user

import (
	"coffee/internal/application/core/errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	Email          string
	HashedPassword string
}

type User struct{}

func New() IUser {
	return &User{}
}

func (user *User) NewUser(email, password string) (UserData, error) {
	if email == "" {
		return UserData{}, fmt.Errorf(errors.EmptyEmail)
	}

	if password == "" {
		return UserData{}, fmt.Errorf(errors.EmptyPassword)
	}

	if !validateEmail(email) {
		return UserData{}, fmt.Errorf(errors.EmailIsNotValid)
	}

	hashedPassword, err := hashPassword(password)

	if err != nil {
		return UserData{}, err
	}

	return UserData{
		Email:          email,
		HashedPassword: hashedPassword,
	}, nil
}

func (user *User) ComparePassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func validateEmail(email string) bool {
	// Test email is valid format
	if match, _ := regexp.MatchString(`^[\w\.\_]{2,}@\w{2,}\.\w{2,}$`, email); !match {
		return false
	}

	return true
}

func hashPassword(password string) (string, error) {
	dat, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(dat), nil
}
