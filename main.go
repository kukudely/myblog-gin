package main

import (
	"myblog-gin/model"
	"myblog-gin/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
