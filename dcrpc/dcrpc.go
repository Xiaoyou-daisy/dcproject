package main

import (
	"dcproject/dcrpc/basic/inits"
	"dcproject/dcrpc/dcrpc"
	"dcproject/dcrpc/internal/config"
	"dcproject/dcrpc/internal/server"
	"dcproject/dcrpc/internal/svc"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/dcrpc.yaml", "the config file")

func main() {

	inits.MongeDB()
	inits.MysqlInit()
	inits.ExampleClient()
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		dcrpc.RegisterDcrpcServer(grpcServer, server.NewDcrpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
