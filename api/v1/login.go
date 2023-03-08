package v1

import (
	"github.com/gin-gonic/gin"
	"my-blog/middleware"
	"my-blog/model"
	"my-blog/utils/errormsg"
	"net/http"
)

// AdminLoginController 后台登录
func AdminLoginController(ctx *gin.Context) {
	var data model.User
	var token string
	_ = ctx.ShouldBindJSON(&data)

	code := model.VerifyAdminLogin(data.Username, data.Password)

	if code == errormsg.SUCCESS {
		token, code = middleware.GenerateToken(data.Username)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"token":  token,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// FrontLoginController 后台登录
func FrontLoginController(ctx *gin.Context) {
	var data model.User
	var token string
	_ = ctx.ShouldBindJSON(&data)

	data, code := model.VerifyFrontLogin(data.Username, data.Password)

	if code == errormsg.SUCCESS {
		token, code = middleware.GenerateToken(data.Username)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"token":  token,
		"data":   data.Username,
		"id":     data.ID,
		"msg":    errormsg.GetErrorMsg(code),
	})
}
