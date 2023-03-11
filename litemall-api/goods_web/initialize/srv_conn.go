package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"litemall-api/goods_web/global"
	"litemall-api/goods_web/proto"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		//grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	zap.S().Infof("[InitSrvConn] 连接:consulInfo.Host:%s,consulInfo.Port:%d", consulInfo.Host, consulInfo.Port)
	global.GoodsSrvClient = proto.NewGoodsClient(userConn)
}

//初始化连接器：1、从consul获取用户服务信息 2、连接grpc

func InitSrvConn_old() {
	/***********************从consul获取用户服务信息*************************/
	zap.S().Infof("从consul获取用户服务信息:")
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	zap.S().Infof("consulInfo.Host:%s", consulInfo.Host)
	zap.S().Infof("consulInfo.Port:%d", consulInfo.Port)
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(any(err))
	}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.GoodsSrvInfo.Name))
	if err != nil {
		panic(any(err))
	}
	for _, value := range data {
		zap.S().Infof("user info value from consul :%v", value)
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}

	if userSrvHost == "" {
		//ctx.JSON(http.StatusBadRequest, gin.H{
		//	"msg": "user-srv不能获取",
		//})
		zap.S().Fatalln("goods-srv不能获取")
		return
	}
	/***********************从consul获取用户服务信息*************************/

	/***********************连接grpc*************************/
	zap.S().Infof("连接grpc")
	//userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList]连接失败：", err.Error())
	}
	global.GoodsSrvClient = proto.NewGoodsClient(userConn)
	//claims, _ := ctx.Get("claims")
	//currentClaims := claims.(*models.CustomClaims)
	//zap.S().Infof("访问用户：%d", currentClaims.ID)
	//userSrvClient := proto.NewUserClient(userConn)
	/***********************连接grpc*************************/
}
