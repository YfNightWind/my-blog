package v1

import (
	"github.com/YfNightWind/my-blog/model"
	"github.com/YfNightWind/my-blog/utils/errormsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddCommentController 新增评论
func AddCommentController(ctx *gin.Context) {
	var commentJson model.Comment

	_ = ctx.ShouldBindJSON(&commentJson)
	code = model.AddComment(&commentJson)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   commentJson,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// GetCommentController 查询单个评论
func GetCommentController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, code := model.GetComment(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// GetCommentListController 获取评论列表(后台)
func GetCommentListController(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.Param("pagenum"))
	pageSize, _ := strconv.Atoi(ctx.Param("pagesize"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code := model.GetCommentList(pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"total":  total,
		"data":   data,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// GetCommentNumberController 获取评论数量
func GetCommentNumberController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	totalNumber := model.GetCommentNumber(id)

	ctx.JSON(http.StatusOK, gin.H{
		"total": totalNumber,
	})
}

// DeleteCommentController 删除评论
func DeleteCommentController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	code = model.DeleteComment(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// ArticleGetCommentListController 展示文章底下的评论
func ArticleGetCommentListController(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.Param("pagenum"))
	pageSize, _ := strconv.Atoi(ctx.Param("pagesize"))
	id, _ := strconv.Atoi(ctx.Param("id"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code := model.ArticleGetCommentList(id, pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"total":  total,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// PassTheCommentController 通过评论
func PassTheCommentController(ctx *gin.Context) {
	var commentJson model.Comment
	_ = ctx.ShouldBindJSON(&commentJson)

	id, _ := strconv.Atoi(ctx.Param("id"))

	code = model.PassTheComment(id, &commentJson)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// RemoveTheCommentController 撤下该评论
func RemoveTheCommentController(ctx *gin.Context) {
	var commentJson model.Comment
	_ = ctx.ShouldBindJSON(&commentJson)

	id, _ := strconv.Atoi(ctx.Param("id"))

	code = model.RemoveTheComment(id, &commentJson)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetErrorMsg(code),
	})
}
