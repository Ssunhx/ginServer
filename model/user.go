package model

import (
	"ginserver/utils"
	"gorm.io/gorm"
)

type Role struct {
	ID   int
	Name string
}

type User struct {
	// 包含了默认的四个字段
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20); default:123456" json:"password"`
	Mobile   string `gorm:"type:char(11)" json:"mobile"`
	Sex      int    `gorm:"type:tinyint(1); default:0" json:"sex"`
	Age      string `gorm:"type:char(13)" json:"age"`
	//Role  	 string `gorm:"type:tinyint(1); default:0" json:"role"`
	RoleID int  `gorm:"default:1" json:"roleid"`
	Role   Role `gorm:"ForeignKey:RoleID"`
}

func CheckUser(username string) int {
	var user User
	DB.Where("username = ?", username).First(&user)
	// 用户名已存在
	if user.ID > 0 {
		return utils.USERNAME_USED
	}
	return utils.SUCCESS
}

func CreateUser(user *User) int {
	err := DB.Create(&user).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}

func CheckLogin(username string, password string) int {
	var user User
	DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return utils.USER_NOT_EXIST
	}

	if utils.ScryptPassword(password) != user.Password {
		return utils.USER_PASSWORD_ERROR
	}
	return utils.SUCCESS
}
