package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"react-go/backend/application/service"
	"react-go/backend/domain/model"
)

func SignUp(c *gin.Context) {
	params := &model.User{}
	if err := c.BindJSON(params); err != nil {
		log.Println(err)
		return
	}

	ctx := c.Request.Context()

	user, err := service.SignUpService(ctx, params)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, user)
}
