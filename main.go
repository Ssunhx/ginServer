package main

import (
	"ginserver/model"
	"ginserver/router"
)

func main() {
	// 初始化 DB
	model.InitDB()

	// 迁移表结构
	model.MigrateTables()

	router.InitRouter()
}
