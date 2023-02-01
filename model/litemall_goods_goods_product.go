package model

import "github.com/jinzhu/gorm"

type GoodsProduct struct {
	gorm.Model
	GoodsId               uint    `gorm:"type:mediumint(8) unsigned NOT NULL DEFAULT '0'"`
	GoodsSpecificationIds string  `gorm:"type:varchar(50) NOT NULL DEFAULT ''"`
	GoodsSn               string  `gorm:"type:varchar(60) NOT NULL DEFAULT ''"`
	GoodsNumber           uint    `gorm:"type:mediumint(8) unsigned NOT NULL DEFAULT '0'"`
	RetailPrice           float64 `gorm:"type:decimal(10,2) unsigned NOT NULL DEFAULT '0.00'"`
}

/*
-- ----------------------------
-- Table structure for nideshop_product
-- ----------------------------
DROP TABLE IF EXISTS `nideshop_product`;
CREATE TABLE `nideshop_product` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `goods_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `goods_specification_ids` varchar(50) NOT NULL DEFAULT '',
  `goods_sn` varchar(60) NOT NULL DEFAULT '',
  `goods_number` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `retail_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=245 DEFAULT CHARSET=utf8mb4;
*/
