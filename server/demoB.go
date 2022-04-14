package server

import (
	"context"
	"github.com/Gouplook/rpcxB/logics"
)

type DemoA struct {
}

func (d *DemoA) Add(ctx context.Context, a int, b int, rs *int) error {

	logic := new(logics.DemoALogic)

	err := logic.Add(ctx, a, b, rs)

	if err != nil {
		panic(err)
	}
	return nil
}
