package model

import (
	"ginserver/utils"
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	ImageName string `gorm:"type:varchar(100);not null" json:"imagename"`
	ImageDesc string `gorm:"type:varchar(100);not null" json:"imagedesc"`
	ImageTag  string `gorm:"type:varchar(100);not null" json:"imagetag"`
	ImagePath string `gorm:"type:varchar(100);not null" json:"imagepath"`
	AuthID    int    `gorm:"default:1" json:"userid"`
	ImageAuth User   `gorm:"ForeignKey:AuthID" json:"imageauth"`
}

func UploadImg(img *Image) int {
	err := DB.Create(&img).Error

	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}
