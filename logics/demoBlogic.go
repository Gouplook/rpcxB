package logics

import "context"

type DemoALogic struct {
}

func (d *DemoALogic) Add(ctx context.Context, a, b int, rs *int) error {
	// 逻辑代码，数据库读写
	//......
	// 调用rpcxA 中的方法

	return nil
}
