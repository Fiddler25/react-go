package users

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"react-go/backend/config/db"
	"react-go/backend/models"
)

func LoggedIn(c *gin.Context, user models.User) (result bool) {
	db := db.DBConnector()
	defer db.Close()

	paramsPassword := user.Password
	if err := db.QueryRow("SELECT * FROM users where Email = ?", user.Email).Scan(&user.ID, &user.Email, &user.Password); err != nil {
		c.Status(http.StatusBadRequest)
	}
	encryptedPassword := user.Password

	result = false
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(paramsPassword))
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		session := sessions.Default(c)
		session.Set("loginUser", user.ID)
		session.Save()
		c.Status(http.StatusOK)
		result = true
	}

	return result
}
