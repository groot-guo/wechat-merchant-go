package itemservicelogic

import (
	"context"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateItemInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateItemInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateItemInfoLogic {
	return &UpdateItemInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateItemInfoLogic) UpdateItemInfo(in *sku.CreateOrUpdateItemInfoReq) (*sku.CommonRsp, error) {
	// todo: add your logic here and delete this line

	return &sku.CommonRsp{}, nil
}
