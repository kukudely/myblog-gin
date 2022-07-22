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
	r.POST("/login_ajax_check", v1.LoginFront)
	r.Run(utils.HttpPort)
}
