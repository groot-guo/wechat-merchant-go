package skuinventoryservicelogic

import (
	"context"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductSkuInventorySagaRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductSkuInventorySagaRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductSkuInventorySagaRollbackLogic {
	return &DeductSkuInventorySagaRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductSkuInventorySagaRollbackLogic) DeductSkuInventorySagaRollback(in *sku.UpdateSkuInventoryInfoReq) (*sku.CommonRsp, error) {
	u := l.svcCtx.Use.SkuInventoryTab

	result, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).UpdateColumnSimple(u.Inventory.Add(in.GetInventoryQty()))
	if err != nil {
		l.Errorf("UpdateSkuInventoryInfo error: %v", err)
		return nil, err
	}
	l.Infof("UpdateSkuInventoryInfo: %v", result)
	return &sku.CommonRsp{
		Code: 0,
		Msg:  "success",
	}, nil
}
