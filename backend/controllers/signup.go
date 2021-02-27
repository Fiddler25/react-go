package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"react-go/backend/config/db"
	"react-go/backend/models"
)

func Signup(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err := json.Unmarshal([]byte(string(value)), &user); err != nil {
		log.Fatal(err)
	}

	createUser(user.ID, user.UserID, user.Password)
}

func createUser(id int, userID string, password string) {
	db := db.DBConnector()
	defer db.Close()

	sql, err := db.Prepare("INSERT INTO users(ID, UserID, Password) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	sql.Exec(id, userID, password)
}
