package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SkuInventoryTabModel = (*customSkuInventoryTabModel)(nil)

type (
	// SkuInventoryTabModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSkuInventoryTabModel.
	SkuInventoryTabModel interface {
		skuInventoryTabModel
		withSession(session sqlx.Session) SkuInventoryTabModel
	}

	customSkuInventoryTabModel struct {
		*defaultSkuInventoryTabModel
	}
)

// NewSkuInventoryTabModel returns a model for the database table.
func NewSkuInventoryTabModel(conn sqlx.SqlConn) SkuInventoryTabModel {
	return &customSkuInventoryTabModel{
		defaultSkuInventoryTabModel: newSkuInventoryTabModel(conn),
	}
}

func (m *customSkuInventoryTabModel) withSession(session sqlx.Session) SkuInventoryTabModel {
	return NewSkuInventoryTabModel(sqlx.NewSqlConnFromSession(session))
}
