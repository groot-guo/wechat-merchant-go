// Code generated by goctl. DO NOT EDIT.
// Source: sku.proto

package skuservice

import (
	"context"

	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommonRsp            = sku.CommonRsp
	CreateOrUpdateSkuReq = sku.CreateOrUpdateSkuReq
	SkuInfo              = sku.SkuInfo
	SkuReq               = sku.SkuReq
	SkuResp              = sku.SkuResp

	SkuService interface {
		GetSkuInfo(ctx context.Context, in *SkuReq, opts ...grpc.CallOption) (*SkuResp, error)
		CreateSkuInfo(ctx context.Context, in *CreateOrUpdateSkuReq, opts ...grpc.CallOption) (*CommonRsp, error)
		Update(ctx context.Context, in *CreateOrUpdateSkuReq, opts ...grpc.CallOption) (*CommonRsp, error)
	}

	defaultSkuService struct {
		cli zrpc.Client
	}
)

func NewSkuService(cli zrpc.Client) SkuService {
	return &defaultSkuService{
		cli: cli,
	}
}

func (m *defaultSkuService) GetSkuInfo(ctx context.Context, in *SkuReq, opts ...grpc.CallOption) (*SkuResp, error) {
	client := sku.NewSkuServiceClient(m.cli.Conn())
	return client.GetSkuInfo(ctx, in, opts...)
}

func (m *defaultSkuService) CreateSkuInfo(ctx context.Context, in *CreateOrUpdateSkuReq, opts ...grpc.CallOption) (*CommonRsp, error) {
	client := sku.NewSkuServiceClient(m.cli.Conn())
	return client.CreateSkuInfo(ctx, in, opts...)
}

func (m *defaultSkuService) Update(ctx context.Context, in *CreateOrUpdateSkuReq, opts ...grpc.CallOption) (*CommonRsp, error) {
	client := sku.NewSkuServiceClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}
