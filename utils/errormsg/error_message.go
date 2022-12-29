package errormsg

const (
	SUCCESS = 200
	ERROR   = 500
	// cpde = 100...用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007

	// code = 200...文章模块的错误

	// code = 300...分类模块的错误
	ERROR_CATEGORYNAME_USED = 3001
)

var codeMsg = map[int]string{
	SUCCESS:                 "成功！",
	ERROR:                   "错误",
	ERROR_USERNAME_USED:     "用户名已被使用",
	ERROR_PASSWORD_WRONG:    "用户密码错误",
	ERROR_USER_NOT_EXIST:    "该用户不存在",
	ERROR_TOKEN_NOT_EXIST:   "token不存在",
	ERROR_TOKEN_RUNTIME:     "token过期",
	ERROR_TOKEN_WRONG:       "token错误",
	ERROR_TOKEN_TYPE_WRONG:  "token格式错误",
	ERROR_CATEGORYNAME_USED: "该分类已存在",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
