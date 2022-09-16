package main

import (
	"myblog-gin/model"
	"myblog-gin/routes"
	"myblog-gin/tool"
)

func main() {
	model.InitDb()
	tool.InitRedisStore()
	routes.InitRouter()
}
