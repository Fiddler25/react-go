package users

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"react-go/backend/config/db"
	"react-go/backend/models"
)

func UserLogin(c *gin.Context) (res bool) {
	rowPassword, encryptedPassword := getRegisteredPassword(c)

	res = true
	if err := bcrypt.CompareHashAndPassword(
		[]byte(encryptedPassword),
		[]byte(rowPassword),
	); err != nil {
		res = false
		log.Println(err)
	}

	return res
}

func getRegisteredPassword(c *gin.Context) (rowPassword string, encryptedPassword string) {
	db := db.DBConnector()
	defer db.Close()

	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println(err)
	}

	var user models.User
	if err := json.Unmarshal([]byte(string(value)), &user); err != nil {
		log.Println(err)
	}

	rowPassword = user.Password
	if err := db.QueryRow(
		"SELECT Email, Password FROM users WHERE Email = ?", user.Email,
	).Scan(&user.Email, &user.Password); err != nil {
		c.Status(http.StatusBadRequest)
		log.Println(err)
	}
	encryptedPassword = user.Password

	return rowPassword, encryptedPassword
}
