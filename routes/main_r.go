package routes

import (
	c "studentApi/config"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func MainRouter(router *gin.Engine) {
	env := c.GetWebServerEnv()
	config := cors.DefaultConfig()
	config.AllowOrigins = env.AllowOrigins
	config.AllowHeaders = env.AllowHeaders
	config.AllowMethods = env.AllowMethods
	config.AllowCredentials = true

	router.Use(cors.New(config))
	
}