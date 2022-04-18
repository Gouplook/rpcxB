package main

import (
	"fmt"
	"github.com/Gouplook/dzgin"
	"github.com/Gouplook/rpcxB/routers"
	"github.com/smallnest/rpcx/server"
)

func main() {

	// 初始化配置
	fmt.Println(dzgin.AppConfig.String("rpchost"))

	//启动服务
	rpcServer := server.NewServer()
	//服务注册
	routers.InitRpcRouters(rpcServer)

	//启动链路追踪

	//address := "0.0.0.0:9002"
	address := fmt.Sprintf("%v:%v", dzgin.AppConfig.String("rpchost"), dzgin.AppConfig.String("rpcport"))
	if err := rpcServer.Serve("tcp", address); err != nil {
		// 打印日志
		fmt.Println("rpcServer fail")
	}

}
