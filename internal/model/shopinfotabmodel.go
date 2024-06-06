package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ShopInfoTabModel = (*customShopInfoTabModel)(nil)

type (
	// ShopInfoTabModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShopInfoTabModel.
	ShopInfoTabModel interface {
		shopInfoTabModel
		withSession(session sqlx.Session) ShopInfoTabModel
	}

	customShopInfoTabModel struct {
		*defaultShopInfoTabModel
	}
)

// NewShopInfoTabModel returns a model for the database table.
func NewShopInfoTabModel(conn sqlx.SqlConn) ShopInfoTabModel {
	return &customShopInfoTabModel{
		defaultShopInfoTabModel: newShopInfoTabModel(conn),
	}
}

func (m *customShopInfoTabModel) withSession(session sqlx.Session) ShopInfoTabModel {
	return NewShopInfoTabModel(sqlx.NewSqlConnFromSession(session))
}
