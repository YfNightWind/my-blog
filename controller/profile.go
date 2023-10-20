package v1

import (
	"net/http"
	"strconv"

	"github.com/YfNightWind/my-blog/model"
	"github.com/YfNightWind/my-blog/utils/errormsg"
	"github.com/gin-gonic/gin"
)

// GetProfileController 获取个人信息
func GetProfileController(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))

	data, code := model.GetProfile(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errormsg.GetErrorMsg(code),
	})
}

// UpdateProfileController 更新个人信息
func UpdateProfileController(ctx *gin.Context) {
	var data model.Profile

	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&data)
	code = model.UpdateProfile(id, &data)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetErrorMsg(code),
	})
}
