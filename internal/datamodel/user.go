package datamodel

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/think-free/ABCFitness-challenge/internal/errors"
)

type BaseUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

type CreateUserRequest struct {
	BaseUser
}

type User struct {
	ID string `json:"id"`
	BaseUser
}

func NewUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	id := uuid.New().String()

	u := &User{
		ID:       id,
		BaseUser: req.BaseUser,
	}

	if !u.isValid() {
		return nil, errors.ErrorValidationError()
	}

	return u, nil
}

func (u *User) isValid() bool {
	// TODO: add true validation
	return strings.Contains(u.Email, "@") &&
		strings.Contains(u.Email, ".") &&
		strings.Contains(u.Phone, "+")
}
