package model

import "github.com/jinzhu/gorm"

//评论

type Comment struct {
	gorm.Model
	ValueId    uint   `gorm:"type:tinyint(1) unsigned; NOT NULL DEFAULT '0';index"`
	Content    string `gorm:"type:varchar(6550); NOT NULL DEFAULT ''"`
	Status     uint   `gorm:"type:tinyint(3) unsigned; NOT NULL DEFAULT '0'"`
	UserId     uint   `gorm:"type:tinyint(11) unsigned; NOT NULL DEFAULT '0'"`
	NewContent string `gorm:"type:varchar(6550); NOT NULL DEFAULT ''"`
}

/*
DROP TABLE IF EXISTS `nideshop_comment`;
CREATE TABLE `nideshop_comment` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `value_id` int(11) unsigned NOT NULL DEFAULT '0',
  `content` varchar(6550) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '储存为base64编码',
  `add_time` bigint(12) unsigned NOT NULL DEFAULT '0',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0',

  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `new_content` varchar(6550) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `id_value` (`value_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1006 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
*/
