package config

import (
	"fmt"

	"log"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type (
	NacosService struct {
		RegisterClient naming_client.INamingClient
		ConfigClient   config_client.IConfigClient
		AppConfig      *NacosConfig
	}
	NacosConfig struct {
		Host        string
		GrpcPort    uint64
		Namespace   string
		Group       string
		ServiceName string
		DataId      string
		Weight      float64
		Ip          string
		Port        uint64
		LogLevel    string
	}
)

func NewNacosService(config *NacosConfig) (*NacosService, error) {
	param := vo.NacosClientParam{
		ClientConfig: &constant.ClientConfig{
			NamespaceId: config.Namespace,
			CacheDir:    ".\\nacos\\cache",
			LogLevel:    config.LogLevel,
			LogDir:      ".\\ncaos\\logs",
			AppName:     config.ServiceName,
		},
		ServerConfigs: []constant.ServerConfig{
			{
				IpAddr:   config.Host,
				GrpcPort: config.GrpcPort,
			},
		},
	}

	namingClient, err := clients.NewNamingClient(param)
	if err != nil {
		return nil, err
	}

	configInfo, err := clients.NewConfigClient(param)
	if err != nil {
		log.Fatal(err)
	}

	return &NacosService{
		RegisterClient: namingClient,
		ConfigClient:   configInfo,
		AppConfig:      config,
	}, nil
}

func (this *NacosService) Register() error {
	// 服务注册
	success, err := this.RegisterClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          this.AppConfig.Ip,
		Port:        this.AppConfig.Port,
		ServiceName: this.AppConfig.ServiceName,
		Weight:      this.AppConfig.Weight,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"project": "go-zero"},
		GroupName:   this.AppConfig.Group, // 分组名称
	})
	if err != nil || !success {
		log.Fatal("RegisterInstance failed:", err)
		return err
	}
	return nil
}

// ListenConfig 监听config 并更新, 在线上环境应当先获取
func (this *NacosService) ListenConfig() error {
	return this.ConfigClient.ListenConfig(vo.ConfigParam{
		DataId:  this.AppConfig.DataId,
		Group:   this.AppConfig.Group,
		AppName: this.AppConfig.ServiceName,
		OnChange: func(namespace, group, dataId, data string) {
			log.Fatalln(namespace + group + dataId + data)
		},
	})
}

func (this *NacosService) GetConfig() error {
	s, err := this.ConfigClient.GetConfig(vo.ConfigParam{
		DataId:  this.AppConfig.DataId,
		Group:   this.AppConfig.Group,
		AppName: this.AppConfig.ServiceName,
	})
	if err != nil {
		return err
	}
	log.Println(s)
	fmt.Println(s)
	return nil
}

func (this *NacosService) Close() {
	this.ConfigClient.CloseClient() // 程序结束关闭
	this.RegisterClient.CloseClient()
}
