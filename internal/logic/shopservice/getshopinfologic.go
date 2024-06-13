package shopservicelogic

import (
	"context"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopInfoLogic {
	return &GetShopInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetShopInfoLogic) GetShopInfo(in *sku.GetShopInfoReq) (*sku.GetShopInfoResp, error) {
	// todo: add your logic here and delete this line

	return &sku.GetShopInfoResp{}, nil
}
