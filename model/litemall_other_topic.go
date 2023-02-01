package model

import "github.com/jinzhu/gorm"

//主题

type topic struct {
	gorm.Model
	Title           string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''''''"`
	Content         string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''''''"`
	Avatar          string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''''''"`
	ItemPicUrl      string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''''''"`
	SubTitle        string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''''''"`
	TopicCategoryId uint    `gorm:"type:smallint(11) unsigned NOT NULL DEFAULT '0'"`
	PriceInfo       float64 `gorm:"type:decimal(10,2) unsigned NOT NULL DEFAULT '0.00'"`
	ReadCount       string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''''''"`
	ScenePicUrl     string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''''''"`
	TopicTemplateId uint    `gorm:"type:int(11) unsigned NOT NULL DEFAULT '0'"`
	TopicTagId      uint    `gorm:"type:int(11) unsigned NOT NULL DEFAULT '0'"`
	SortOrder       uint    `gorm:"type:int(11) unsigned NOT NULL DEFAULT '100'"`
	IsShow          uint    `gorm:"type:int(1) unsigned NOT NULL DEFAULT '1'"`
}

/*
-- ----------------------------
-- Table structure for nideshop_topic
-- ----------------------------
DROP TABLE IF EXISTS `nideshop_topic`;
CREATE TABLE `nideshop_topic` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '''''',
  `content` text,
  `avatar` varchar(255) NOT NULL DEFAULT '',
  `item_pic_url` varchar(255) NOT NULL DEFAULT '',
  `subtitle` varchar(255) NOT NULL DEFAULT '''',
  `topic_category_id` int(11) unsigned NOT NULL DEFAULT '0',
  `price_info` decimal(10,2) unsigned NOT NULL DEFAULT '0.00',
  `read_count` varchar(255) NOT NULL DEFAULT '0',
  `scene_pic_url` varchar(255) NOT NULL DEFAULT '',
  `topic_template_id` int(11) unsigned NOT NULL DEFAULT '0',
  `topic_tag_id` int(11) unsigned NOT NULL DEFAULT '0',
  `sort_order` int(11) unsigned NOT NULL DEFAULT '100',
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '1',
  KEY `topic_id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=316 DEFAULT CHARSET=utf8mb4;
*/
