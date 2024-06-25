package skuinventoryservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"
)

type DeductSkuInventoryTccSubmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductSkuInventoryTccSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductSkuInventoryTccSubmitLogic {
	return &DeductSkuInventoryTccSubmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductSkuInventoryTccSubmitLogic) DeductSkuInventoryTccSubmit(in *sku.UpdateSkuInventoryInfoReq) (*sku.CommonRsp, error) {
	u := l.svcCtx.Use.SkuInventoryTab
	//data, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).First()
	//if err != nil {
	//	return nil, err
	//}
	//if data.Inventory < in.GetInventoryQty() {
	//	return nil, errors.New("inventory is over")
	//}

	//data.Inventory -= in.GetInventoryQty()

	result, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).UpdateColumnSimple(u.Inventory.Div(in.GetInventoryQty()))
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
