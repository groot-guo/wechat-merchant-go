package svc

import (
	"gorm.io/gorm"
	"wechat-merchant-go/internal/config"
	"wechat-merchant-go/internal/dao/query"
)

// 传指针，利用nacos 热更新完成基本操作

type ServiceContext struct {
	Config *config.Config
	Use    *query.Query
}

func NewServiceContext(c *config.Config, db *gorm.DB) *ServiceContext {
	// db init
	// go-zero sqlx 在 in 查询上，对于数组传值不友好
	//db := sqlx.NewMysql(c.Mysql.Url)

	return &ServiceContext{
		Config: c,
		Use:    query.Use(db),
	}
}
