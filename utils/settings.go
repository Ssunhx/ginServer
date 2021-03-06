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

	// 七牛云配置
	QiniuAccessKey  string
	QiniuSecretKey  string
	QiniuBucketName string
	QiniuDomain     string
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
	LoadQiNiu(file)
}

// 初始化 server 配置
func LoadServer(file *ini.File) {
	mode := file.Section("server").Key("Mode").String()
	if mode == "1" {
		AppMode = "debug"
	} else if mode == "0" {
		AppMode = "test"
	} else if mode == "2" {
		AppMode = "release"
	} else {
		AppMode = ""
	}
	HttpPort = file.Section("server").Key("HttpPort").String()
}

// 初始化 DB 配置
func LoadDB(file *ini.File) {
	Host = file.Section("database").Key("Host").String()
	Port = file.Section("database").Key("Port").String()
	UserName = file.Section("database").Key("UserName").String()
	Password = file.Section("database").Key("Password").String()
	DBName = file.Section("database").Key("DBName").String()
}

// 初始化七牛云 oss
func LoadQiNiu(file *ini.File) {
	QiniuAccessKey = file.Section("qiniu").Key("accessKey").String()
	QiniuSecretKey = file.Section("qiniu").Key("secretKey").String()
	QiniuBucketName = file.Section("qiniu").Key("qiniu_bucket_name").String()
	QiniuDomain = file.Section("qiniu").Key("qiniu_domain").String()
}
