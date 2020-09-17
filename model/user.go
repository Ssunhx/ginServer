package model

import "gorm.io/gorm"

type Role struct {
	ID   int
	Name string
}

type User struct {
	// 包含了默认的四个字段
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20); default:123456" json:"password"`
	Mobile   string `gorm:"type:char(11)" json:"mobile"`
	Sex      int    `gorm:"type:tinyint(1); default:0" json:"sex"`
	Age      string `gorm:"type:char(13)" json:"age"`
	//Role  	 string `gorm:"type:tinyint(1); default:0" json:"role"`
	RoleID int  `gorm:"default:1" json:"roleid"`
	Role   Role `gorm:"ForeignKey:RoleID"`
}

func CheckUser(username string) {

}
