package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Log() gin.HandlerFunc {
	// TODO 加入SQL语句展示
	log := logrus.New()
	filePath := "log/"
	linkName := "latest-log.log"
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err: ", err)
	}
	log.Out = src

	log.SetLevel(logrus.DebugLevel)

	// 由于logrus不支持按日志按文件生成引入file-rotatelogs来支持日志按文件生成
	logWriter, _ := rotateLogs.New(
		filePath+"blog-%Y%m%d.log",
		rotateLogs.WithMaxAge(14*24*time.Hour), // 两周清除一次
		rotateLogs.WithRotationTime(1),         // 1天分割一次
		rotateLogs.WithLinkName(linkName),      // 根目录下生成最新的log
	)

	// 引入lfshook来组合
	writeMap := lfshook.WriterMap{
		logrus.PanicLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.DebugLevel: logWriter,
		logrus.TraceLevel: logWriter,
	}

	// 日志输出前，执行钩子的内容(输出时间)
	timeFormat := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 对时间格式化
	})
	log.AddHook(timeFormat)

	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		stopTime := time.Since(startTime)

		// 请求时间
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))

		// 主机名
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		// 状态码
		statusCode := ctx.Writer.Status()

		// 请求IP
		clientIP := ctx.ClientIP()

		// 客户端信息
		userAgent := ctx.Request.UserAgent()

		// 请求大小
		dataSize := ctx.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}

		// 请求方法
		method := ctx.Request.Method

		// 请求路径
		reqUri := ctx.Request.RequestURI

		// 记录
		entry := log.WithFields(logrus.Fields{
			"HostName":  hostName,
			"Status":    statusCode,
			"SpendTime": spendTime,
			"Method":    method,
			"IP":        clientIP,
			"Path":      reqUri,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})

		// 处理一些可能发生的错误
		if len(ctx.Errors) > 0 {
			entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		switch {
		case statusCode >= 500:
			entry.Error()
		case statusCode >= 400:
			entry.Warn()
		default:
			entry.Info()
		}
	}
}
