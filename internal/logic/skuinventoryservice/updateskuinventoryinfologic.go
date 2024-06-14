package skuinventoryservicelogic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"
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
	if in.GetSkuId() == "" {
		return nil, errors.New("skuId is empty")
	}
	if in.GetInventoryQty() < 0 {
		return nil, errors.New("inventoryQty is negative")
	}

	u := l.svcCtx.Use.SkuInventoryTab
	data, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).First()
	if err != nil {
		return nil, err
	}
	if data.Inventory < in.GetInventoryQty() {
		return nil, errors.New("inventory is over")
	}

	result, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).Update(u.Inventory, in.GetInventoryQty())
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
