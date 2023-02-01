package model

import "github.com/jinzhu/gorm"

type AfterSale struct {
	gorm.Model
	AftersaleSn string  `gorm:"type:varchar(63)  NOT NULL DEFAULT ''"`
	OrderId     uint    `gorm:"type:int(11) unsigned NOT NULL DEFAULT '0'"`
	Type        uint    `gorm:"type:tinyint(11) unsigned NOT NULL DEFAULT '0'"`
	Reason      string  `gorm:"type:varchar(31)  NOT NULL DEFAULT ''"`
	Amount      float64 `gorm:"decimal(10,2) unsigned NOT NULL DEFAULT '0.00'"`
	Pictures    string  `gorm:"type:varchar(1023)  NOT NULL DEFAULT ''"`
	Comment     string  `gorm:"type:varchar(511)  NOT NULL DEFAULT ''"`
	Status      uint    `gorm:"type:smallint(5) NOT NULL DEFAULT '0'"`
	HandleTime  uint    `gorm:"type:bigint(12) unsigned NOT NULL DEFAULT '0'"`
}
