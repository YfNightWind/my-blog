package routers

import (
	"fmt"
	"net/http"

	"github.com/YfNightWind/my-blog/controller"
	"github.com/YfNightWind/my-blog/middleware"
	"github.com/YfNightWind/my-blog/utils"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// 渲染多个HTML模板 todo 将前端分离出来部署
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
	public := r.Group("api")
	{
		// 用户模块的路由接口
		public.GET("users", controller.GetUserListController)       // 获取用户列表
		public.GET("user/:id", controller.GetUserController)        // 查询单个用户
		public.POST("user/add", controller.AddUserController)       // 用户注册
		public.POST("login", controller.AdminLoginController)       // 后台登录
		public.POST("front/login", controller.FrontLoginController) // 前台登录

		// 分类模块的路由接口
		public.GET("category", controller.GetCategoryListController) // 获取分类列表
		public.GET("category/:id", controller.GetCategoryController) // 查询单个分类

		// 文章模块的路由接口
		public.GET("article", controller.GetArticleListController)                  // 获取文章列表
		public.GET("article/info/:id", controller.GetArticleInfoController)         // 获取单个文章信息
		public.GET("article/list/:id", controller.GetCategoryArticleListController) // 获取分类下的所有文章

		// 个人信息
		public.GET("profile/:id", controller.GetProfileController) // 获取个人信息

		// ChatGPT
		public.POST("chat", controller.Chat)

		// 评论模块路由接口
		public.GET("comment/info/:id", controller.GetCommentController)               // 获取文章评论
		public.GET("article/comment/:id", controller.ArticleGetCommentListController) // 获取文章下的评论
		public.GET("comment/number/:id", controller.GetCommentNumberController)       // 获取文章评论数量
	}

	// 需要使用token中间件的
	authorized := r.Group("api")
	authorized.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		authorized.PUT("user/:id", controller.EditUserController)                       // 编辑用户
		authorized.DELETE("user/:id", controller.DeleteUserController)                  // 删除用户
		authorized.PUT("user/change/password/:id", controller.ChangePasswordController) // 修改密码

		// 分类模块的路由接口
		authorized.POST("category/add", controller.CreateCategoryController)   // 添加分类
		authorized.PUT("category/:id", controller.EditCategoryController)      // 编辑分类
		authorized.DELETE("category/:id", controller.DeleteCategoryController) // 删除分类

		// 文章模块的路由接口
		authorized.POST("article/add", controller.CreateArticleController)   // 添加文章
		authorized.PUT("article/:id", controller.EditArticleController)      // 编辑文章
		authorized.DELETE("article/:id", controller.DeleteArticleController) // 删除文章

		// 上传
		authorized.POST("upload", controller.UploadController) // 上传文件

		// 个人信息
		authorized.PUT("profile/:id", controller.UpdateProfileController) // 更新个人信息

		// 评论模块路由接口
		authorized.POST("add/comment", controller.AddCommentController)             // 新增评论
		authorized.GET("comment/list", controller.GetCommentListController)         // 后台获取评论列表
		authorized.DELETE("delete/comment/:id", controller.DeleteCommentController) // 后台删除评论
		authorized.PUT("pass/comment/:id", controller.PassTheCommentController)     // 后台审核通过评论
		authorized.PUT("remove/comment/:id", controller.RemoveTheCommentController) // 后台撤下评论
	}

	// 404处理
	r.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "找不到该页面，请检查请求地址。")
	})

	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println(err)
	}
}
