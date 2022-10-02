package main

import (
	v1 "myblog-gin/api/v1"
	"myblog-gin/model"
	"myblog-gin/routes"
	"myblog-gin/tool"
)

func main() {
	model.InitDb()
	tool.InitRedisStore()
	v1.InitSearchJson()
	routes.InitRouter()
}
