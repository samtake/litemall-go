package model

import "github.com/jinzhu/gorm"

//客服、自助表达

type Issue struct {
	gorm.Model
	Question string `gorm:"varchar(255) NOT NULL DEFAULT ''"`
	Answer   string `gorm:"varchar(255) NOT NULL DEFAULT ''"`
}

/*
-- ----------------------------
-- Table structure for nideshop_goods_issue
-- ----------------------------
DROP TABLE IF EXISTS `nideshop_goods_issue`;
CREATE TABLE `nideshop_goods_issue` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `goods_id` text,
  `question` varchar(255) DEFAULT NULL,
  `answer` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
*/
