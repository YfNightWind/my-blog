package v1

import (
	"github.com/gin-gonic/gin"
	"my-blog/middleware"
	"my-blog/model"
	"my-blog/utils/errormsg"
	"net/http"
)

// LoginController 登录成功吼返回token
func LoginController(ctx *gin.Context) {
	var data model.User
	var token string
	_ = ctx.ShouldBindJSON(&data)

	code := model.VerifyLogin(data.Username, data.Password)

	if code == errormsg.SUCCESS {
		token, code = middleware.GenerateToken(data.Username)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"token":  token,
		"msg":    errormsg.GetErrorMsg(code),
	})
}
