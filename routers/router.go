package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "my-blog/api/v1"
	"my-blog/middleware"
	"my-blog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)

	// 路由初始化
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Log())  // 使用自定义日志中间件
	r.Use(middleware.Cors()) // 跨域中间件

	// 公共部分
	public := r.Group("api/v1")
	{
		public.GET("users", v1.GetUserListController)                       // 获取用户列表
		public.GET("user/:id", v1.GetUserController)                        // 查询单个用户
		public.GET("category", v1.GetCategoryListController)                // 获取分类列表
		public.GET("article", v1.GetArticleListController)                  // 获取文章列表
		public.GET("article/list/:id", v1.GetCategoryArticleListController) // 获取分类下的所有文章
		public.GET("article/info/:id", v1.GetArticleInfoController)         // 获取单个文章信息
		public.POST("user/add", v1.AddUserController)                       // 用户注册
		public.POST("login", v1.LoginController)                            // 用户登录

	}

	// 需要使用token中间件的
	authorized := r.Group("api/v1")
	authorized.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		authorized.PUT("user/:id", v1.EditUserController)      // 编辑用户
		authorized.DELETE("user/:id", v1.DeleteUserController) // 删除用户

		// 分类模块的路由接口
		authorized.POST("category/add", v1.CreateCategoryController)   // 添加分类
		authorized.PUT("category/:id", v1.EditCategoryController)      // 编辑分类
		authorized.DELETE("category/:id", v1.DeleteCategoryController) // 删除分类

		// 文章模块的路由接口
		authorized.POST("article/add", v1.CreateArticleController)   // 添加文章
		authorized.PUT("article/:id", v1.EditArticleController)      // 编辑文章
		authorized.DELETE("article/:id", v1.DeleteArticleController) // 删除文章

		// 上传
		authorized.POST("upload", v1.UploadController) // 上传文件
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println(err)
	}
}
