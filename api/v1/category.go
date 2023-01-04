package v1

import (
	"github.com/gin-gonic/gin"
	"my-blog/model"
	"my-blog/utils/errormsg"
	"net/http"
	"strconv"
)

// CreateCategoryController 添加分类
func CreateCategoryController(ctx *gin.Context) {
	var data model.Category
	_ = ctx.ShouldBindJSON(&data)
	code = model.IsCategoryExist(data.Name)
	if code == errormsg.SUCCESS {
		// 写进数据库
		model.CreateCategory(&data)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// DeleteCategoryController 删除分类
func DeleteCategoryController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code = model.DeleteCategory(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   nil,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// GetCategoryListController 查询分类列表
func GetCategoryListController(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	data, total := model.GetCategoryList(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status": errormsg.SUCCESS,
		"data":   data,
		"total":  total,
		"msg":    errormsg.GetErrorMsg(errormsg.SUCCESS),
	})
}

// EditCategoryController 编辑分类
func EditCategoryController(ctx *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&data)

	// 修改分类名时候，检查更新后的分类名是否存在
	code = model.IsCategoryExist(data.Name)

	if code == errormsg.SUCCESS {
		model.EditCategory(id, &data)
	}

	if code == errormsg.ERROR_CATEGORYNAME_USED {
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetErrorMsg(code),
	})
}
