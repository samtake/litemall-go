package model

import "github.com/jinzhu/gorm"

//区域

type Region struct {
	gorm.Model
	ParentId uint   `gorm:"type:mediumint(8) unsigned; NOT NULL DEFAULT '0';index:user_id;index:parent_id"`
	Name     string `gorm:"type:varchar(50) NOT NULL DEFAULT ''"`
	Type     uint   `gorm:"type:tinyint(1) NOT NULL DEFAULT '2';index:type"`
	AgencyId uint   `gorm:"type:smallint(5) unsigned NOT NULL DEFAULT '0';index:agency_id"`
}

/*
-- ----------------------------
-- Table structure for nideshop_region
-- ----------------------------
DROP TABLE IF EXISTS `nideshop_region`;
CREATE TABLE `nideshop_region` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` smallint(5) unsigned NOT NULL DEFAULT '0',
  `name` varchar(120) NOT NULL DEFAULT '',
  `type` tinyint(1) NOT NULL DEFAULT '2',
  `agency_id` smallint(5) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`),
  KEY `region_type` (`type`),
  KEY `agency_id` (`agency_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4044 DEFAULT CHARSET=utf8mb4;
*/
