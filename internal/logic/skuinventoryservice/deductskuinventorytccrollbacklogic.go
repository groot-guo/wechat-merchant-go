package skuinventoryservicelogic

import (
	"context"
	"errors"
	"gorm.io/gorm"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductSkuInventoryTccRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductSkuInventoryTccRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductSkuInventoryTccRollbackLogic {
	return &DeductSkuInventoryTccRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductSkuInventoryTccRollbackLogic) DeductSkuInventoryTccRollback(in *sku.UpdateSkuInventoryInfoReq) (*sku.CommonRsp, error) {
	u := l.svcCtx.Use.SkuInventoryTab
	_, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &sku.CommonRsp{
			Code: 0,
			Msg:  "success",
		}, nil
	}
	if err != nil {
		l.Errorf("DeductSkuInventoryTccRollback query error: %v", err)
		return nil, err
	}

	result, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).UpdateColumnSimple(u.Inventory.Add(in.GetInventoryQty()))
	if err != nil {
		l.Errorf("DeductSkuInventoryTccRollback update error: %v", err)
		return nil, err
	}
	l.Infof("DeductSkuInventoryTccRollback update : %v", result)
	return &sku.CommonRsp{
		Code: 0,
		Msg:  "success",
	}, nil
}
