package routers

import (
	"github.com/gin-gonic/gin"
)

// InitRoute ..
func InitRoute() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	setTestRoutes(router)
	router.Run()
	return router
}
