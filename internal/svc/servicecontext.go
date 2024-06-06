package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"wechat-merchant-go/internal/config"
)

// 传指针，利用nacos 热更新完成基本操作

type ServiceContext struct {
	Config *config.Config
	Db     *sqlx.SqlConn
}

func NewServiceContext(c *config.Config, db *sqlx.SqlConn) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Db:     db,
	}
}
