package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"gozero/bookstore/api/internal/config"
	"gozero/bookstore/api/internal/handler"
	"gozero/bookstore/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/bookstore-api.yaml", "the config file")

func main() {
	flag.Parse()
	logC := logx.LogConf{
		Mode: "file",
		Path: "logs",
	}
	logx.Info("22344444")
	logx.MustSetup(logC)
	logx.Error("1111111111")
	logx.Info("loginfo11111111111")
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
