package routes

import (
	v1 "myblog-gin/api/v1"
	"myblog-gin/middleware"
	"myblog-gin/utils"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("front", "web/dist/index.html")
	p.AddFromFiles("admin", "web/admin/index.html")
	p.AddFromFiles("searchjson", "Search.json")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.Use(middleware.Cors())
	// r.Use(middleware.Log())
	r.HTMLRender = createMyRender()
	r.Static("/assets", "./web/dist/assets")
	r.Static(" ", "./web/admin/assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v1.GetUsers)
		// auth.PUT("user/:id", v1.EditUser)
		// auth.DELETE("user/:id", v1.DeleteUser)
		//修改密码
		auth.PUT("admin/changepw/:id", v1.ChangeUserPassword)
		// 分类模块的路由接口
		// auth.GET("admin/category", v1.GetCate)
		auth.POST("category/add", v1.AddCategory)
		// auth.PUT("category/:id", v1.EditCate)
		// auth.DELETE("category/:id", v1.DeleteCate)
		// 文章模块的路由接口
		// auth.GET("admin/article/info/:id", v1.GetArtInfo)
		auth.GET("admin/article", v1.GetArt)
		auth.POST("article/add", v1.AddArticle)
		// auth.PUT("article/:id", v1.EditArt)
		// auth.DELETE("article/:id", v1.DeleteArt)
		// 上传文件
		// auth.POST("upload", v1.UpLoad)
		// 更新个人设置
		auth.GET("admin/profile/:id", v1.GetProfile)
		auth.PUT("profile/:id", v1.UpdateProfile)
		// 评论模块
		// auth.GET("comment/list", v1.GetCommentList)
		// auth.DELETE("delcomment/:id", v1.DeleteComment)
		// auth.PUT("checkcomment/:id", v1.CheckComment)
		// auth.PUT("uncheckcomment/:id", v1.UncheckComment)
	}

	/*
		前端展示页面接口
	*/
	router := r.Group("api/v1")
	{
		router.GET("/search.json", func(c *gin.Context) {
			c.HTML(200, "searchjson", nil)

		})
		// 用户信息模块
		router.POST("user/add", v1.AddUser)
		// router.GET("user/:id", v1.GetUserInfo)
		// router.GET("users", v1.GetUsers)

		// 文章分类信息模块
		router.GET("category", v1.GetCate)
		// router.GET("category/:id", v1.GetCateInfo)

		// 文章模块
		router.GET("article", v1.GetArt)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetOneArtInfo)
		router.GET("recentart", v1.GetRecArt)

		// 登录控制模块
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

		// 获取个人设置信息
		router.GET("profile/:id", v1.GetProfile)

		// 评论模块
		router.POST("addcomment", v1.AddComment)
		// router.GET("comment/info/:id", v1.GetComment)
		router.GET("commentfront/:id", v1.GetCommentListFront)
		// router.GET("commentcount/:id", v1.GetCommentCount)
	}

	r.Run(utils.HttpPort)
}
