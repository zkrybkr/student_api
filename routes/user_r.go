package routes

import (
	c "studentApi/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {

	r.GET("/user", c.ListUsers)

	r.POST("/user/add", c.CreateUser)
	
}
