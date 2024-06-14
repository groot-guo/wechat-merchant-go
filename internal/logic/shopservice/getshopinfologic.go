package shopservicelogic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"
)

type GetShopInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopInfoLogic {
	return &GetShopInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetShopInfoLogic) GetShopInfo(in *sku.GetShopInfoReq) (*sku.GetShopInfoResp, error) {
	if in.GetShopId() == 0 || in.GetShopName() == "" {
		return nil, errors.New("params error")
	}
	u := l.svcCtx.Use.ShopInfoTab
	data, err := u.WithContext(l.ctx).Where(u.ShopID.Eq(in.GetShopId())).First()
	if err != nil {
		return nil, err
	}

	return &sku.GetShopInfoResp{
		Common: &sku.CommonRsp{
			Code: 0,
			Msg:  "success",
		},
		ShopId:   data.ShopID,
		ShopName: data.ShopName,
	}, nil
}
