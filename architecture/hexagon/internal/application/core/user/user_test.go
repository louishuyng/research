package user

import (
	"coffee/internal/application/core/errors"
	"testing"
)

func TestNewUserWhenEmailIsEmpty(t *testing.T) {
	user := New()

	_, err := user.NewUser("", "password")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if err.Error() != errors.EmptyEmail {
		t.Errorf("Expected error %s, got %s", errors.EmptyEmail, err.Error())
	}
}

func TestNewUserWhenPasswordIsEmpty(t *testing.T) {
	user := New()

	_, err := user.NewUser("test@gmail.com", "")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if err.Error() != errors.EmptyPassword {
		t.Errorf("Expected error %s, got %s", errors.EmptyPassword, err.Error())
	}
}

func TestNewUserWhenEmailIsNotValid(t *testing.T) {
	user := New()

	_, err := user.NewUser("testgmail.com", "password")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if err.Error() != errors.EmailIsNotValid {
		t.Errorf("Expected error, got %s", errors.EmailIsNotValid)
	}
}

func TestNewUserSuccess(t *testing.T) {
	user := New()

	userData, err := user.NewUser("test@gmail.com", "password")

	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}

	if userData.Email != "test@gmail.com" {
		t.Errorf("Expected email %s, got %s", "test@gmail.com", userData.Email)
	}

	if userData.HashedPassword == "" {
		t.Error("Expected hashed password, got empty")
	}
}

func TestComparePasswordWhenPasswordIsNotMatch(t *testing.T) {
	user := New()

	err := user.ComparePassword("password", "hashedPassword")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestComparePasswordSuccess(t *testing.T) {
	user := New()

	userData, _ := user.NewUser("test@gmail.com", "password")

	err := user.ComparePassword("password", userData.HashedPassword)

	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}
}
