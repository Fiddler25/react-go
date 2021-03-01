package sessions

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"react-go/backend/models"
	"react-go/backend/services/users"
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

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	users.CreateUser(user.Email, string(encryptedPassword))
}
