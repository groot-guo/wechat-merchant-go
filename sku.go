package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"log"
	"os"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"wechat-merchant-go/internal/config"
	"wechat-merchant-go/internal/server"
	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"
)

var configFile = flag.String("f", "etc/sku.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	// 创建 Nacos 客户端
	sc := []constant.ServerConfig{
		{
			IpAddr:   c.Nacos.Host,
			GrpcPort: c.Nacos.Port,
		},
	}
	file, err := os.ReadFile("./naming/c2b3cfff-e8bd-4066-ae52-648a8eb6ceeb")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))
	cc := constant.ClientConfig{
		NamespaceId: c.Nacos.Namespace,
		CacheDir:    ".",
		LogLevel:    "info",
		//NotLoadCacheAtStart: false,
		// 其他配置...
	}
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	//configInfo, err := clients.NewConfigClient(vo.NacosClientParam{
	//	ClientConfig:  &cc,
	//	ServerConfigs: sc,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//configInfo.ListenConfig(vo.ConfigParam{
	//	DataId:           "",
	//	Group:            "",
	//	Content:          "",
	//	Tag:              "",
	//	AppName:          "",
	//	BetaIps:          "",
	//	CasMd5:           "",
	//	Type:             "",
	//	SrcUser:          "",
	//	EncryptedDataKey: "",
	//	KmsKeyId:         "",
	//	UsageType:        "",
	//	OnChange:         nil,
	//})

	// 服务注册
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        8090,
		ServiceName: "wechat-merchant-go",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"project": "go-zero"},
		//ClusterName: "CLUSTER_NAME", // 集群名称
		//GroupName:   "GROUP_NAME",   // 分组名称
	})

	if err != nil || !success {
		log.Fatal("RegisterInstance failed:", err)
	}
	db := sqlx.NewMysql(c.Mysql.Url)
	ctx := svc.NewServiceContext(&c, &db)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		sku.RegisterSkuServiceServer(grpcServer, server.NewSkuServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
