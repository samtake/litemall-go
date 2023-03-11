package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"litemall-api/goods_web/global"
	"litemall-api/goods_web/initialize"
	"litemall-api/goods_web/utils/register/consul"
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
	initialize.InitLogger()
	initialize.InitConfig()
	Router := initialize.Routers()
	if err := initialize.InitTrans("zh"); err != nil {
		panic(any(err))
	}

	initialize.InitSrvConn()

	//目前只有本地开发环境
	//viper.AutomaticEnv()
	////如果是本地开发环境端口号固定，线上环境启动获取端口号
	//debug := viper.GetBool("MXSHOP_DEBUG")
	////debug = false
	//if !debug{
	//	port, err := utils.GetFreePort()
	//	if err == nil {
	//		global.ServerConfig.Port = port
	//	}
	//}

	//统一在nacos上拿
	register_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := register_client.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}

	//port := goods_srv
	zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}

	//接收终止信号
	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = register_client.DeRegister(serviceId); err != nil {
		zap.S().Info("注销失败:", err.Error())
	} else {
		zap.S().Info("注销成功:")
	}
}
