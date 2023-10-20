package v1

import (
	"net/http"

	"github.com/YfNightWind/my-blog/server"
	"github.com/YfNightWind/my-blog/utils/errormsg"
	"github.com/gin-gonic/gin"
)

func UploadController(ctx *gin.Context) {
	file, fileHeader, _ := ctx.Request.FormFile("file")

	fileSize := fileHeader.Size

	url, code := server.UploadFile(file, fileSize)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetErrorMsg(code),
		"url":    url,
	})
}
