package service

import (
	"context"
	"log"
	"react-go/backend/config/db"
	"react-go/backend/domain/model"
	"react-go/backend/domain/repository"
)

func SignUpService(ctx context.Context, params *model.User) (user *model.User, err error) {
	user, err = repository.NewUser(params.Email, params.Password)
	if err != nil {
		return nil, err
	}

	id, err := createUser(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = id

	return user, nil
}

func createUser(ctx context.Context, user *model.User) (int, error) {
	db := db.DBConnector()
	defer db.Close()

	sql := "INSERT INTO users (Email, Password) VALUES (?, ?)"
	stmt, err := db.PrepareContext(ctx, sql)
	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	result, err := stmt.ExecContext(ctx, user.Email, user.Password)
	id, err := result.LastInsertId()

	return int(id), nil
}
