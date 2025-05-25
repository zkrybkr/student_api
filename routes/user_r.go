package routes

import (
	c "studentApi/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/user", c.ListUsers)
	v1.POST("/user", c.CreateUser)
}
