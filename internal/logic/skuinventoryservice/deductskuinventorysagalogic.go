package skuinventoryservicelogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
		return &sku.CommonRsp{}, status.New(codes.Aborted, "inventory is over").Err()
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
