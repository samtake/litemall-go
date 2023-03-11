package model

import (
	"github.com/jinzhu/gorm"
)

//平台优惠券

type Coupon struct {
	gorm.Model
	Name           string  `gorm:"varchar(50) NOT NULL DEFAULT ''"`
	TypeMoney      float64 `gorm:"decimal(10,2) NOT NULL DEFAULT '0.00'"`
	SendType       float64 `gorm:"decimal(10,2) NOT NULL DEFAULT '0.00'"`
	MiniAmount     float64 `gorm:"decimal(10,2) unsigned NOT NULL DEFAULT '0.00'"`
	MaxAmount      float64 `gorm:"decimal(10,2) unsigned NOT NULL DEFAULT '0.00'"`
	SendStartDate  float64 `gorm:"decimal(10,2) unsigned NOT NULL DEFAULT '0.00'"`
	SendEndDate    float64 `gorm:"decimal(10,2) unsigned NOT NULL DEFAULT '0.00'"`
	UseStartDate   float64 `gorm:"decimal(10,2) unsigned NOT NULL DEFAULT '0.00'"`
	UseEndDate     float64 `gorm:"decimal(10,2) unsigned NOT NULL DEFAULT '0.00'"`
	MinGoodsAmount float64 `gorm:"decimal(10,2) unsigned NOT NULL DEFAULT '0.00'"`
}

/*
-- ----------------------------
-- Table structure for nideshop_coupon
-- ----------------------------
DROP TABLE IF EXISTS `nideshop_coupon`;
CREATE TABLE `nideshop_coupon` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(60) NOT NULL DEFAULT '',
  `type_money` decimal(10,2) NOT NULL DEFAULT '0.00',
  `send_type` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `min_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00',
  `max_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00',
  `send_start_date` int(11) NOT NULL DEFAULT '0',
  `send_end_date` int(11) NOT NULL DEFAULT '0',
  `use_start_date` int(11) NOT NULL DEFAULT '0',
  `use_end_date` int(11) NOT NULL DEFAULT '0',
  `min_goods_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
*/
