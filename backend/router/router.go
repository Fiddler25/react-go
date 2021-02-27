package router

import (
	"github.com/gin-gonic/gin"
	"react-go/backend/controllers"
)

func Init() {
	router := gin.Default()
	endpoints(router)

	router.Run(":3000")
}

func endpoints(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1api := api.Group("/v1")
		{
			v1api.POST("/signup", controllers.Signup)
		}
	}
}
