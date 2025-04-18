package server

import (
	"context"
	"mime/multipart"

	"github.com/YfNightWind/my-blog/utils"
	"github.com/YfNightWind/my-blog/utils/errormsg"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// 七牛云对象存储Go SDK
// https://developer.qiniu.com/kodo/1238/go
//

var (
	AccessKey = utils.AccessKey
	SecretKey = utils.SecretKey
	Bucket    = utils.Bucket
	ImgUrl    = utils.QiNiuServer
)

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{Scope: Bucket}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong, // 使用华东地区机房
		UseHTTPS:      true,
		UseCdnDomains: true,
	}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, nil)
	if err != nil {
		return "", errormsg.ERROR
	}

	url := ImgUrl + ret.Key
	return url, errormsg.SUCCESS
}
