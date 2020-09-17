package utils

import "log"

var logger log.Logger

func init() {
	// 设置 log 输出格式， 日期 东八区时间 文件行号 内容
	logger.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
}

func GetLogger() log.Logger {
	return logger
}
