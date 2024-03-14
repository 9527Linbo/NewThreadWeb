package route

import (
	"NewThread/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)

	// 路由组（团队）
	group_router := router.Group("/groups")
	{
		group_router.GET("list", controller.GroupInfo)            //团队信息
		group_router.GET("teachers", controller.GroupTeacherInfo) //老师信息
		group_router.GET("students", controller.GroupStudentInfo) //管理层学生信息
		group_router.GET("yearlist", controller.Yearlist)         //届
	}

	// 路由组（荣誉）
	honour_router := router.Group("/honours")
	{
		honour_router.GET("list", controller.HonoursList)
		honour_router.GET("students", controller.HonoursStudents)
		honour_router.GET("projects", controller.HonoursProjects)
		honour_router.GET("milestone", controller.HonoursMilestone)
		honour_router.GET("milestones", controller.HonoursMilestones)
	}

	// 路由组（帖子）
	group_post_router := router.Group("/post")
	{
		//group_post_router.GET("/list", controller.PostInfo)
		group_post_router.GET("/sharelist", controller.PageShareInfo)
		group_post_router.GET("/newslist", controller.PageNewsInfo)
		group_post_router.GET("/activitylist", controller.PageActivityInfo)
		group_post_router.GET("/readshare", controller.ReadShare)
		group_post_router.GET("/readnews", controller.ReadNews)
		group_post_router.GET("/readactivity", controller.ReadActivity)

		//帖子评论
		group_post_router.GET("/comment", controller.CommentInfo_topThree) //获取评论及前三条次级评论
	}

	// 路由组（文件）
	group_file_router := router.Group("/file")
	{
		group_file_router.POST("/upload", controller.UploadFile)    //上传文件
		group_file_router.GET("/download", controller.DownloadFile) //下载文件
		group_file_router.GET("/list", controller.FileList)         //获取文件列表
	}
	return router
}
