package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/changecoder/webservice/handler"
)

// InitRouter Initialize Router
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	user := r.Group("v1/api/proxy/user")
	{
		user.POST("/login", handler.UserLogin)
		user.GET("/status", handler.UserStatus)
		user.GET("/info", handler.UserInfo)
	}

	return r
}
