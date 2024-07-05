package skuinventoryservicelogic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"
)

type DeductSkuInventoryTccTryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductSkuInventoryTccTryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductSkuInventoryTccTryLogic {
	return &DeductSkuInventoryTccTryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductSkuInventoryTccTryLogic) DeductSkuInventoryTccTry(in *sku.UpdateSkuInventoryInfoReq) (*sku.CommonRsp, error) {
	u := l.svcCtx.Use.SkuInventoryTab
	data, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).First()
	if err != nil {
		return nil, status.New(codes.Aborted, "inventory is over").Err()
	}
	if data.Inventory < in.GetInventoryQty() {
		return nil, errors.New("inventory is over")
	}

	return &sku.CommonRsp{
		Code: 0,
		Msg:  "success",
	}, nil
}
