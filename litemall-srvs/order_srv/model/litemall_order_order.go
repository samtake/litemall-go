package model

type Order struct {
	BaseModel
	OrderSn        string  `gorm:"type:varchar(63)  NOT NULL DEFAULT '';index:order_sn  COMMENT '订单号，平台自己生成的订单号'"`
	PayType        string  `gorm:"type:varchar(20) comment 'alipay(支付宝)， wechat(微信)'"`
	payTime        int32   `gorm:"type:int(11) unsigned NOT NULL DEFAULT '0'"`
	PayStatus      int32   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '0';index:pay_status comment 'PAYING(待支付), TRADE_SUCCESS(成功)， TRADE_CLOSED(超时关闭), WAIT_BUYER_PAY(交易创建), TRADE_FINISHED(交易结束)'"`
	PayName        string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	PayId          int32   `gorm:"type:tinyint(3) NOT NULL DEFAULT '0';index:pay_id COMMENT '支付宝的订单号 查账'"`
	OrderStatus    int32   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '0';index:order_status"`
	UserId         int32   `gorm:"type:mediumint(8) unsigned NOT NULL DEFAULT '0';index:user_id"`
	ShippingStatus int32   `gorm:"type:tinyint(1) unsigned NOT NULL DEFAULT '0';index:shipping_status"`
	Consignee      string  `gorm:"type:varchar(60) NOT NULL DEFAULT ''"`
	Country        string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	City           string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	District       string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	Address        string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	Mobile         string  `gorm:"type:varchar(60) NOT NULL DEFAULT ''"`
	Postscript     string  `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	ShippingFee    float64 `gorm:"type:decimal(10,2) NOT NULL DEFAULT '0.00'"`
	ActualPrice    float64 `gorm:"type:decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '实际需要支付的金额'"`
	Integral       int32   `gorm:"type:int(10) unsigned NOT NULL DEFAULT '0'"`
	IntegralMoney  float64 `gorm:"type:decimal(10, 2) NOT NULL DEFAULT '0.00'"`
	OrderPrice     float64 `gorm:"type:decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '订单总价'"`
	GoodsPrice     float64 `gorm:"type:decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '商品总价'"`
	AddTime        int32   `gorm:"type:int(11) unsigned NOT NULL DEFAULT '0'"`
	ConfirmTime    int32   `gorm:"type:int(11) unsigned NOT NULL DEFAULT '0'"`
	FreightPrice   int32   `gorm:"type:int(10) unsigned NOT NULL DEFAULT '0' COMMENT '配送费用'"`
	CouponId       int32   `gorm:"type:mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '使用的优惠券id'"`
	ParentId       int32   `gorm:"type:mediumint(8) unsigned NOT NULL DEFAULT '0'"`
	CouponPrice    float32 `gorm:"type:decimal(10, 2) NOT NULL DEFAULT '0.00'"`
	CallbackStatus bool    `gorm:"type:enum('true', 'false') DEFAULT 'true'"`
}

/*
-- ----------------------------
-- Table structure for nideshop_order
-- ----------------------------
DROP TABLE IF EXISTS `nideshop_order`;
CREATE TABLE `nideshop_order` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `order_sn` varchar(20) NOT NULL DEFAULT '',
  `user_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `order_status` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `shipping_status` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `pay_status` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `consignee` varchar(60) NOT NULL DEFAULT '',
  `country` smallint(5) unsigned NOT NULL DEFAULT '0',
  `province` smallint(5) unsigned NOT NULL DEFAULT '0',
  `city` smallint(5) unsigned NOT NULL DEFAULT '0',
  `district` smallint(5) unsigned NOT NULL DEFAULT '0',
  `address` varchar(255) NOT NULL DEFAULT '',
  `mobile` varchar(60) NOT NULL DEFAULT '',
  `postscript` varchar(255) NOT NULL DEFAULT '',
  `shipping_fee` decimal(10,2) NOT NULL DEFAULT '0.00',
  `pay_name` varchar(120) NOT NULL DEFAULT '',
  `pay_id` tinyint(3) NOT NULL DEFAULT '0',
  `actual_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '实际需要支付的金额',
  `integral` int(10) unsigned NOT NULL DEFAULT '0',
  `integral_money` decimal(10,2) NOT NULL DEFAULT '0.00',
  `order_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '订单总价',
  `goods_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品总价',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0',
  `confirm_time` int(11) unsigned NOT NULL DEFAULT '0',
  `pay_time` int(11) unsigned NOT NULL DEFAULT '0',
  `freight_price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '配送费用',
  `coupon_id` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '使用的优惠券id',
  `parent_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `coupon_price` decimal(10,2) NOT NULL DEFAULT '0.00',
  `callback_status` enum('true','false') DEFAULT 'true',
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_sn` (`order_sn`),
  KEY `user_id` (`user_id`),
  KEY `order_status` (`order_status`),
  KEY `shipping_status` (`shipping_status`),
  KEY `pay_status` (`pay_status`),
  KEY `pay_id` (`pay_id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4;
*/
