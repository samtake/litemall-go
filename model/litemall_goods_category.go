package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name         string `gorm:"type:varchar(90) NOT NULL DEFAULT ''"`
	Keywords     string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	FrontDesc    string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	ParentId     uint   `gorm:"type:tinyint(11) unsigned NOT NULL DEFAULT '1';index:parent_id"`
	SortOrder    uint   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '50'"`
	ShowIndex    uint   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '1'"`
	IsShow       uint   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '1';index:is_show"`
	BannerUrl    string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	IconUrl      string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	ImgUrl       string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	WapBannerUrl string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	Level        string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	Type         uint   `gorm:"type:tinyint(11) unsigned NOT NULL DEFAULT '1'"`
	FrontName    string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
}

/*
----------------------
-- Table structure for nideshop_category
-- ----------------------------
DROP TABLE IF EXISTS `nideshop_category`;
CREATE TABLE `nideshop_category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(90) NOT NULL DEFAULT '',
  `keywords` varchar(255) NOT NULL DEFAULT '',
  `front_desc` varchar(255) NOT NULL DEFAULT '',
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0',
  `sort_order` tinyint(1) unsigned NOT NULL DEFAULT '50',
  `show_index` tinyint(1) NOT NULL DEFAULT '0',
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `banner_url` varchar(255) NOT NULL DEFAULT '',
  `icon_url` varchar(255) NOT NULL,
  `img_url` varchar(255) NOT NULL,
  `wap_banner_url` varchar(255) NOT NULL,
  `level` varchar(255) NOT NULL,
  `type` int(11) NOT NULL DEFAULT '0',
  `front_name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1036008 DEFAULT CHARSET=utf8mb4;

*/
