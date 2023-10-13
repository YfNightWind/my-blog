package v1

import (
	"github.com/YfNightWind/my-blog/model"
	"github.com/YfNightWind/my-blog/utils/errormsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateArticleController 添加文章
func CreateArticleController(ctx *gin.Context) {
	var data model.Article

	_ = ctx.ShouldBindJSON(&data)
	code = model.CreateArticle(&data)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// DeleteArticleController 删除文章
func DeleteArticleController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code = model.DeleteArticle(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   nil,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// GetArticleListController 查询文章列表
func GetArticleListController(ctx *gin.Context) {
	title := ctx.Query("title")
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	data, code, total := model.GetArticleList(title, pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"total":  total,
		"msg":    errormsg.GetErrorMsg(errormsg.SUCCESS),
	})
}

// GetCategoryArticleListController 查询分类下的所有文章
func GetCategoryArticleListController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	data, code, total := model.GetCategoryArticleList(id, pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"total":  total,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// GetArticleInfoController 查询单个文章信息
func GetArticleInfoController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := model.GetArticleInfo(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// EditArticleController 编辑文章
func EditArticleController(ctx *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&data)

	code = model.EditArticle(id, &data)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetErrorMsg(code),
	})
}
