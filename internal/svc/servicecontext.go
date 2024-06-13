package svc

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"wechat-merchant-go/internal/config"
	"wechat-merchant-go/internal/dao/query"
)

// 传指针，利用nacos 热更新完成基本操作

type ServiceContext struct {
	Config *config.Config
	Use    *query.Query
}

func NewServiceContext(c *config.Config) *ServiceContext {
	// db init
	// go-zero sqlx 在 in 查询上，对于数组传值不友好
	//db := sqlx.NewMysql(c.Mysql.Url)

	// init gorm db todo log 初始化
	db, err := gorm.Open(mysql.Open(c.Mysql.Url), &gorm.Config{
		Logger: config.NewGormLog(logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		}),
	})
	if err != nil {
		panic("failed to connect database")
	}
	//
	//err = db.Use(prometheus.New(prometheus.Config{
	//	DBName:          "wetchant-merchant",         // 使用 `DBName` 作为指标 label
	//	RefreshInterval: 15,                          // 指标刷新频率（默认为 15 秒）
	//	PushAddr:        "prometheus pusher address", // 如果配置了 `PushAddr`，则推送指标
	//	StartServer:     true,                        // 启用一个 http 服务来暴露指标
	//	HTTPServerPort:  8086,                        // 配置 http 服务监听端口，默认端口为 8080 （如果您配置了多个，只有第一个 `HTTPServerPort` 会被使用）
	//	MetricsCollector: []prometheus.MetricsCollector{
	//		&prometheus.MySQL{
	//			VariableNames: []string{"Threads_running"},
	//		},
	//	}, // 用户自定义指标
	//}))
	//if err != nil {
	//	panic("failed to connect database")
	//}

	use := query.Use(db)
	return &ServiceContext{
		Config: c,
		Use:    use,
	}
}
