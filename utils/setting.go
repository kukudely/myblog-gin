package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file,err := ini.Load("../config/config.ini")
	if err != nil{
		fmt.Println("配置文件读取错误，请检查文件路径；",err)
	}
}
func LoadServer(file *ini.File){
	AppMode = 
}