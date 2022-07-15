package routes

import (
	"gintest/utils"
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("front", "web/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.HTMLRender = createMyRender()
	r.Static("/static", "./web")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	router := r.Group("api/v2")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello,Gin!",
			})
		})
	}
	r.Run(utils.HttpPort)
}
