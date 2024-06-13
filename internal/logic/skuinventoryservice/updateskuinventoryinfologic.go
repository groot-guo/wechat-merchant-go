package skuinventoryservicelogic

import (
	"context"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSkuInventoryInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSkuInventoryInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSkuInventoryInfoLogic {
	return &UpdateSkuInventoryInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSkuInventoryInfoLogic) UpdateSkuInventoryInfo(in *sku.UpdateSkuInventoryInfoReq) (*sku.CommonRsp, error) {
	// todo: add your logic here and delete this line

	return &sku.CommonRsp{}, nil
}
