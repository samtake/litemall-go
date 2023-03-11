package model

import "github.com/jinzhu/gorm"

//系统相关

type System struct {
	gorm.Model
	KeyName  string `gorm:"varchar(255) NOT NULL DEFAULT ''"`
	KeyValue string `gorm:"varchar(255) NOT NULL DEFAULT ''"`
}
