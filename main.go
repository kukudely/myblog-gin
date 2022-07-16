package main

import (
	"gintest/model"
	"gintest/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
