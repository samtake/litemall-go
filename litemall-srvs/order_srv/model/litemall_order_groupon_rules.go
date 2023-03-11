package model

type GrouponRules struct {
	BaseModel
	GoodsId        uint   `gorm:"type:mediumint(11) unsigned NOT NULL DEFAULT '0';index:goods_id"`
	GoodsName      string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	PicUrl         string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	DiscountMember uint   `gorm:"type:mediumint(11) unsigned NOT NULL DEFAULT '0'"`
	ExpireTime     string `gorm:"type:type:varchar(25) NOT NULL DEFAULT ''"`
	Status         uint   `gorm:"type:smallint(6) unsigned NOT NULL DEFAULT '0';index:groupon_id"`
	RulesId        uint   `gorm:"type:mediumint(11) unsigned NOT NULL DEFAULT '0';index:rules_id"`
	UserId         uint   `gorm:"type:mediumint(11) unsigned NOT NULL DEFAULT '0';index:user_id"`
	CreatorUserId  uint   `gorm:"type:mediumint(11) unsigned NOT NULL DEFAULT '0';index:creator_user_id"`
}
