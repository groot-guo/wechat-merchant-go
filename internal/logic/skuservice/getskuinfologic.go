package skuservicelogic

import (
	"context"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSkuInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSkuInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSkuInfoLogic {
	return &GetSkuInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSkuInfoLogic) GetSkuInfo(in *sku.SkuReq) (*sku.SkuResp, error) {
	if in.GetPageNumber() == 0 || in.GetPageSize() == 0 {
		return &sku.SkuResp{
			Common: &sku.CommonRsp{
				Code: -1000,
				Msg:  "the page_number or page_size is zero",
			},
		}, nil
	}

	if len(in.GetSkuId()) == 0 || len(in.GetSkuId()) > 200 {
		return &sku.SkuResp{
			Common: &sku.CommonRsp{
				Code: -100,
				Msg:  "the request sku_id len is too long or empty, only support 200 length.",
			},
		}, nil
	}

	u := l.svcCtx.Use.SkuInfoTab
	limit := in.GetPageSize()
	offset := (in.GetPageNumber() - 1) * limit
	dataList, err := u.WithContext(l.ctx).Where(u.SkuID.In(in.GetSkuId()...)).Limit(int(limit)).Offset(int(offset)).Find()
	if err != nil {
		l.Errorf("err:%+v", err)
		return nil, err
	}

	resultList := make([]*sku.SkuInfo, 0)
	for _, data := range dataList {
		pre := &sku.SkuInfo{
			SkuId:       data.SkuID,
			SkuName:     data.SkuName,
			SkuImage:    data.SkuName,
			ItemId:      data.ItemID,
			ItemName:    data.ItemName,
			ProductId:   data.ProductID,
			ProductName: data.ProductName,
			Ctime:       data.Ctime,
			Mtime:       data.Mtime,
		}
		resultList = append(resultList, pre)
	}

	return &sku.SkuResp{
		Common: &sku.CommonRsp{
			Code: 0,
			Msg:  "ok",
		},
		Data: resultList,
	}, nil
}
