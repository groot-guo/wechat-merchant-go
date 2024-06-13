package shopservicelogic

import (
	"context"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateShopInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateShopInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateShopInfoLogic {
	return &UpdateShopInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateShopInfoLogic) UpdateShopInfo(in *sku.UpdateShopReq) (*sku.CommonRsp, error) {
	// todo: add your logic here and delete this line

	return &sku.CommonRsp{}, nil
}
