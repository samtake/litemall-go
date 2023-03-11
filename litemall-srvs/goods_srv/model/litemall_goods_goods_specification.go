package model

//商品规格表

type GoodsSpecification struct {
	BaseModel
	GoodsId       int32  `gorm:"type:mediumint(8) unsigned NOT NULL DEFAULT '0'"`
	Specification string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	Value         string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	PicUrl        string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
}
