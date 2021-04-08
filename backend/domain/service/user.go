package service

import (
	"golang.org/x/crypto/bcrypt"
	"react-go/backend/domain/model"
	"react-go/backend/domain/repository"
)

type UserService interface {
	NewUser(email, password string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) NewUser(email, password string) (*model.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Email:    email,
		Password: string(hashed),
	}, nil
}
