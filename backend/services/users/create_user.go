package users

import (
	"log"
	"react-go/backend/config/db"
)

func CreateUser(email string, password string) {
	db := db.DBConnector()
	defer db.Close()

	sql, err := db.Prepare("INSERT INTO users(Email, Password) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	sql.Exec(email, password)
}
