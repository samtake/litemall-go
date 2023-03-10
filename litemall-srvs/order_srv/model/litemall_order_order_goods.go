package model

//记录订单商品信息

type OrderGoods struct {
	BaseModel
	OrderId                   int32   `gorm:"type:mediumint(8) unsigned NOT NULL DEFAULT '0';index:order_id"`
	GoodsId                   int32   `gorm:"type:mediumint(8) unsigned NOT NULL DEFAULT '0';index:goods_id"`
	GoodsName                 string  `gorm:"type:varchar(60) NOT NULL DEFAULT ''"`
	GoodsSn                   string  `gorm:"type:varchar(60) NOT NULL DEFAULT ''"`
	ProductId                 int32   `gorm:"type:mediumint(8) unsigned NOT NULL DEFAULT '0'"`
	Number                    int32   `gorm:"type:mediumint(8) unsigned NOT NULL DEFAULT '1'"`
	MarketPrice               float64 `gorm:"type:decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '市场价'"`
	RetailPrice               float64 `gorm:"type:decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '零售价格'"`
	GoodsSpecifitionNameValue string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	IsReal                    int32   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '0'"`
	GoodsSpecifitionIds       string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	ListPicUrl                string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
}

/*
-- ----------------------------
-- Table structure for nideshop_order_goods
-- ----------------------------
DROP TABLE IF EXISTS `nideshop_order_goods`;
CREATE TABLE `nideshop_order_goods` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `goods_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `goods_name` varchar(120) NOT NULL DEFAULT '',
  `goods_sn` varchar(60) NOT NULL DEFAULT '',
  `product_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `number` smallint(5) unsigned NOT NULL DEFAULT '1',
  `market_price` decimal(10,2) NOT NULL DEFAULT '0.00',
  `retail_price` decimal(10,2) NOT NULL DEFAULT '0.00',
  `goods_specifition_name_value` text NOT NULL,
  `is_real` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `goods_specifition_ids` varchar(255) NOT NULL DEFAULT '',
  `list_pic_url` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`),
  KEY `goods_id` (`goods_id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4;
*/
