package controller

import (
	"net/http"
	"strconv"

	"github.com/YfNightWind/my-blog/model"
	"github.com/YfNightWind/my-blog/utils/errormsg"
	"github.com/YfNightWind/my-blog/utils/validator"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	model.User
}

// AddUserController 添加用户
func (u UserController) AddUserController(ctx *gin.Context) {
	var data model.User
	var msg string
	_ = ctx.ShouldBindJSON(&data)

	msg, code = validator.Validate(&data)
	if code != errormsg.SUCCESS {
		ctx.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    msg,
		})
		return
	}

	code = model.IsUserExist(data.Username)
	if code == errormsg.SUCCESS {
		// 写进数据库
		model.AddUser(&data)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// DeleteUserController 删除用户
func (u UserController) DeleteUserController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code = model.DeleteUser(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   nil,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// GetUserController 查询单个用户
func (u UserController) GetUserController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := u.GetUser(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"total":  1,
		"msg":    errormsg.GetErrorMsg(errormsg.SUCCESS),
	})
}

// GetUserListController 查询用户列表
func (u UserController) GetUserListController(ctx *gin.Context) {
	username := ctx.Query("username")
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	data, total := model.GetUserList(username, pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status": errormsg.SUCCESS,
		"data":   data,
		"total":  total,
		"msg":    errormsg.GetErrorMsg(errormsg.SUCCESS),
	})
}

// EditUserController 编辑用户
// TODO 判断用户已被软删除之后如何解决
func (u UserController) EditUserController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&u.User)

	// 修改用户名时候，检查更新后的用户名是否存在
	code = model.UpdateUser(id, u.Username)

	if code == errormsg.SUCCESS {
		model.EditUser(id, &u.User)
	}

	if code == errormsg.ERROR_USERNAME_USED {
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// ChangePasswordController 修改密码
func (u UserController) ChangePasswordController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&u.User)

	code = u.ChangePassword(id, &u.User)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// AdjustToolsPageAccessController 修改工具页面权限
func (u UserController) AdjustToolsPageAccessController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))

	_ = ctx.ShouldBindJSON(&u.User)
	err, code := u.AdjustToolsPageAccess(id)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetErrorMsg(int(code)),
	})
}
