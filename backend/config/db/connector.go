package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"react-go/backend/models"
)

func Init() {
	db := DBConnector()
	defer db.Close()

	query, err := db.Query("SELECT * FROM users")
	defer query.Close()
	if err != nil {
		log.Fatal(err)
	}

	for query.Next() {
		var user models.User
		err := query.Scan(&user.ID, &user.UserID, &user.Password)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func DBConnector() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))
	log.Println("Connected to mysql.")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
