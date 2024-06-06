// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	skuInfoTabFieldNames          = builder.RawFieldNames(&SkuInfoTab{})
	skuInfoTabRows                = strings.Join(skuInfoTabFieldNames, ",")
	skuInfoTabRowsExpectAutoSet   = strings.Join(stringx.Remove(skuInfoTabFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	skuInfoTabRowsWithPlaceHolder = strings.Join(stringx.Remove(skuInfoTabFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	skuInfoTabModel interface {
		Insert(ctx context.Context, data *SkuInfoTab) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*SkuInfoTab, error)
		FindOneByProductId(ctx context.Context, productId uint64) (*SkuInfoTab, error)
		FindOneBySkuId(ctx context.Context, skuId string) (*SkuInfoTab, error)
		Update(ctx context.Context, data *SkuInfoTab) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultSkuInfoTabModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SkuInfoTab struct {
		Id           uint64 `db:"id"`
		SkuId        string `db:"sku_id"`
		SkuName      string `db:"sku_name"`
		ItemId       int64  `db:"item_id"`
		ItemName     string `db:"item_name"`
		ProductId    uint64 `db:"product_id"`
		ProductName  string `db:"product_name"`
		CategoryId   uint64 `db:"category_id"`
		CategoryName string `db:"category_name"`
		ShopId       uint64 `db:"shop_id"`
		Ctime        uint64 `db:"ctime"`
		Mtime        uint64 `db:"mtime"`
	}
)

func newSkuInfoTabModel(conn sqlx.SqlConn) *defaultSkuInfoTabModel {
	return &defaultSkuInfoTabModel{
		conn:  conn,
		table: "`sku_info_tab`",
	}
}

func (m *defaultSkuInfoTabModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSkuInfoTabModel) FindOne(ctx context.Context, id uint64) (*SkuInfoTab, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", skuInfoTabRows, m.table)
	var resp SkuInfoTab
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSkuInfoTabModel) FindOneByProductId(ctx context.Context, productId uint64) (*SkuInfoTab, error) {
	var resp SkuInfoTab
	query := fmt.Sprintf("select %s from %s where `product_id` = ? limit 1", skuInfoTabRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, productId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSkuInfoTabModel) FindOneBySkuId(ctx context.Context, skuId string) (*SkuInfoTab, error) {
	var resp SkuInfoTab
	query := fmt.Sprintf("select %s from %s where `sku_id` = ? limit 1", skuInfoTabRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, skuId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSkuInfoTabModel) Insert(ctx context.Context, data *SkuInfoTab) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, skuInfoTabRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.SkuId, data.SkuName, data.ItemId, data.ItemName, data.ProductId, data.ProductName, data.CategoryId, data.CategoryName, data.ShopId, data.Ctime, data.Mtime)
	return ret, err
}

func (m *defaultSkuInfoTabModel) Update(ctx context.Context, newData *SkuInfoTab) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, skuInfoTabRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.SkuId, newData.SkuName, newData.ItemId, newData.ItemName, newData.ProductId, newData.ProductName, newData.CategoryId, newData.CategoryName, newData.ShopId, newData.Ctime, newData.Mtime, newData.Id)
	return err
}

func (m *defaultSkuInfoTabModel) tableName() string {
	return m.table
}