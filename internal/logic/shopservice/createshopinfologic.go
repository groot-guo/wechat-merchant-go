package shopservicelogic

import (
	"context"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &sku.CommonRsp{}, nil
}
