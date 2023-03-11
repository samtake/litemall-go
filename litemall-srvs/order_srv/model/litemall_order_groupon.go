package model

type Groupon struct {
	BaseModel
	OrderId         uint   `gorm:"type:mediumint(11) unsigned NOT NULL DEFAULT '0';index:order_id"`
	GrouponId       uint   `gorm:"type:mediumint(11) unsigned NOT NULL DEFAULT '0';index:groupon_id"`
	RulesId         uint   `gorm:"type:mediumint(11) unsigned NOT NULL DEFAULT '0';index:rules_id"`
	UserId          uint   `gorm:"type:mediumint(11) unsigned NOT NULL DEFAULT '0';index:user_id"`
	ShareUrl        string `gorm:"type:varchar(255) NOT NULL DEFAULT ''"`
	CreatorUserId   uint   `gorm:"type:mediumint(11) unsigned NOT NULL DEFAULT '0';index:creator_user_id"`
	CreatorUserTime string `gorm:"type:type:varchar(255) NOT NULL DEFAULT ''"`
}
