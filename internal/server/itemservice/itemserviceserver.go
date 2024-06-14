// Code generated by goctl. DO NOT EDIT.
// Source: sku.proto

package server

import (
	"context"

	"wechat-merchant-go/internal/logic/itemservice"
	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"
)

type ItemServiceServer struct {
	svcCtx *svc.ServiceContext
	sku.UnimplementedItemServiceServer
}

func NewItemServiceServer(svcCtx *svc.ServiceContext) *ItemServiceServer {
	return &ItemServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ItemServiceServer) GetItemInfo(ctx context.Context, in *sku.GetItemInfoReq) (*sku.GetItemInfoResp, error) {
	l := itemservicelogic.NewGetItemInfoLogic(ctx, s.svcCtx)
	return l.GetItemInfo(in)
}

func (s *ItemServiceServer) CreateItemInfo(ctx context.Context, in *sku.CreateOrUpdateItemInfoReq) (*sku.CommonRsp, error) {
	l := itemservicelogic.NewCreateItemInfoLogic(ctx, s.svcCtx)
	return l.CreateItemInfo(in)
}

func (s *ItemServiceServer) UpdateItemInfo(ctx context.Context, in *sku.CreateOrUpdateItemInfoReq) (*sku.CommonRsp, error) {
	l := itemservicelogic.NewUpdateItemInfoLogic(ctx, s.svcCtx)
	return l.UpdateItemInfo(in)
}
