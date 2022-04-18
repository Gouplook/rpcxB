package routers

import (
	"fmt"
	"github.com/Gouplook/rpcxB/service"
	"github.com/smallnest/rpcx/server"
)

//搜索注册路由
func InitRpcRouters(rpcServer *server.Server) {

	//err := rpcServer.RegisterName("Demo/Add", new(service.DemoB), "")
	err := rpcServer.Register(new(service.DemoB), "")
	if err != nil {
		fmt.Println("failed to register rpcRouter: %v", err)
	}
}
