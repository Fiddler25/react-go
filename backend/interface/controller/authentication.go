package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"react-go/backend/application"
	"react-go/backend/domain/model"
)

type AuthenticationController interface {
	InitAuthenticationAPI(g *gin.RouterGroup)
	SignUp(g *gin.Context)
}

type authenticationController struct {
	aApp application.AuthenticationService
}

func NewAuthenticationController(uAPP application.AuthenticationService) AuthenticationController {
	return &authenticationController{
		aApp: uAPP,
	}
}

func (c *authenticationController) InitAuthenticationAPI(g *gin.RouterGroup) {
	g.POST("/signup", c.SignUp)
}

func (c *authenticationController) SignUp(g *gin.Context) {
	params := &model.User{}
	if err := g.BindJSON(params); err != nil {
		log.Println(err)
		return
	}

	ctx := g.Request.Context()

	user, err := c.aApp.SignUp(ctx, params)
	if err != nil {
		log.Println(err)
	}

	uDTO := translateFromUserToUserDTO(user)
	g.JSON(http.StatusOK, uDTO)
}
