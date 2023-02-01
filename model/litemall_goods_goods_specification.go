package model

import "github.com/jinzhu/gorm"

type GoodsSpecification struct {
	gorm.Model
	GoodsId       uint   `gorm:"type:mediumint(8) unsigned NOT NULL DEFAULT '0'"`
	Specification string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	Value         string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	PicUrl        string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
}
