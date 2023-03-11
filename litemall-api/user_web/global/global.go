package global

import (
	ut "github.com/go-playground/universal-translator"
	"litemall-api/user_web/config"
	"litemall-api/user_web/proto"
)

var (
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	Trans         ut.Translator
	UserSrvClient proto.UserClient
	NacosConfig   *config.NacosConfig = &config.NacosConfig{}
)
