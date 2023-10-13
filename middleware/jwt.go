package middleware

import (
	"github.com/YfNightWind/my-blog/utils"
	"github.com/YfNightWind/my-blog/utils/errormsg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(username string) (string, int) {
	userClaim := MyClaims{
		// 使用用户名生成
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Hour * time.Duration(1))), // 10天
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "AlexLin",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errormsg.ERROR
	}
	return token, errormsg.SUCCESS

}

// VerifyToken 验证token
func VerifyToken(reqToken string) (*MyClaims, int) {
	var claims MyClaims
	token, _ := jwt.ParseWithClaims(reqToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	// https://pkg.go.dev/github.com/golang-jwt/jwt/v4#ParseWithClaims
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, errormsg.SUCCESS
	} else {
		return nil, errormsg.ERROR
	}
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		tokenHeader := ctx.GetHeader("Authorization")

		// 正常的应该长这样子
		// Bearer xxx.xxx.xxx
		//
		// 分割之后，长度为2
		// [Bearer xxx.xxx.xxx]
		//
		// 以此来判断token是否有效

		//
		// TODO 增加Token过期后生成新的Token
		//

		if tokenHeader == "" {
			code = errormsg.ERROR_TOKEN_NOT_EXIST
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errormsg.GetErrorMsg(code),
			})
			ctx.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			code = errormsg.ERROR_TOKEN_TYPE_WRONG
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errormsg.GetErrorMsg(code),
			})
			ctx.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = errormsg.ERROR_TOKEN_TYPE_WRONG
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errormsg.GetErrorMsg(code),
			})
			ctx.Abort()
			return
		}

		key, tCode := VerifyToken(checkToken[1])
		if tCode == errormsg.ERROR {
			code = errormsg.ERROR_TOKEN_WRONG
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errormsg.GetErrorMsg(code),
			})
			ctx.Abort()
			return
		}

		if jwt.NewNumericDate(time.Now()).Unix() > key.ExpiresAt.Unix() {
			code = errormsg.ERROR_TOKEN_RUNTIME
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errormsg.GetErrorMsg(code),
			})
			ctx.Abort()
			return
		}

		ctx.Set("username", key.Username)
		ctx.Next()
	}
}
