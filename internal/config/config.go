package config

import (
	"github.com/dtm-labs/client/dtmcli"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Dtm dtmcli.DBConf

	Mysql struct {
		Url string
	}
}
