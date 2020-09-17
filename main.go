package main

import "ginserver/utils"

func main() {
	// 初始化 DB
	utils.InitDB()

	// 迁移表结构
	utils.MigrateTables()
}
