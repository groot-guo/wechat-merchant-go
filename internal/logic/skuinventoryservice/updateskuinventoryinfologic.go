package skuinventoryservicelogic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"

	"wechat-merchant-go/internal/svc"
	"wechat-merchant-go/pb/sku"
)

type UpdateSkuInventoryInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSkuInventoryInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSkuInventoryInfoLogic {
	return &UpdateSkuInventoryInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSkuInventoryInfoLogic) UpdateSkuInventoryInfo(in *sku.UpdateSkuInventoryInfoReq) (*sku.CommonRsp, error) {
	if in.GetSkuId() == "" {
		return nil, errors.New("skuId is empty")
	}
	if in.GetInventoryQty() < 0 {
		return nil, errors.New("inventoryQty is negative")
	}
	//xa, err := dtmgrpc.XaGrpcFromRequest(l.ctx)
	//xa.CallBranch()
	//dtmcli.MustGenGid()

	//err := dtmgrpc.XaLocalTransaction(l.ctx, dtmcli.DBConf{
	//	Driver:   "mysql",
	//	Host:     "127.0.0.1",
	//	Port:     13306,
	//	User:     "root",
	//	Password: "test123456789",
	//	Db:       "wechat-merchant",
	//	Schema:   "",
	//}, func(db *sql.DB, xa *dtmgrpc.XaGrpc) error {
	//	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}))
	//	if err != nil {
	//		return err
	//	}
	//	u := l.svcCtx.Use.ReplaceDB(gormDB).SkuInventoryTab
	//	data, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).First()
	//	if err != nil {
	//		return err
	//	}
	//	if data.Inventory < in.GetInventoryQty() {
	//		return errors.New("inventory is over")
	//	}
	//
	//	result, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).Update(u.Inventory, in.GetInventoryQty())
	//	if err != nil {
	//		l.Errorf("UpdateSkuInventoryInfo error: %v", err)
	//		return err
	//	}
	//	l.Infof("UpdateSkuInventoryInfo: %v", result)
	//	dtmimp.OrString("", "", dtmcli.ResultSuccess)
	//	return nil
	//})
	//if err != nil {
	//	return nil, err
	//}

	u := l.svcCtx.Use.SkuInventoryTab
	data, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).First()
	if err != nil {
		return nil, err
	}
	if data.Inventory < in.GetInventoryQty() {
		return nil, errors.New("inventory is over")
	}

	data.Inventory -= in.GetInventoryQty()

	result, err := u.WithContext(l.ctx).Where(u.SkuID.Eq(in.GetSkuId())).Update(u.Inventory, data.Inventory)
	if err != nil {
		l.Errorf("UpdateSkuInventoryInfo error: %v", err)
		return nil, err
	}
	l.Infof("UpdateSkuInventoryInfo: %v", result)
	return &sku.CommonRsp{
		Code: 0,
		Msg:  "success",
	}, nil
}
