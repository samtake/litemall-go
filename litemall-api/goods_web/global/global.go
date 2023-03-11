package global

import (
	ut "github.com/go-playground/universal-translator"
	"litemall-api/goods_web/config"
	"litemall-api/goods_web/proto"
)

var (
	ServerConfig   *config.ServerConfig = &config.ServerConfig{}
	Trans          ut.Translator
	GoodsSrvClient proto.GoodsClient
	NacosConfig    *config.NacosConfig = &config.NacosConfig{}
)
