package main

import (
	"context"
	"github.com/micro/go-micro/v2/logger"

	postProto "micro-service/proto/post"
	userProto "micro-service/proto/user"

	micro "github.com/micro/go-micro/v2"
)

// 使用方式：将下面代码插入到web，api项目中来调用user.QueryUserByName服务。可以参考web/gin.go。
func main() {
	// Create a new service
	service := micro.NewService(
		micro.Name("user.client"),
		micro.Version("v1.0.0"),
		//micro.Registry(etcd.NewRegistry(
		//	// 地址是我本地etcd服务器地址，不要照抄
		//	registry.Addrs("127.0.0.1:2379"),
		//)),
	)
	// Initialise the client and parse command line flags
	service.Init()
	client := service.Client()
	// Create new user client
	user := userProto.NewUserService("go.micro.srv.user", client)
	post := postProto.NewPostService("go.micro.srv.user", client)

	// Call the user
	rsp, err := user.QueryUserByName(context.TODO(), &userProto.Request{UserName: "John"})
	if err != nil {
		logger.Fatal(err)
	}

	// Print response
	logger.Info(rsp.GetUser())

	rsp2, err2 := post.QueryUserPosts(context.TODO(), &postProto.Request{UserID: 1})
	if err2 != nil {
		logger.Fatal(err2)
	}
	// Print response
	logger.Info(rsp2.GetPost().Title)
}
