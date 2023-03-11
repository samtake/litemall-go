package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	proto2 "litemall-srvs/order_srv/proto"
)

var userClient proto2.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(any(err))
	}
	userClient = proto2.NewUserClient(conn)
}

func TestCreateUser() {
	for i := 5; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto2.CreateUserInfo{
			NickName: fmt.Sprintf("ss%d", i),
			Mobile:   fmt.Sprintf("15504705%d", i),
			PassWord: "admin123",
		})
		if err != nil {
			panic(any(err))
		}
		fmt.Println(rsp.Id)
	}
}

func main() {
	Init()
	TestCreateUser()
	//TestGetUserList()

	conn.Close()
}
