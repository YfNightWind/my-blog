package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "my-blog/api/v1"
	"my-blog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)

	// 路由初始化
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	{
		// 用户模块的路由接口
		routerV1.POST("user/add", v1.AddUserController)      // 添加用户
		routerV1.GET("users", v1.GetUserListController)      // 获取用户列表
		routerV1.PUT("user/:id", v1.EditUserController)      // 编辑用户
		routerV1.DELETE("user/:id", v1.DeleteUserController) // 删除用户

		// 分类模块的路由接口
		routerV1.POST("category/add", v1.CreateCategoryController)   // 添加分类
		routerV1.GET("category", v1.GetCategoryListController)       // 获取分类列表
		routerV1.PUT("category/:id", v1.EditCategoryController)      // 编辑分类
		routerV1.DELETE("category/:id", v1.DeleteCategoryController) // 删除分类

		// 文章模块的路由接口
		routerV1.POST("article/add", v1.CreateArticleController)              // 添加文章
		routerV1.GET("article", v1.GetArticleListController)                  // 获取文章列表
		routerV1.GET("article/list/:id", v1.GetCategoryArticleListController) // 获取分类下的所有文章
		routerV1.GET("article/info/:id", v1.GetArticleInfoController)         // 获取单个文章信息
		routerV1.PUT("article/:id", v1.EditArticleController)                 // 编辑文章
		routerV1.DELETE("article/:id", v1.DeleteArticleController)            // 删除文章
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println(err)
	}
}
