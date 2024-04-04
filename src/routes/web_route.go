package route

import (
	"NewThread/src/controller"
	"NewThread/src/middle"

	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)

	router.POST("/login", controller.Login)
	router.POST("/register", controller.Register)
	router.POST("/register/teacher", controller.RegisterTeacher)

	// 路由组（团队）
	group_router := router.Group("/groups")
	group_router.Use(middle.JWTAuth) //登录校验
	{
		group_router.GET("list", controller.GroupInfo)            //团队信息
		group_router.GET("teachers", controller.GroupTeacherInfo) //老师信息
		group_router.GET("students", controller.GroupStudentInfo) //管理层学生信息
		group_router.GET("yearlist", controller.Yearlist)         //届
	}

	// 路由组（荣誉）
	honour_router := router.Group("/honours")
	honour_router.Use(middle.JWTAuth)
	{
		honour_router.GET("list", controller.HonoursList)
		honour_router.GET("students", controller.HonoursStudents)
		honour_router.GET("projects", controller.HonoursProjects)
		honour_router.GET("milestone", controller.HonoursMilestone)
		honour_router.GET("milestones", controller.HonoursMilestones)
	}

	// 路由组（帖子）
	group_post_router := router.Group("/post")
	group_post_router.Use(middle.JWTAuth)
	{
		//group_post_router.GET("/list", controller.PostInfo)
		group_post_router.GET("/sharelist", controller.PageShareInfo)
		group_post_router.GET("/newslist", controller.PageNewsInfo)
		group_post_router.GET("/activitylist", controller.PageActivityInfo)
		group_post_router.GET("/readshare", controller.ReadShare)
		group_post_router.GET("/readnews", controller.ReadNews)
		group_post_router.GET("/readactivity", controller.ReadActivity)

		//帖子评论
		group_post_router.GET("/comment", controller.CommentInfo_topThree) //获取顶级评论及前三条次级评论
		group_post_router.GET("/allcomment", controller.CommentInfo_All)   //获取顶级评论及所有次级评论
		group_post_router.POST("/comment", controller.Comment_Upload)      //上传评论
	}

	// 路由组（文件）
	group_file_router := router.Group("/file")
	group_file_router.Use(middle.JWTAuth)
	{
		group_file_router.POST("/upload", controller.UploadFile)    //上传文件
		group_file_router.GET("/download", controller.DownloadFile) //下载文件
		group_file_router.GET("/list", controller.FileList)         //获取文件列表
	}

	//用户
	group_user_router := router.Group("/user")
	group_user_router.Use(middle.JWTAuth)
	{
		group_user_router.POST("/icon/upload", controller.UploadIcon) //上传头像
		group_user_router.GET("/icon", controller.UsersIcon)          //获取头像
	}

	//增添信息
	group_add_router := router.Group("/add")
	group_add_router.Use(middle.JWTAuth)
	{
		group_add_router.POST("/teacher", controller.AddTeacher)     //增添老师信息
		group_add_router.POST("/student", controller.AddStudent)     //增添学生信息
		group_add_router.POST("/graduate", controller.AddGraduate)   //增添毕业生信息
		group_add_router.POST("/milestone", controller.AddMilestone) //增添里程碑

	}

	return router
}
