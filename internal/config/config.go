package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		Url string
	}

	Nacos struct {
		Host      string
		Port      uint64
		Namespace string
		Group     string
	}
}
