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

// EditUserController 编辑用户
// TODO 判断用户已被软删除之后如何解决
func EditUserController(ctx *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&data)

	// 修改用户名时候，检查更新后的用户名是否存在
	code = model.IsUserExist(data.Username)

	if code == errormsg.SUCCESS {
		model.EditUser(id, &data)
	}

	if code == errormsg.ERROR_USERNAME_USED {
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetErrorMsg(code),
	})
}
