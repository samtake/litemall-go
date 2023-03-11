package model

import "github.com/jinzhu/gorm"

// 购物车

type Cart struct {
	gorm.Model
	UserId      uint    `gorm:"type:mediumint(45) unsigned; NOT NULL DEFAULT '0'"`
	SessionId   uint    `gorm:"type:mediumint(32) unsigned; NOT NULL DEFAULT ''"`
	GoodsId     uint    `gorm:"type:mediumint(8) unsigned NOT NULL DEFAULT '0'"`
	GoodsSn     string  `gorm:"type:varchar(60) NOT NULL DEFAULT ''"`
	ProductId   uint    `gorm:"type:mediumint(8) unsigned NOT NULL DEFAULT '0'"`
	GoodsName   string  `gorm:"type:varchar(120) NOT NULL DEFAULT ''"`
	MarketPrice float64 `gorm:"type:decimal(10,2) unsigned NOT NULL DEFAULT '0.00'"`
}

/*
CREATE TABLE `nideshop_cart` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `session_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',


  `goods_id` mediumint(8) unsigned NOT NULL DEFAULT '0',

  `goods_sn` varchar(60) NOT NULL DEFAULT '',

  `product_id` mediumint(8) unsigned NOT NULL DEFAULT '0',


  `goods_name` varchar(120) NOT NULL DEFAULT '',

  `market_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00',
  `retail_price` decimal(10,2) NOT NULL DEFAULT '0.00',
  `number` smallint(5) unsigned NOT NULL DEFAULT '0',
  `goods_specifition_name_value` text NOT NULL COMMENT '规格属性组成的字符串，用来显示用',
  `goods_specifition_ids` varchar(60) NOT NULL DEFAULT '' COMMENT 'product表对应的goods_specifition_ids',
  `checked` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `list_pic_url` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `session_id` (`session_id`)
) ENGINE=InnoDB AUTO_INCREMENT=99 DEFAULT CHARSET=utf8mb4;
*/
