package itemservicelogic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"

	"wechat-merchant-go/internal/dao/model"
	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"
)

type CreateItemInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateItemInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateItemInfoLogic {
	return &CreateItemInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateItemInfoLogic) CreateItemInfo(in *sku.CreateOrUpdateItemInfoReq) (*sku.CommonRsp, error) {
	if in.GetShopId() == 0 || in.GetItemId() == 0 || in.GetItemName() == "" || in.GetCategoryId() == 0 || in.GetCategoryNam() == "" {
		return nil, errors.New("params error")
	}

	shopTab := l.svcCtx.Use.ShopInfoTab
	shopData, err := shopTab.WithContext(l.ctx).Where(shopTab.ShopID.Eq(in.GetShopId())).First()
	if err != nil {
		l.Errorf("CreateItemInfo error:%+v", err)
		return nil, err
	}
	if shopData == nil {
		return nil, errors.New("shop not found")
	}

	// todo category id 校验, 忘记加这张表了

	u := l.svcCtx.Use.ItemInfoTab
	err = u.WithContext(l.ctx).Create(&model.ItemInfoTab{
		ItemID:       in.GetItemId(),
		ItemName:     in.GetItemName(),
		CategoryID:   in.GetCategoryId(),
		CategoryName: in.GetCategoryNam(),
		ShopID:       in.GetShopId(),
	})
	if err != nil {
		l.Errorf("CreateItemInfo error:%+v", err)
		return nil, err
	}

	return &sku.CommonRsp{
		Code: 0,
		Msg:  "success",
	}, nil
}
