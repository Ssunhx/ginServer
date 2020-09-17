package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 全局 DB 对象
var DB *gorm.DB
var err error

func InitDB() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=true&loc=Local", UserName,
		Password, Host, DBName)

	DB, err = gorm.Open(mysql.New(mysql.Config{
		//DriverName:                "",	// 可以使用自定义的驱动
		DSN: dsn, // 数据库链接配置
		//Conn:                      nil,	// 连接现有的数据库
		SkipInitializeWithVersion: false, // 根据当前 mysql 版本自动配置
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  false, // 禁用 datetime 精度， 5.6之前不支持
		DontSupportRenameIndex:    false, // 重命名索引时采用删除并新建的方式， 5.7之前不支持
		DontSupportRenameColumn:   false, // 用 change 重命名列， 8之前和 mariadb 不支持
	}), &gorm.Config{})

	if err != nil {
		logger.Println("connect database error, please check your config")
	}

	// 数据库连接池设置
	sqlDB, _ := DB.DB()

	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// 打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)

	// 连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Second * 10)
}

func GetDB() *gorm.DB {
	return DB
}

// 迁移数据库表
func MigrateTables() {
	//DB.AutoMigrate()
}
