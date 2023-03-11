package model

import "github.com/jinzhu/gorm"

//

type Storage struct {
	gorm.Model
	Key  string `gorm:"type:varchar(63) NOT NULL DEFAULT ''"`
	Name string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	Type string `gorm:"type:varchar(20) NOT NULL DEFAULT ''"`
	Size uint   `gorm:"type:tinyint(11) unsigned NOT NULL DEFAULT '0'"`
	Url  string `gorm:"type:varchar(255) NOT NULL DEFAULT '';COMMENT '--'"`
}
