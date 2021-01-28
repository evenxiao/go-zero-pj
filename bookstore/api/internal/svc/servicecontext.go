package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"gozero/bookstore/api/internal/config"
	"gozero/bookstore/rpc/add/adder"
	"gozero/bookstore/rpc/check/checker"
)

type ServiceContext struct {
	Config config.Config
	Adder   adder.Adder          // 手动代码
	Checker checker.Checker      // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),         // 手动代码
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),   // 手动代码
	}
}
