package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiNiuServer string

	ApiKey string

	Salt []byte
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
	LoadQiNiu(file)
	LoadChatGPT(file)
	//LoadSaltCode(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKet").MustString("blog-key")
}

func LoadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("")
	DbPassword = file.Section("database").Key("DbPassword").MustString("")
	DbName = file.Section("database").Key("DbName").MustString("my-blog")
}

func LoadQiNiu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiNiuServer = file.Section("qiniu").Key("QiNiuServer").String()
}

func LoadChatGPT(file *ini.File) {
	ApiKey = file.Section("ChatGPT").Key("ApiKey").String()
}

//func LoadSaltCode(file *ini.File) {
//	for i := 1; i <= 8; i++ {
//		file.Section("Salt").Key("saltCode" + string(rune(i)))
//	}
//	data, err := file.WriteTo()
//	if err != nil {
//		return
//	}
//}
