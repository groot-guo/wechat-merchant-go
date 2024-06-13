package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"
)

type GetSkuInventoryInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSkuInventoryInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSkuInventoryInfoLogic {
	return &GetSkuInventoryInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSkuInventoryInfoLogic) GetSkuInventoryInfo(in *sku.SkuInventoryReq) (*sku.SkuInventoryResp, error) {

	u := l.svcCtx.Use.SkuInventoryTab
	dataList, err := u.WithContext(l.ctx).Where(u.SkuID.In(in.GetSkuId()...)).Find()
	if err != nil {
		l.Infof("err:%v", err)
		return nil, err
	}

	resultData := make([]*sku.SkuInventoryInfo, 0, len(dataList))
	for _, data := range dataList {
		resultData = append(resultData, &sku.SkuInventoryInfo{
			SkuId:        data.SkuID,
			InventoryQty: data.Inventory,
			DamageQty:    data.Damage,
		})
	}

	return &sku.SkuInventoryResp{
		Common: &sku.CommonRsp{
			Code: 0,
			Msg:  "ok",
		},
		Data: resultData,
	}, nil
}
