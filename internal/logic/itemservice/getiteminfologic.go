package itemservicelogic

import (
	"context"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetItemInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetItemInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetItemInfoLogic {
	return &GetItemInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetItemInfoLogic) GetItemInfo(in *sku.GetItemInfoReq) (*sku.GetItemInfoResp, error) {
	// todo: add your logic here and delete this line

	return &sku.GetItemInfoResp{}, nil
}
