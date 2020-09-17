package utils

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	// http 设置
	AppMode  string
	HttpPort string

	// database 配置
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
)

func init() {
	// 初始化参数

	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("load config file error")
		logger.Println("load config file error")
	}
	LoadServer(file)
	LoadDB(file)
}

// 初始化 server 配置
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").String()
	HttpPort = file.Section("server").Key("httpPort").String()
}

// 初始化 DB 配置
func LoadDB(file *ini.File) {
	Host = file.Section("database").Key("Host").String()
	Port = file.Section("database").Key("Port").String()
	UserName = file.Section("database").Key("UserName").String()
	Password = file.Section("database").Key("Password").String()
	DBName = file.Section("database").Key("DBName").String()
}
