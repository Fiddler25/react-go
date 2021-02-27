package users

import (
	"log"
	"react-go/backend/config/db"
)

func CreateUser(id int, email string, password string) {
	db := db.DBConnector()
	defer db.Close()

	sql, err := db.Prepare("INSERT INTO users(ID, Email, Password) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	sql.Exec(id, email, password)
}
