package shopservicelogic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"
)

type UpdateShopInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateShopInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateShopInfoLogic {
	return &UpdateShopInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateShopInfoLogic) UpdateShopInfo(in *sku.UpdateShopReq) (*sku.CommonRsp, error) {
	if in.GetShopId() == 0 || in.GetShopName() == "" {
		return nil, errors.New("params error")
	}

	u := l.svcCtx.Use.ShopInfoTab
	result, err := u.WithContext(l.ctx).Where(u.ShopID.Eq(in.GetShopId())).Update(u.ShopName, in.GetShopName())
	if err != nil {
		l.Errorf("UpdateShopInfo err:%v", err)
		return nil, err
	}

	l.Infof("UpdateShopInfo result: %v", result)
	return &sku.CommonRsp{
		Code: 0,
		Msg:  "success",
	}, nil
}
