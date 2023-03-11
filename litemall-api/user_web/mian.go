package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"litemall-api/user_web/global"
	initialize2 "litemall-api/user_web/initialize"
	myValidator "litemall-api/user_web/validator"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	/*
		init:
		1.logger
		2.config
		3.router
		4.translate
		5.srvconn
	*/
	initialize2.InitLogger()
	initialize2.InitConfig()
	Router := initialize2.Routers()
	if err := initialize2.InitTrans("zh"); err != nil {
		panic(any(err))
	}

	initialize2.InitSrvConn()

	//注册验证器
	//validator:2
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myValidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	//port := 8021
	zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
