package db

import (
	"context"
	"log"
	"react-go/backend/domain/model"
	"react-go/backend/domain/repository"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (repo *userRepository) InsertUser(ctx context.Context, user *model.User) (*model.User, error) {
	db := DBConnector()
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

	user.ID = int(id)

	return user, nil
}
