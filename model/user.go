package model

import (
	"errors"
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

// 使用钩子函数加密用户密码，不用显示调用，只需要指向结构体即可
// 其他类似的钩子函数还有：
// create object ：BeforeSave、BeforeCreate、AfterCreate、AfterSave
// update object：BeforeSave、BeforeUpdate、AfterUpdate、 AfterSave
// delete object：BeforeDelete、AfterDelete
// query object: AfterFind

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	pwd := utils.ScryptPassword(u.Password)
	if pwd == "" {
		return errors.New("scrypt password error")
	}
	u.Password = pwd
	return nil
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
