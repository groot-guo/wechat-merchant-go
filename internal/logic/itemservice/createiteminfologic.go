package itemservicelogic

import (
	"context"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &sku.CommonRsp{}, nil
}
