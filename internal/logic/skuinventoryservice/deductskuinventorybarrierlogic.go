package skuinventoryservicelogic

import (
	"context"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductSkuInventoryBarrierLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductSkuInventoryBarrierLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductSkuInventoryBarrierLogic {
	return &DeductSkuInventoryBarrierLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductSkuInventoryBarrierLogic) DeductSkuInventoryBarrier(in *sku.UpdateSkuInventoryInfoReq) (*sku.CommonRsp, error) {
	// todo: add your logic here and delete this line

	return &sku.CommonRsp{}, nil
}
