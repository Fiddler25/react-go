package auth

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"react-go/backend/config/db"
	"react-go/backend/models"
)

func Login(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err := json.Unmarshal([]byte(string(value)), &user); err != nil {
		log.Fatal(err)
	}
	paramsPassword := user.Password

	db := db.DBConnector()
	defer db.Close()

	if err = db.QueryRow("SELECT * FROM users where Email = ?", user.Email).Scan(&user.ID, &user.Email, &user.Password); err != nil {
		c.Status(http.StatusBadRequest)
	}
	encryptedPassword := user.Password

	result := false
	err = bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(paramsPassword))
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		session := sessions.Default(c)
		session.Set("loginUser", user.ID)
		session.Save()
		
		result = true
	}

	c.JSON(200, gin.H{"result": result})
}
