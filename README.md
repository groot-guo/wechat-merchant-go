# wechat-merchant-go

-----

## Description

​	本项目主要为微信小程序中，开发一个商城小程序demo 的后端服务。主要负责功能：商品信息管理，商品库存管理。与另一个 `wechat-merchant-java` 后端项目 java 开发互相协作。学习相关 框架使用，熟悉微服务框架的使用流程。

​	基于 go-zero 微服务，提供 `rpc` 服务和 http 服务能力，利用 `etcd` 或 `nacos` 作为注册中心和配置中心，利用 `dtm`与 java spring 服务进行分布式事务的处理。

### 已开发功能：

- [ ] 商品信息
- [ ] 商品库存

### 组件

- [ ] `nacos`
- [ ] `etcd`
- [ ] `sqlx` （go-zero 自带）
- [ ] gorm
- [ ] `dtm` 

### 使用时遇到的问题（解决打✔️）

- [ ] go-zero `sqlx` 执行 in 查询时，占位符解析不支持 `[]interface` 类型，支持`[]uint8`
- [ ] `nacos` 初始化客户端服务时，注册微服务信息，需要打开本地缓存，但不会自动创建，windows 中，文件存在，无法打开，go 可以打开这个文件
- [ ] go-zero 连接 docker 启动的 etcd 服务，会有不断重试的警告，需要改为本地启动（或者 k8s 集群？应该可以，需要尝试）  

### 根据 db 生成 go-zero model 代码
```shell
  goctl model mysql datasource -d ./internal/model -t "item_info_tab,shop_info_tab,sku_info_tab,sku_inventory_tab" --url "root:test123456@tcp(127.0.0.1:13306)/wechat-merchant" --strict=true
```

### 根据 proto 文件生成项目代码
不需要生成 client 代码, 加 -m 生成多个服务
```shell
  goctl rpc protoc ./pb/sku.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. -client=false -m 
```

### gorm 生成 code
使用 `gentool` 工具生成代码，下面是生成命令，指定数据库和表
```shell
 gentool -db mysql -dsn "root:test123456@tcp(127.0.0.1:13306)/wechat-merchant" -fieldSignable=true  -tables="item_info_tab,shop_info_tab,sku_info_tab,sku_inventory_tab" -withUnitTest=true
```

### 接入 dtm
```shell
 go get github.com/dtm-labs/client
```

### 注意
使用 `gotcl` 安装 protoc 命令时，仅安装了命令，没有携带包含protoc 文件中protobuf 原生的proto文件（也可能我没找到？），会导致生成文件时，出错。
