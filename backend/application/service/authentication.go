package service

import (
	"context"
	"react-go/backend/domain/model"
	"react-go/backend/domain/repository"
)

func SignUpService(ctx context.Context, params *model.User) (user *model.User, err error) {
	user, err = repository.NewUser(params.Email, params.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
