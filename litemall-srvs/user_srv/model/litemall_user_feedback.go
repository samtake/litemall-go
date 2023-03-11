package model

import "github.com/jinzhu/gorm"

//反馈

type Feedback struct {
	gorm.Model
	ValueId    uint   `gorm:"type:tinyint(1) unsigned; NOT NULL DEFAULT '0';index"`
	Content    string `gorm:"type:varchar(6550); NOT NULL DEFAULT ''"`
	Status     uint   `gorm:"type:tinyint(3) unsigned; NOT NULL DEFAULT '0'"`
	UserId     uint   `gorm:"type:tinyint(11) unsigned; NOT NULL DEFAULT '0'"`
	NewContent string `gorm:"type:varchar(6550); NOT NULL DEFAULT ''"`
}

/*
CREATE TABLE `nideshop_feedback` (
  `msg_id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `user_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `user_name` varchar(60) NOT NULL DEFAULT '',
  `user_email` varchar(60) NOT NULL DEFAULT '',
  `msg_title` varchar(200) NOT NULL DEFAULT '',
  `msg_type` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `msg_status` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `msg_content` text NOT NULL,
  `msg_time` int(10) unsigned NOT NULL DEFAULT '0',
  `message_img` varchar(255) NOT NULL DEFAULT '0',
  `order_id` int(11) unsigned NOT NULL DEFAULT '0',
  `msg_area` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`msg_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
*/
