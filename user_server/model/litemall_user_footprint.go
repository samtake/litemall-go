package model

import "github.com/jinzhu/gorm"

//足迹

type Footprint struct {
	gorm.Model
	UserId  uint `gorm:"type:int(11) unsigned; NOT NULL DEFAULT '0'"`
	GoodsId uint `gorm:"type:int(11) unsigned; NOT NULL DEFAULT '0'"`
	Deleted uint `gorm:"type:tinyint(1) unsigned; NOT NULL DEFAULT '0'"`
}

/*
DROP TABLE IF EXISTS `nideshop_footprint`;
CREATE TABLE `nideshop_footprint` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `goods_id` int(11) NOT NULL DEFAULT '0',
  `add_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=93 DEFAULT CHARSET=utf8mb4;
*/
