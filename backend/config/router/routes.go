package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"react-go/backend/controllers/sessions"
)

func Init() {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

	endpoints(router)

	router.Run(":3000")
}

func endpoints(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1api := api.Group("/v1")
		{
			v1api.POST("/signup", sessions.Signup)
		}
	}
}
