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
		group_router.GET("teachers", controller.GroupTeacherInfo)
		group_router.GET("students", controller.GroupStudentInfo)
	}
	honour_router := router.Group("/honours")
	{
		honour_router.GET("list", controller.HonoursList)
		honour_router.GET("students", controller.HonoursStudents)
		honour_router.GET("projects", controller.HonoursProjects)
	}
	return router
}
