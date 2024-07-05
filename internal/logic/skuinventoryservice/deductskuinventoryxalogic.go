package skuinventoryservicelogic

import (
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"
)

type DeductSkuInventoryXaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductSkuInventoryXaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductSkuInventoryXaLogic {
	return &DeductSkuInventoryXaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductSkuInventoryXaLogic) DeductSkuInventoryXa(in *sku.UpdateSkuInventoryInfoReq) (*sku.CommonRsp, error) {
	if in.GetSkuId() == "" {
		return nil, errors.New("skuId is empty")
	}
	if in.GetInventoryQty() < 0 {
		return nil, errors.New("inventoryQty is negative")
	}

	err := dtmgrpc.XaLocalTransaction(l.ctx, l.svcCtx.Config.Dtm, func(db *sql.DB, xa *dtmgrpc.XaGrpc) error {
		gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}))
		if err != nil {
			return err
		}
		u := l.svcCtx.Use.ReplaceDB(gormDB).SkuInventoryTab
		data, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).First()
		if err != nil {
			return err
		}
		if data.Inventory < in.GetInventoryQty() {
			return status.New(codes.FailedPrecondition, "inventory is over").Err()
		}
		data.Inventory -= in.GetInventoryQty()

		result, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).Update(u.Inventory, in.GetInventoryQty())
		if err != nil {
			l.Errorf("UpdateSkuInventoryInfo error: %v", err)
			return err
		}
		l.Infof("UpdateSkuInventoryInfo: %v", result)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &sku.CommonRsp{
		Code: 0,
		Msg:  "success",
	}, nil
}
