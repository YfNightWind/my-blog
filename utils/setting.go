package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

// 初始化
func init() {
	file, err := ini.Load("config/config.ini")

	if err != nil {
		fmt.Println("配置文件读取错误，请检查内容: ", err)
	}

	// 加载config.ini里面的配置
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("你的地址")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("你的数据库用户名")
	DbPassword = file.Section("database").Key("DbPassword").MustString("你的数据库密码")
	DbName = file.Section("database").Key("DbName").MustString("my-blog")
}
