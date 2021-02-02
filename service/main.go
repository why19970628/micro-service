package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"micro-service/service/handler"
	//"github.com/micro/go-micro/v2/client/selector"
	//"github.com/micro/go-micro/v2/registry"
	//"github.com/micro/go-micro/v2/registry/etcd"

	post "micro-service/proto/post"
	user "micro-service/proto/user"
)

func main() {
	//reg := etcd.NewRegistry(func(options *registry.Options) {
	//	options.Addrs = []string{
	//		"xx.xx.xx.xx:2379",
	//	}
	//})
	//micro.Selector(selector.NewSelector(func(options *selector.Options) {
	//	options.Registry=reg
	//}))




	// New Service
	service := micro.NewService(
		// micro.Registry(reg),
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		// 配置etcd为注册中心，配置etcd路径，默认端口是2379
		//micro.Registry(etcd.NewRegistry(
		//	registry.Addrs("127.0.0.1:2379"),
		//)),

		// 注册服务生命周期,过了这么长时间需要重新注册
		//micro.RegisterTTL(time.Second*30),

		// 多长时间注册一次
		//micro.RegisterInterval(time.Second*15),

	)

	// Initialise service
	service.Init()
	server := service.Server()
	// Register Handler
	_ = user.RegisterUserHandler(server, new(handler.User))
	_ = post.RegisterPostHandler(server, new(handler.Post))

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
