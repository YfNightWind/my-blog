package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-blog/utils"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)

	// 路由初始化
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println(err)
	}
}
