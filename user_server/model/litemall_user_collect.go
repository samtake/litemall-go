package model

import "github.com/jinzhu/gorm"

//收藏表

type Collect struct {
	gorm.Model
	UserId      uint `gorm:"type:mediumint(8) unsigned; NOT NULL DEFAULT '0';index:user_id"`
	ValueId     uint `gorm:"type:mediumint(8) unsigned; NOT NULL DEFAULT '0';index:user_id"`
	IsAttention uint `gorm:"type:int(1) unsigned; NOT NULL DEFAULT '0';COMMENT '是否是关注';index:user_id"`
	TypeId      uint `gorm:"type:int(2) unsigned; NOT NULL DEFAULT '0'"`
	Deleted     uint `gorm:"type:tinyint(1) unsigned; NOT NULL DEFAULT '0'"`
}

/*
DROP TABLE IF EXISTS `nideshop_collect`;
CREATE TABLE `nideshop_collect` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `value_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0',
  `is_attention` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是关注',
  `type_id` int(2) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `goods_id` (`value_id`),
  KEY `is_attention` (`is_attention`)
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8mb4;
*/
