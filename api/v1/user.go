package v1

import (
	"github.com/gin-gonic/gin"
	"my-blog/model"
	"my-blog/utils/errormsg"
	"net/http"
	"strconv"
)

var code int

// IsUserExistController TODO 查询用户是否存在
func IsUserExistController(ctx *gin.Context) {

}

// AddUserController 添加用户
func AddUserController(ctx *gin.Context) {
	// TODO 用户名必须传
	var data model.User
	_ = ctx.ShouldBindJSON(&data)
	code = model.IsUserExist(data.Username)
	if code == errormsg.SUCCESS {
		// 写进数据库
		model.AddUser(&data)
	}
	if code == errormsg.ERROR_USERNAME_USED {
		// 用户已存在 返回1001
		code = errormsg.ERROR_USERNAME_USED
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// DeleteUserController 删除用户
func DeleteUserController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code = model.DeleteUser(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   nil,
		"msg":    errormsg.GetErrorMsg(code),
	})

}

// TODO 查询单个用户

// GetUserListController 查询用户列表
func GetUserListController(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	data := model.GetUserList(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status": errormsg.SUCCESS,
		"data":   data,
		"msg":    errormsg.GetErrorMsg(errormsg.SUCCESS),
	})
}

// EditUserController TODO 编辑用户
func EditUserController(ctx *gin.Context) {

}
