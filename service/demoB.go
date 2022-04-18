package service

import (
	"context"
	"github.com/Gouplook/rpcxB/logics"
)

type DemoB struct {
}

func (d *DemoB) Add(ctx context.Context, a *int, b *int, rs *int) error {
	logic := new(logics.DemoALogic)
	err := logic.Add(ctx, a, b, rs)
	if err != nil {
		panic(err)
	}
	return nil
}
