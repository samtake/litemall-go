package model

import "github.com/jinzhu/gorm"

//广告、轮播图
type ad struct {
	gorm.Model
	AdPositionId uint   `gorm:"type:smallint(5) unsigned NOT NULL DEFAULT '0';index:user_id;index:position_id"`
	MediaType    uint   `gorm:"type:tinyint(3) unsigned NOT NULL DEFAULT '0'"`
	Name         string `gorm:"type:varchar(60) NOT NULL DEFAULT ''"`
	Link         string `gorm:"type:varchar(60) NOT NULL DEFAULT ''"`
	ImageUrl     string `gorm:"type:text NOT NULL"`
	Content      string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	EndTime      uint   `gorm:"type:int(11) NOT NULL DEFAULT '0'"`
	Enabled      uint   `gorm:"type:int(3) NOT NULL DEFAULT '1'"`
}

/*
-- ----------------------------
-- Table structure for nideshop_ad
-- ----------------------------
DROP TABLE IF EXISTS `nideshop_ad`;
CREATE TABLE `nideshop_ad` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `ad_position_id` smallint(5) unsigned NOT NULL DEFAULT '0',
  `media_type` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `name` varchar(60) NOT NULL DEFAULT '',
  `link` varchar(255) NOT NULL DEFAULT '',
  `image_url` text NOT NULL,
  `content` varchar(255) NOT NULL DEFAULT '',
  `end_time` int(11) NOT NULL DEFAULT '0',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `position_id` (`ad_position_id`),
  KEY `enabled` (`enabled`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
*/
