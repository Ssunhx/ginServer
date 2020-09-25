package model

import (
	"ginserver/utils"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = utils.GetDB()
}
