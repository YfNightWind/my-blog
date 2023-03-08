package routers

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	v1 "my-blog/api/v1"
	"my-blog/middleware"
	"my-blog/utils"
	"net/http"
)

// 渲染多个HTML模板
func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("front", "static/front/index.html")
	p.AddFromFiles("admin", "static/admin/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)

	// 路由初始化
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Log())         // 使用自定义日志中间件
	r.Use(middleware.Cors())        // 跨域中间件
	r.HTMLRender = createMyRender() // 渲染HTML

	r.Static("front/static", "static/front/static")
	r.Static("admin/static", "static/admin/static")
	r.StaticFile("/favicon.ico", "/static/front/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin", nil)
	})

	// 公共部分
	public := r.Group("api/v1")
	{
		// 用户模块的路由接口
		public.GET("users", v1.GetUserListController)       // 获取用户列表
		public.GET("user/:id", v1.GetUserController)        // 查询单个用户
		public.POST("user/add", v1.AddUserController)       // 用户注册
		public.POST("login", v1.AdminLoginController)       // 后台登录
		public.POST("front/login", v1.FrontLoginController) // 前台登录

		// 分类模块的路由接口
		public.GET("category", v1.GetCategoryListController) // 获取分类列表
		public.GET("category/:id", v1.GetCategoryController) // 查询单个分类

		// 文章模块的路由接口
		public.GET("article", v1.GetArticleListController)                  // 获取文章列表
		public.GET("article/info/:id", v1.GetArticleInfoController)         // 获取单个文章信息
		public.GET("article/list/:id", v1.GetCategoryArticleListController) // 获取分类下的所有文章

		// 个人信息
		public.GET("profile/:id", v1.GetProfileController) // 获取个人信息
	}

	// 需要使用token中间件的
	authorized := r.Group("api/v1")
	authorized.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		authorized.PUT("user/:id", v1.EditUserController)                      // 编辑用户
		authorized.DELETE("user/:id", v1.DeleteUserController)                 // 删除用户
		authorized.PUT("user/changePassword/:id", v1.ChangePasswordController) // 修改密码

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

		// 个人信息
		authorized.PUT("profile/:id", v1.UpdateProfileController) // 更新个人信息
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println(err)
	}
}
