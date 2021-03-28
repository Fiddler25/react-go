package repository

import (
	"golang.org/x/crypto/bcrypt"
	"react-go/backend/domain/model"
)

func NewUser(email, password string) (*model.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Email:    email,
		Password: string(hashed),
	}, nil
}
