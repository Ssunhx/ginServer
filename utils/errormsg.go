package utils

const (
	SUCCESS = 200
	ERROR   = 500
)

var codeMsg = map[int]string{
	SUCCESS: "SUCCESS",
	ERROR:   "ERROR",
}

func GetErrMsg(code int) string {
	value, err := codeMsg[code]
	if !err {
		logger.Println("key error")
		return ""
	}
	return value
}
