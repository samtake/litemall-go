package model

import "github.com/jinzhu/gorm"

//搜索记录

type SearchHistory struct {
	gorm.Model
	UserId  string `gorm:"type:varchar(45); NOT NULL DEFAULT '0'"`
	Keyword string `gorm:"type:varchar(50); NOT NULL DEFAULT ''"`
	From    string `gorm:"type:varchar(45); NOT NULL DEFAULT ''; COMMENT '搜索来源，如PC、小程序、APP等'"`
	Deleted uint   `gorm:"type:tinyint(1) unsigned; NOT NULL DEFAULT '0'"`
}

/*
DROP TABLE IF EXISTS `nideshop_search_history`;
CREATE TABLE `nideshop_search_history` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `keyword` char(50) NOT NULL,
  `from` varchar(45) NOT NULL DEFAULT '' COMMENT '搜索来源，如PC、小程序、APP等',
  `add_time` int(11) NOT NULL DEFAULT '0' COMMENT '搜索时间',
  `user_id` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4;
*/
