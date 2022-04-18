package logics

import (
	"context"
	"fmt"
	"github.com/Gouplook/rpcxinterfaxe/client/demoA"
	demoA2 "github.com/Gouplook/rpcxinterfaxe/interface/demoA"
)

type DemoALogic struct {
}

func (d *DemoALogic) Add(ctx context.Context, a, b *int, rs *int) error {
	rpcA := new(demoA.DemoA).Init()

	args := demoA2.ArgsAdd{
		Num_1: *a,
		Num_2: *b,
	}
	var reply demoA2.ReplyAdd

	err := rpcA.Add(ctx, &args, &reply)
	if err != nil {
		panic(err)
	}
	*rs = reply.Sum
	fmt.Println(reply.Sum)
	//*rs = 10

	return nil
}
