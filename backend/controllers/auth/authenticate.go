package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"react-go/backend/services/users"
)

func Signup(c *gin.Context) {
	user := users.RequestHandler(c)

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	users.CreateUser(user.Email, string(encryptedPassword))
}

func Login(c *gin.Context) {
	user := users.RequestHandler(c)
	result := users.LoggedIn(c, user)

	c.JSON(200, gin.H{"result": result})
}
