package model

//商品表

type Goods struct {
	BaseModel

	BrandId int32 `gorm:"type:int;not null"`
	Brand   Brand

	CategoryId int32 `gorm:"type:int;not null"`
	Category   Category

	GoodsAttributeId int32 `gorm:"type:int;not null"`
	GoodsAttribute   GoodsAttribute

	GoodsProductId int32 `gorm:"type:int;not null"`
	GoodsProduct   GoodsProduct

	GoodsSpecificationId int32 `gorm:"type:int;not null"`
	GoodsSpecification   GoodsSpecification

	GoodsSn           string  `gorm:"type:varchar(60) NOT NULL DEFAULT ''"`
	Name              string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	GoodsNumber       int32   `gorm:"type:tinyint(8) unsigned NOT NULL DEFAULT '0'"`
	Keywords          string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	GoodsBrief        string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	GoodsDesc         string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	isOnSale          int32   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '1'"`
	SortOrder         int32   `gorm:"type:tinyint(4) unsigned NOT NULL DEFAULT '100'"`
	IsDelete          int32   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '0'"`
	AttributeCategory int32   `gorm:"type:tinyint(11) unsigned NOT NULL DEFAULT '0'"`
	CounterPrice      float64 `gorm:"type:decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '专柜价格'"`
	ExtraPrice        float64 `gorm:"type:decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '附加价格'"`
	IsNew             int32   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '0'"`
	GoodsUnit         string  `gorm:"type:varchar(45) NOT NULL COMMENT '商品单位'"`

	RetailPrice       float64 `gorm:"type:decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '零售价格'"`
	SellVolume        int32   `gorm:"int(11) unsigned NOT NULL DEFAULT '0' COMMENT '销售量'"`
	PrimaryProductId  int32   `gorm:"type:int(11) unsigned NOT NULL DEFAULT '0' COMMENT '主sku　product_id'"`
	UnitPrice         float64 `gorm:"type:decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '单位价格，单价'"`
	PromotionDesc     string  `gorm:"type:varchar(255) NOT NULL"`
	PromotionTag      string  `gorm:"type:varchar(45) NOT NULL"`
	AppExclusivePrice float64 `gorm:"type:decimal(10, 2) unsigned NOT NULL COMMENT 'APP专享价'"`
	IsAppExclusive    int32   `gorm:"type:tinyint(1) unsigned NOT NULL COMMENT '是否是APP专属'"`
	IsLimited         int32   `gorm:"type:tinyint(1) unsigned NOT NULL"`
	IsHot             int32   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '0'"`

	//PrimaryPicUrl string `gorm:"type:varchar(255) NOT NULL COMMENT '商品主

	Images          GormList `gorm:"type:varchar(1000);not null COMMENT '商品主图'"`
	DescImages      GormList `gorm:"type:varchar(1000);not null COMMENT '商品列表图'"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null COMMENT '封面图'"`
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
