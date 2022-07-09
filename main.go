package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	//localhost:8080
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello,Gin!",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
