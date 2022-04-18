package main

import (
	"context"
	"fmt"
	"github.com/Gouplook/rpcxB/logics"
	"testing"
)

func TestAdd(t *testing.T) {
	add := new(logics.DemoALogic)
	a := 12
	b := 13
	var rs int
	err := add.Add(context.Background(), &a, &b, &rs)
	if err != nil {
		panic(err)
	}

	fmt.Println(rs)
}
