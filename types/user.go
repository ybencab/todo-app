package types

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(username, email, password string) (*User, error) {
	if len(email) == 0 || len(password) == 0 || len(username) == 0 {
		return nil, errors.New("username, email, and password must not be empty")
	}

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        uuid.New(),
		Username:  username,
		Email:     email,
		Password:  string(hashed_password),
		CreatedAt: time.Now(),
	}, nil
}
