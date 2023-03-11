package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

//用户表

type User struct {
	gorm.Model
	Username      string     `gorm:"type:varchar(60); NOT NULL DEFAULT ''"`
	Password      string     `gorm:"type:varchar(100); NOT NULL DEFAULT ''"`
	Gender        int32      `gorm:"type:tinyint(3) unsigned; NOT NULL DEFAULT '0'"`
	Role          int32      `gorm:"type:tinyint(3) unsigned; NOT NULL DEFAULT '0'"`
	Telephone     string     `gorm:"type:varchar(20)"`
	Birthday      *time.Time `gorm:"type:datetime"`
	LastLoginTime *time.Time `gorm:"type:datetime"`
	LastLoginIp   string     `gorm:"type:varchar(255); NOT NULL DEFAULT ''"`
	UserLevel     uint       `gorm:"type:int(11) unsigned; NOT NULL DEFAULT '0'"`
	NickName      string     `gorm:"type:varchar(63);NOT NULL DEFAULT '0'"`
	Mobile        string     `gorm:"type:varchar(20); NOT NULL DEFAULT '0'"`
	Avatar        string     `gorm:"type:varchar(255);not null DEFAULT ''"`
	WeixinOpenid  string     `gorm:"type:varchar(63); NOT NULL DEFAULT '0'"`
	SessionKey    string     `gorm:"type:varchar(100); NOT NULL DEFAULT '0'"`
	Status        uint       `gorm:"type:tinyint(3) unsigned; NOT NULL DEFAULT '0'"`
	Deleted       uint       `gorm:"type:tinyint(1) unsigned; NOT NULL DEFAULT '0'"`
}

/*

// gorm.Model 的定义
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

DROP TABLE IF EXISTS `nideshop_user`;
CREATE TABLE `nideshop_user` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(60) NOT NULL DEFAULT '',
  `password` varchar(32) NOT NULL DEFAULT '',
  `gender` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `birthday` int(11) unsigned NOT NULL DEFAULT '0',
  `register_time` int(11) unsigned NOT NULL DEFAULT '0',
  `last_login_time` int(11) unsigned NOT NULL DEFAULT '0',
  `last_login_ip` varchar(255) NOT NULL DEFAULT '',
  `user_level_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `nickname` varchar(60) NOT NULL,
  `mobile` varchar(20) NOT NULL,
  `register_ip` varchar(255) NOT NULL DEFAULT '',
  `avatar` varchar(255) NOT NULL DEFAULT '',
  `weixin_openid` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
*/
