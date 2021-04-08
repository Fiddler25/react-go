package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var G *gin.Engine

func init() {
	g := gin.New()
	G = g

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4000"}
	G.Use(cors.New(config))
}
