package model

import "github.com/jinzhu/gorm"

//用户的优惠券

type CouponUser struct {
	gorm.Model
	CouponId     uint   `gorm:"type:tinyint(3) unsigned; NOT NULL DEFAULT '0'"`
	CouponNumber string `gorm:"type:varchar(20);not null DEFAULT ''"`
	UserId       uint   `gorm:"type:int(11) unsigned; NOT NULL DEFAULT '0'; index"`
	OrderId      uint   `gorm:"type:int(11) unsigned NOT NULL DEFAULT '0'"`
}

/*
DROP TABLE IF EXISTS `nideshop_user_coupon`;
CREATE TABLE `nideshop_user_coupon` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `coupon_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `coupon_number` varchar(20) NOT NULL DEFAULT '',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `used_time` int(10) unsigned NOT NULL DEFAULT '0',
  `order_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4;
*/
