package model

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ImageName string `gorm:"type:varchar(100);not null" json:"imagename"`
	ImageDesc string `gorm:"type:varchar(100);not null" json:"imagedesc"`
	ImageTag  string `gorm:"type:varchar(100);not null" json:"imagetag"`
	ImagePath string `gorm:"type:varchar(100);not null" json:"imagepath"`
	AuthID    int    `gorm:"default:1" json:"userid"`
	ImageAuth User   `gorm:"ForeignKey:AuthID" json:"imageauth"`
}
