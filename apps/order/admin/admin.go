package main

import (
	"flag"
	"fmt"

	"github.com/zhoushuguang/lebron/apps/order/admin/internal/config"
	"github.com/zhoushuguang/lebron/apps/order/admin/internal/handler"
	"github.com/zhoushuguang/lebron/apps/order/admin/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the etc file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
