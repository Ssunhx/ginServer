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
	UPLOAD_FILE_IS_EMPTY    = 2001

	NO_PER_TO_DELETE_IMAGE = 3001
	DELETE_IMAGE_SUCCESS   = 3002
	IMAGE_NOT_FOUND        = 3003
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
	UPLOAD_FILE_IS_EMPTY:    "未选择文件",
	NO_PER_TO_DELETE_IMAGE:  "没有权限删除图片",
	DELETE_IMAGE_SUCCESS:    "图片删除成功",
	IMAGE_NOT_FOUND:         "图片未找到",
}

func GetErrMsg(code int) string {
	value, err := codeMsg[code]
	if !err {
		logger.Println("key error")
		return ""
	}
	return value
}
