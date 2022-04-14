package routers

import (
	"github.com/smallnest/rpcx/server"
)

//搜索注册路由
func InitRpcRouters(rpcServer *server.Server) {

	// API段访问需要注册路由
	//err := rpcServer.RegisterName("Demo/Add", new(server2.DemoA), "")
	//if err != nil {
	//	fmt.Println("failed to register rpcRouter: %v", err)
	//}
}
