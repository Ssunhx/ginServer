package model

import (
	"ginserver/utils"
	"gorm.io/gorm"
)

type Image struct {
	// 返回 json 时，不想要传给前端的字段用 -
	gorm.Model
	ImageName string `gorm:"type:varchar(100);not null" json:"imagename"`
	ImageDesc string `gorm:"type:varchar(100);not null" json:"imagedesc"`
	ImageTag  string `gorm:"type:varchar(100);not null" json:"imagetag"`
	ImagePath string `gorm:"type:varchar(100);not null" json:"imagepath"`
	AuthID    int    `gorm:"default:1" json:"authid"`
	ImageAuth User   `gorm:"ForeignKey:AuthID" json:"-"`
}

// 上传image
func UploadImg(img *Image) int {
	err := DB.Create(&img).Error

	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}

// 获取用户的图片
func GetImgByUserId(userid int) []Image {
	var images []Image
	err := DB.Where("auth_id", userid).Find(&images).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return images
}

// 根据imageid 获取图片信息
func GetImgById(imgid int) Image {
	var img Image
	DB.Where("id", imgid).First(&img)
	return img
}

// 删除图片
func DeleteImg(imageid int) int {
	var img Image
	err := DB.Where("id", imageid).Delete(&img).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.DELETE_IMAGE_SUCCESS
}

// 获取已经被删除的图片
func GetDeleted(userid int) []Image {
	var images []Image
	err := DB.Unscoped().Where("auth_id", userid).Find(&images).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return images
}
