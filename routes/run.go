package routes

import (
	c "studentApi/config"

	"github.com/gin-gonic/gin"
)
func RunService() error {
	env := c.GetWebServerEnv()
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	MainRouter(router)
	if err := router.Run(env.Port); err != nil {
		return err
	}
	
	return nil
}