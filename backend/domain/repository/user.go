package repository

import (
	"context"
	"react-go/backend/domain/model"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *model.User) (*model.User, error)
}
