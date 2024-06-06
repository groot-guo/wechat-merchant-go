package logic

import (
	"context"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSkuInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSkuInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSkuInfoLogic {
	return &GetSkuInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSkuInfoLogic) GetSkuInfo(in *sku.SkuReq) (*sku.SkuResp, error) {
	// todo: add your logic here and delete this line

	return &sku.SkuResp{}, nil
}
