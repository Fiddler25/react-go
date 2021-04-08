package main

import (
	"react-go/backend/application"
	"react-go/backend/domain/service"
	"react-go/backend/infra/db"
	"react-go/backend/infra/router"
	"react-go/backend/interface/controller"
)

func main() {
	apiV1 := router.G.Group("/v1")

	ac := initializeAuthenticationController()
	ac.InitAuthenticationAPI(apiV1)

	if err := router.G.Run(":3000"); err != nil {
		panic(err.Error())
	}
}

func initializeAuthenticationController() controller.AuthenticationController {
	uRepo := db.NewUserRepository()
	uService := service.NewUserService(uRepo)
	di := application.NewAuthenticationServiceDIInput(uRepo, uService)
	aApp := application.NewAuthenticationService(di)

	return controller.NewAuthenticationController(aApp)
}
