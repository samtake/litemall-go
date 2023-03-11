package model

//商品属性表

type GoodsAttribute struct {
	BaseModel
	GoodsId     int32  `gorm:"type:tinyint(11) unsigned NOT NULL DEFAULT '1';index:goods_id"`
	AttributeId int32  `gorm:"type:tinyint(11) unsigned NOT NULL DEFAULT '1';index:attr_id"`
	Value       string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
}

/*
-- ----------------------------
-- Table structure for nideshop_goods_attribute
-- ----------------------------
DROP TABLE IF EXISTS `nideshop_goods_attribute`;
CREATE TABLE `nideshop_goods_attribute` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `goods_id` int(11) unsigned NOT NULL DEFAULT '0',
  `attribute_id` int(11) unsigned NOT NULL DEFAULT '0',
  `value` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `goods_id` (`goods_id`),
  KEY `attr_id` (`attribute_id`)
) ENGINE=InnoDB AUTO_INCREMENT=872 DEFAULT CHARSET=utf8mb4;
*/
