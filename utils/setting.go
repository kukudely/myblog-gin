package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	ReAddr     string
	RePort     string
	RePassWord string
	ReDb       int
)

func init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径；", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadRedis(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("147258369")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
func LoadRedis(file *ini.File) {
	ReAddr = file.Section("redis").Key("ReAddr").MustString("180.76.187.248")
	RePort = file.Section("redis").Key("RePort").MustString("6379")
	RePassWord = file.Section("redis").Key("RePassWord").MustString("")
	ReDb = file.Section("redis").Key("ReDb").MustInt(0)

}
