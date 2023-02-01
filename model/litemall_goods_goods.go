package model

import "github.com/jinzhu/gorm"

type Goods struct {
	gorm.Model
	CategoryId        uint    `gorm:"type:tinyint(11) unsigned NOT NULL DEFAULT '1';index:cat_id"`
	GoodsSn           string  `gorm:"type:varchar(60) NOT NULL DEFAULT '';index:goods_sn"`
	Name              string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	BrandId           uint    `gorm:"type:tinyint(11) unsigned NOT NULL DEFAULT '1';index:brand_id"`
	GoodsNumber       uint    `gorm:"type:tinyint(8) unsigned NOT NULL DEFAULT '0';index:goods_number"`
	Keywords          string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	GoodsBrief        string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	GoodsDesc         string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	isOnSale          uint    `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '1'"`
	SortOrder         uint    `gorm:"type:tinyint(4) unsigned NOT NULL DEFAULT '100';index:sort_order"`
	IsDelete          uint    `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '0'"`
	AttributeCategory uint    `gorm:"type:tinyint(11) unsigned NOT NULL DEFAULT '0'"`
	CounterPrice      float64 `gorm:"type:decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '专柜价格'"`
	ExtraPrice        float64 `gorm:"type:decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '附加价格'"`
	IsNew             uint    `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '0'"`
	GoodsUnit         string  `gorm:"type:varchar(45) NOT NULL COMMENT '商品单位'"`
	PrimaryPicUrl     string  `gorm:"type:varchar(255) NOT NULL COMMENT '商品主图'"`
	ListPicUrl        string  `gorm:"type:varchar(255) NOT NULL COMMENT '商品列表图'"`
	RetailPrice       float64 `gorm:"type:decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '零售价格'"`
	SellVolume        uint    `gorm:"int(11) unsigned NOT NULL DEFAULT '0' COMMENT '销售量'"`
	PrimaryProductId  uint    `gorm:"type:int(11) unsigned NOT NULL DEFAULT '0' COMMENT '主sku　product_id'"`
	UnitPrice         float64 `gorm:"type:decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '单位价格，单价'"`
	PromotionDesc     string  `gorm:"type:varchar(255) NOT NULL"`
	PromotionTag      string  `gorm:"type:varchar(45) NOT NULL"`
	AppExclusivePrice float64 `gorm:"type:decimal(10, 2) unsigned NOT NULL COMMENT 'APP专享价'"`
	IsAppExclusive    uint    `gorm:"type:tinyint(1) unsigned NOT NULL COMMENT '是否是APP专属'"`
	IsLimited         uint    `gorm:"type:tinyint(1) unsigned NOT NULL"`
	IsHot             uint    `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '0'"`
}

/*
-- ----------------------------
-- Table structure for nideshop_goods
-- ----------------------------
DROP TABLE IF EXISTS `nideshop_goods`;
CREATE TABLE `nideshop_goods` (
  `id` int(11) unsigned NOT NULL,
  `category_id` int(11) unsigned NOT NULL DEFAULT '0',
  `goods_sn` varchar(60) NOT NULL DEFAULT '',
  `name` varchar(120) NOT NULL DEFAULT '',
  `brand_id` int(11) unsigned NOT NULL DEFAULT '0',
  `goods_number` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `keywords` varchar(255) NOT NULL DEFAULT '',
  `goods_brief` varchar(255) NOT NULL DEFAULT '',
  `goods_desc` text,
  `is_on_sale` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0',
  `sort_order` smallint(4) unsigned NOT NULL DEFAULT '100',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `attribute_category` int(11) unsigned NOT NULL DEFAULT '0',
  `counter_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '专柜价格',
  `extra_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '附加价格',
  `is_new` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `goods_unit` varchar(45) NOT NULL COMMENT '商品单位',
  `primary_pic_url` varchar(255) NOT NULL COMMENT '商品主图',
  `list_pic_url` varchar(255) NOT NULL COMMENT '商品列表图',
  `retail_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '零售价格',
  `sell_volume` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '销售量',
  `primary_product_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '主sku　product_id',
  `unit_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '单位价格，单价',
  `promotion_desc` varchar(255) NOT NULL,
  `promotion_tag` varchar(45) NOT NULL,
  `app_exclusive_price` decimal(10,2) unsigned NOT NULL COMMENT 'APP专享价',
  `is_app_exclusive` tinyint(1) unsigned NOT NULL COMMENT '是否是APP专属',
  `is_limited` tinyint(1) unsigned NOT NULL,
  `is_hot` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `goods_sn` (`goods_sn`),
  KEY `cat_id` (`category_id`),
  KEY `brand_id` (`brand_id`),
  KEY `goods_number` (`goods_number`),
  KEY `sort_order` (`sort_order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of nideshop_goods
-- ----------------------------
*/
