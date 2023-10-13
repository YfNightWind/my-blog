package main

import (
	"github.com/YfNightWind/my-blog/model"
	"github.com/YfNightWind/my-blog/routers"
)

func main() {
	// 引用数据库
	model.InitDb()
	// 引用路由
	routers.InitRouter()
}
