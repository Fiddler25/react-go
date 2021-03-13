package sessions

import (
	"github.com/gin-gonic/gin"
	"react-go/backend/services/users"
)

func Login(c *gin.Context) {
	result := users.UserLogin(c)

	c.JSON(200, gin.H{"result": result})
}
