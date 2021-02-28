package auth

import (
	"github.com/gin-gonic/gin"
	"react-go/backend/services/users"
)

func Login(c *gin.Context) {
	user := users.RequestHandler(c)
	result := users.LoggedIn(c, user)

	c.JSON(200, gin.H{"result": result})
}
