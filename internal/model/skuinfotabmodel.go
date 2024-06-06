package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SkuInfoTabModel = (*customSkuInfoTabModel)(nil)

type (
	// SkuInfoTabModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSkuInfoTabModel.
	SkuInfoTabModel interface {
		skuInfoTabModel
		withSession(session sqlx.Session) SkuInfoTabModel
	}

	customSkuInfoTabModel struct {
		*defaultSkuInfoTabModel
	}
)

// NewSkuInfoTabModel returns a model for the database table.
func NewSkuInfoTabModel(conn sqlx.SqlConn) SkuInfoTabModel {
	return &customSkuInfoTabModel{
		defaultSkuInfoTabModel: newSkuInfoTabModel(conn),
	}
}

func (m *customSkuInfoTabModel) withSession(session sqlx.Session) SkuInfoTabModel {
	return NewSkuInfoTabModel(sqlx.NewSqlConnFromSession(session))
}
