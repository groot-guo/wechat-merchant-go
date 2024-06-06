package logic

import (
	"context"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSkuInventoryInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSkuInventoryInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSkuInventoryInfoLogic {
	return &GetSkuInventoryInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSkuInventoryInfoLogic) GetSkuInventoryInfo(in *sku.SkuInventoryReq) (*sku.SkuInventoryResp, error) {
	// todo: add your logic here and delete this line

	return &sku.SkuInventoryResp{}, nil
}
