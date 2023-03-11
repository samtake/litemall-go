package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"litemall-srvs/user_srv/global"
	"litemall-srvs/user_srv/handler"
	initialize2 "litemall-srvs/user_srv/initialize"
	pb "litemall-srvs/user_srv/proto"
	"net"
)

func main() {
	IP := flag.String("ip", "127.0.0.1", "ip地址")
	Port := flag.Int("port", 50051, "端口号")

	initialize2.InitLogger()
	initialize2.InitConfig()
	initialize2.InitDB()
	flag.Parse()
	zap.S().Info("ip:", *IP)
	zap.S().Info("Port:", *Port)
	server := grpc.NewServer()
	pb.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic(any("failed to listen:" + err.Error()))
	}

	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//服务注册
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host,
		consulInfo.Port)
	zap.S().Infof("consulInfo.Host: %s", consulInfo.Host)
	zap.S().Infof("consulInfo.Port: %d", consulInfo.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(any(err))
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", *IP, *Port),
		Timeout:                        "10s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
		//GRPCUseTLS:                     false,
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serviceID := fmt.Sprintf("%s", uuid.NewV4())
	registration.ID = serviceID
	registration.Port = *Port
	registration.Tags = []string{"user-srv22", "sam", "user", "srv"}
	registration.Address = "127.0.0.1"
	registration.Check = check
	//1. 如何启动两个服务
	//2. 即使我能够通过终端启动两个服务，但是注册到consul中的时候也会被覆盖
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(any(err))
	}

	err = server.Serve(lis)
	if err != nil {
		panic(any("failed to start grpc" + err.Error()))
	}
}
