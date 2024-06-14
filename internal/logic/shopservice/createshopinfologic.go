package shopservicelogic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"

	"wechat-merchant-go/internal/dao/model"
	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"
)

type CreateShopInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateShopInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateShopInfoLogic {
	return &CreateShopInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateShopInfoLogic) CreateShopInfo(in *sku.CreateShopReq) (*sku.CommonRsp, error) {
	if in.GetShopId() == 0 || in.GetShopName() == "" { // todo 一般要确定是否有非法字符，进行确认
		return nil, errors.New("params error")
	}

	u := l.svcCtx.Use.ShopInfoTab
	err := u.WithContext(l.ctx).Create(&model.ShopInfoTab{
		ShopID:   in.GetShopId(),
		ShopName: in.GetShopName(),
	})
	if err != nil {
		return nil, err
	}

	return &sku.CommonRsp{
		Code: 0,
		Msg:  "success",
	}, nil
}
