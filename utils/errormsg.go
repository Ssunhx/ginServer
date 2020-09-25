package utils

const (
	SUCCESS = 200
	ERROR   = 500

	// user error
	USERNAME_USED = 1001
)

var codeMsg = map[int]string{
	SUCCESS:       "SUCCESS",
	ERROR:         "ERROR",
	USERNAME_USED: "用户名已存在",
}

func GetErrMsg(code int) string {
	value, err := codeMsg[code]
	if !err {
		logger.Println("key error")
		return ""
	}
	return value
}
