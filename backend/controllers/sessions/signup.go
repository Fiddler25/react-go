package sessions

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"react-go/backend/models"
	"react-go/backend/services/users"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		log.Println(err)
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	res := users.CreateUser(user.Email, string(encryptedPassword))
	c.JSON(200, gin.H{"result": res})
}
