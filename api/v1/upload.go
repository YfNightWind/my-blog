package v1

import (
	"github.com/gin-gonic/gin"
	"my-blog/server"
	"my-blog/utils/errormsg"
	"net/http"
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
