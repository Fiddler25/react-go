package users

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"react-go/backend/models"
)

func RequestHandler(c *gin.Context) (u models.User) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err := json.Unmarshal([]byte(string(value)), &user); err != nil {
		log.Fatal(err)
	}

	return user
}
