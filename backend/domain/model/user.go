package model

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"log"
	"react-go/backend/infra/db"
)

type User struct {
	ID       int
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func NewUser(email, password string) (*User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:    email,
		Password: string(hashed),
	}, nil
}

func CreateUser(ctx context.Context, user *User) (int, error) {
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
