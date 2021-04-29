package routers

import (
	"example.com/controllers/test"
	userControllers "example.com/controllers/user"

	"github.com/gin-gonic/gin"
)

func setTestRoutes(router *gin.Engine) {
	router.GET("/test", test.Test)
	router.POST("/signup", userControllers.SignupUser)
}
