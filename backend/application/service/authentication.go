package service

import (
	"context"
	"react-go/backend/domain/model"
)

func SignUpService(ctx context.Context, params *model.User) (user *model.User, err error) {
	user, err = model.NewUser(params.Email, params.Password)
	if err != nil {
		return nil, err
	}

	id, err := model.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = id

	return user, nil
}
