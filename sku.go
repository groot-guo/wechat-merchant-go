package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"wechat-merchant-go/internal/config"
	itemservice "wechat-merchant-go/internal/server/itemservice"
	shopservice "wechat-merchant-go/internal/server/shopservice"
	skuinventoryservice "wechat-merchant-go/internal/server/skuinventoryservice"
	skuservice "wechat-merchant-go/internal/server/skuservice"
	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"
)

var configFile = flag.String("f", "etc/sku.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(&c)

	logc.MustSetup(c.Log)

	// RegisterService
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		sku.RegisterSkuServiceServer(grpcServer, skuservice.NewSkuServiceServer(ctx))
		sku.RegisterItemServiceServer(grpcServer, itemservice.NewItemServiceServer(ctx))
		sku.RegisterShopServiceServer(grpcServer, shopservice.NewShopServiceServer(ctx))
		sku.RegisterSkuInventoryServiceServer(grpcServer, skuinventoryservice.NewSkuInventoryServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
