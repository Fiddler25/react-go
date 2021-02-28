package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"react-go/backend/config/db"
	"react-go/backend/services/users"
)

func Login(c *gin.Context) {
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
	user := users.RequestHandler(c)

	c.JSON(200, gin.H{"result": result})
}
