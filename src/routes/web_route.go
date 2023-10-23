package route

import (
	"NewThread/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)
	user_router := router.Group("/user")
	{
		user_router.GET("", controller.UserInfo)
	}
	return router
}
