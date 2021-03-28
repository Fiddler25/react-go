package controller

import (
	"react-go/backend/domain/model"
)

type UserDTO struct {
	ID       int
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func translateFromUserToUserDTO(user *model.User) *UserDTO {
	return &UserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
}
