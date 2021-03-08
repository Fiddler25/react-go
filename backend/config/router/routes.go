package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"react-go/backend/controllers/sessions"
)

func Init() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4000"}
	router.Use(cors.New(config))

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
