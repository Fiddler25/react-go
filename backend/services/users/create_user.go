package users

import (
	"log"
	"react-go/backend/config/db"
)

func CreateUser(email string, password string) (result bool) {
	db := db.DBConnector()
	defer db.Close()

	result = true
	sql, err := db.Prepare("INSERT INTO users(Email, Password) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
		result = false
		return
	}
	sql.Exec(email, password)

	return result
}
