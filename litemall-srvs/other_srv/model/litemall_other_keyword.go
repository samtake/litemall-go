package model

import "github.com/jinzhu/gorm"

//热闹关键词表

type Keywords struct {
	gorm.Model
	Keyword   string `gorm:"type:varchar(90) NOT NULL DEFAULT ''"`
	IsHot     uint   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '0'"`
	IsDefault uint   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '0'"`
	IsShow    uint   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '1'"`
	SortOrder uint   `gorm:"type:tinyint(1)1 unsigned NOT NULL DEFAULT '100'"`
	SchemeUrl string `gorm:"type:varchar(255) NOT NULL DEFAULT '';COMMENT '关键词的跳转链接'"`
	Type      uint   `gorm:"type:tinyint(11) unsigned NOT NULL DEFAULT '0'"`
}

/*
-- ----------------------------
-- Table structure for nideshop_keywords
-- ----------------------------
DROP TABLE IF EXISTS `nideshop_keywords`;
CREATE TABLE `nideshop_keywords` (
  `keyword` varchar(90) NOT NULL DEFAULT '',
  `is_hot` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `is_default` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `sort_order` int(11) unsigned NOT NULL DEFAULT '100',
  `scheme _url` varchar(255) NOT NULL DEFAULT '' COMMENT '关键词的跳转链接',
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='热闹关键词表';

*/
