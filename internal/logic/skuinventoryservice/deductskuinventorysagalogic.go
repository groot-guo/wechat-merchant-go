package skuinventoryservicelogic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"
)

type DeductSkuInventorySagaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductSkuInventorySagaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductSkuInventorySagaLogic {
	return &DeductSkuInventorySagaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductSkuInventorySagaLogic) DeductSkuInventorySaga(in *sku.UpdateSkuInventoryInfoReq) (*sku.CommonRsp, error) {
	u := l.svcCtx.Use.SkuInventoryTab
	data, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).First()
	if err != nil {
		return nil, err
	}
	if data.Inventory < in.GetInventoryQty() {
		return nil, errors.New("inventory is over")
	}

	data.Inventory -= in.GetInventoryQty()

	result, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).Update(u.Inventory, data.Inventory)
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
