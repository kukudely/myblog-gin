package routes

import (
	v1 "myblog-gin/api/v1"
	"myblog-gin/utils"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("front", "web/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.HTMLRender = createMyRender()
	r.Static("/assets", "./web/dist/assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	/*
		前端展示页面接口
	*/
	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add", v1.AddUser)
		// router.GET("user/:id", v1.GetUserInfo)
		router.GET("users", v1.GetUsers)

		// 文章分类信息模块
		// router.GET("category", v1.GetCate)
		// router.GET("category/:id", v1.GetCateInfo)

		// 文章模块
		// router.GET("article", v1.GetArt)
		// router.GET("article/list/:id", v1.GetCateArt)
		// router.GET("article/info/:id", v1.GetArtInfo)

		// 登录控制模块
		// router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

		// 获取个人设置信息
		router.GET("profile/:id", v1.GetProfile)

		// 更新个人信息
		router.PUT("profile/:id", v1.UpdateProfile)

		// 评论模块
		// router.POST("addcomment", v1.AddComment)
		// router.GET("comment/info/:id", v1.GetComment)
		// router.GET("commentfront/:id", v1.GetCommentListFront)
		// router.GET("commentcount/:id", v1.GetCommentCount)
	}

	r.Run(utils.HttpPort)
}
