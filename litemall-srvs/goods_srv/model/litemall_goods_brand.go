package model

type Brand struct {
	BaseModel
	Name       string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	Desc       string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	PicUrl     string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	SortOrder  int32   `gorm:"type:tinyint(3) unsigned NOT NULL DEFAULT '0'"`
	FloorPrice float32 `gorm:"type:decimal(10, 2) NOT NULL DEFAULT '0.00'"`
	IsShow     int32   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '1'"`
}

/*
-- ----------------------------
-- Table structure for nideshop_brand
-- ----------------------------
DROP TABLE IF EXISTS `nideshop_brand`;
CREATE TABLE `nideshop_brand` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `list_pic_url` varchar(255) NOT NULL DEFAULT '',
  `simple_desc` varchar(255) NOT NULL DEFAULT '',
  `pic_url` varchar(255) NOT NULL DEFAULT '',
  `sort_order` tinyint(3) unsigned NOT NULL DEFAULT '50',
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `floor_price` decimal(10,2) NOT NULL DEFAULT '0.00',
  `app_list_pic_url` varchar(255) NOT NULL DEFAULT '',
  `is_new` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `new_pic_url` varchar(255) NOT NULL DEFAULT '',
  `new_sort_order` tinyint(2) unsigned NOT NULL DEFAULT '10',
  PRIMARY KEY (`id`),
  KEY `is_show` (`is_show`)
) ENGINE=InnoDB AUTO_INCREMENT=1046012 DEFAULT CHARSET=utf8mb4;
*/
