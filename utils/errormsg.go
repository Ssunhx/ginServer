package utils

const (
	SUCCESS = 200
	ERROR   = 500

	// user error
	USERNAME_USED           = 1001
	USER_NOT_EXIST          = 1002
	USER_PASSWORD_ERROR     = 1003
	USER_TOKEN_NOT_EXIST    = 1004
	USER_TOKEN_VERIFY_ERROR = 1005
	USER_TOKEN_EXPIRED      = 1006
)

var codeMsg = map[int]string{
	SUCCESS:                 "SUCCESS",
	ERROR:                   "ERROR",
	USERNAME_USED:           "用户名已存在",
	USER_NOT_EXIST:          "用户不存在",
	USER_PASSWORD_ERROR:     "用户密码错误",
	USER_TOKEN_NOT_EXIST:    "用户token不存在",
	USER_TOKEN_VERIFY_ERROR: "用户token验证失败",
	USER_TOKEN_EXPIRED:      "用户token过期",
}

func GetErrMsg(code int) string {
	value, err := codeMsg[code]
	if !err {
		logger.Println("key error")
		return ""
	}
	return value
}
