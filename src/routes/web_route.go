package route

import (
	"NewThread/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)
	group_router := router.Group("/groups")
	{
		group_router.GET("list", controller.GroupInfo)
		group_router.GET("teacher", controller.GroupTeacherInfo)
		group_router.GET("student", controller.GroupStudentInfo)
	}
	return router
}
