package skuinventoryservicelogic

import (
	"context"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/core/logx"
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
	if len(in.GetSkuId()) > 200 || len(in.GetSkuId()) == 0 {
		return &sku.SkuInventoryResp{
			Common: &sku.CommonRsp{ // todo 换统一的 error ，改在 proto 协议中
				Code: -100,
				Msg:  "the request sku_id len is too long or empty, only support 200 length.",
			},
		}, nil
	}
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
