package entity

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"ms-control-point/pkg/entity"
	"strings"
)

var (
	ErrEmailIsRequired    = errors.New("email is required")
	ErrNameIsRequired     = errors.New("name is required")
	ErrInvalidEmail       = errors.New("invalid email")
	ErrPasswordIsRequired = errors.New("password is required")
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	if name == "" {
		return nil, ErrNameIsRequired
	}
	if email == "" {
		return nil, ErrEmailIsRequired
	}
	if !strings.ContainsAny(email, "@") {
		return nil, ErrInvalidEmail
	}
	if password == "" {
		return nil, ErrPasswordIsRequired
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
