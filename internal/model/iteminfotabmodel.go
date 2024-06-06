package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ItemInfoTabModel = (*customItemInfoTabModel)(nil)

type (
	// ItemInfoTabModel is an interface to be customized, add more methods here,
	// and implement the added methods in customItemInfoTabModel.
	ItemInfoTabModel interface {
		itemInfoTabModel
		withSession(session sqlx.Session) ItemInfoTabModel
	}

	customItemInfoTabModel struct {
		*defaultItemInfoTabModel
	}
)

// NewItemInfoTabModel returns a model for the database table.
func NewItemInfoTabModel(conn sqlx.SqlConn) ItemInfoTabModel {
	return &customItemInfoTabModel{
		defaultItemInfoTabModel: newItemInfoTabModel(conn),
	}
}

func (m *customItemInfoTabModel) withSession(session sqlx.Session) ItemInfoTabModel {
	return NewItemInfoTabModel(sqlx.NewSqlConnFromSession(session))
}
