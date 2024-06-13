// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"wechat-merchant-go/internal/dao/model"
)

func newSkuInfoTab(db *gorm.DB, opts ...gen.DOOption) skuInfoTab {
	_skuInfoTab := skuInfoTab{}

	_skuInfoTab.skuInfoTabDo.UseDB(db, opts...)
	_skuInfoTab.skuInfoTabDo.UseModel(&model.SkuInfoTab{})

	tableName := _skuInfoTab.skuInfoTabDo.TableName()
	_skuInfoTab.ALL = field.NewAsterisk(tableName)
	_skuInfoTab.ID = field.NewUint64(tableName, "id")
	_skuInfoTab.SkuID = field.NewString(tableName, "sku_id")
	_skuInfoTab.SkuName = field.NewString(tableName, "sku_name")
	_skuInfoTab.ItemID = field.NewInt32(tableName, "item_id")
	_skuInfoTab.ItemName = field.NewString(tableName, "item_name")
	_skuInfoTab.ProductID = field.NewUint32(tableName, "product_id")
	_skuInfoTab.ProductName = field.NewString(tableName, "product_name")
	_skuInfoTab.CategoryID = field.NewUint32(tableName, "category_id")
	_skuInfoTab.CategoryName = field.NewString(tableName, "category_name")
	_skuInfoTab.ShopID = field.NewUint32(tableName, "shop_id")
	_skuInfoTab.Ctime = field.NewUint32(tableName, "ctime")
	_skuInfoTab.Mtime = field.NewUint32(tableName, "mtime")

	_skuInfoTab.fillFieldMap()

	return _skuInfoTab
}

type skuInfoTab struct {
	skuInfoTabDo skuInfoTabDo

	ALL          field.Asterisk
	ID           field.Uint64
	SkuID        field.String
	SkuName      field.String
	ItemID       field.Int32
	ItemName     field.String
	ProductID    field.Uint32
	ProductName  field.String
	CategoryID   field.Uint32
	CategoryName field.String
	ShopID       field.Uint32
	Ctime        field.Uint32
	Mtime        field.Uint32

	fieldMap map[string]field.Expr
}

func (s skuInfoTab) Table(newTableName string) *skuInfoTab {
	s.skuInfoTabDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skuInfoTab) As(alias string) *skuInfoTab {
	s.skuInfoTabDo.DO = *(s.skuInfoTabDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skuInfoTab) updateTableName(table string) *skuInfoTab {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint64(table, "id")
	s.SkuID = field.NewString(table, "sku_id")
	s.SkuName = field.NewString(table, "sku_name")
	s.ItemID = field.NewInt32(table, "item_id")
	s.ItemName = field.NewString(table, "item_name")
	s.ProductID = field.NewUint32(table, "product_id")
	s.ProductName = field.NewString(table, "product_name")
	s.CategoryID = field.NewUint32(table, "category_id")
	s.CategoryName = field.NewString(table, "category_name")
	s.ShopID = field.NewUint32(table, "shop_id")
	s.Ctime = field.NewUint32(table, "ctime")
	s.Mtime = field.NewUint32(table, "mtime")

	s.fillFieldMap()

	return s
}

func (s *skuInfoTab) WithContext(ctx context.Context) *skuInfoTabDo {
	return s.skuInfoTabDo.WithContext(ctx)
}

func (s skuInfoTab) TableName() string { return s.skuInfoTabDo.TableName() }

func (s skuInfoTab) Alias() string { return s.skuInfoTabDo.Alias() }

func (s skuInfoTab) Columns(cols ...field.Expr) gen.Columns { return s.skuInfoTabDo.Columns(cols...) }

func (s *skuInfoTab) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skuInfoTab) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["id"] = s.ID
	s.fieldMap["sku_id"] = s.SkuID
	s.fieldMap["sku_name"] = s.SkuName
	s.fieldMap["item_id"] = s.ItemID
	s.fieldMap["item_name"] = s.ItemName
	s.fieldMap["product_id"] = s.ProductID
	s.fieldMap["product_name"] = s.ProductName
	s.fieldMap["category_id"] = s.CategoryID
	s.fieldMap["category_name"] = s.CategoryName
	s.fieldMap["shop_id"] = s.ShopID
	s.fieldMap["ctime"] = s.Ctime
	s.fieldMap["mtime"] = s.Mtime
}

func (s skuInfoTab) clone(db *gorm.DB) skuInfoTab {
	s.skuInfoTabDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skuInfoTab) replaceDB(db *gorm.DB) skuInfoTab {
	s.skuInfoTabDo.ReplaceDB(db)
	return s
}

type skuInfoTabDo struct{ gen.DO }

func (s skuInfoTabDo) Debug() *skuInfoTabDo {
	return s.withDO(s.DO.Debug())
}

func (s skuInfoTabDo) WithContext(ctx context.Context) *skuInfoTabDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skuInfoTabDo) ReadDB() *skuInfoTabDo {
	return s.Clauses(dbresolver.Read)
}

func (s skuInfoTabDo) WriteDB() *skuInfoTabDo {
	return s.Clauses(dbresolver.Write)
}

func (s skuInfoTabDo) Session(config *gorm.Session) *skuInfoTabDo {
	return s.withDO(s.DO.Session(config))
}

func (s skuInfoTabDo) Clauses(conds ...clause.Expression) *skuInfoTabDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skuInfoTabDo) Returning(value interface{}, columns ...string) *skuInfoTabDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skuInfoTabDo) Not(conds ...gen.Condition) *skuInfoTabDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skuInfoTabDo) Or(conds ...gen.Condition) *skuInfoTabDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skuInfoTabDo) Select(conds ...field.Expr) *skuInfoTabDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skuInfoTabDo) Where(conds ...gen.Condition) *skuInfoTabDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skuInfoTabDo) Order(conds ...field.Expr) *skuInfoTabDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skuInfoTabDo) Distinct(cols ...field.Expr) *skuInfoTabDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skuInfoTabDo) Omit(cols ...field.Expr) *skuInfoTabDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skuInfoTabDo) Join(table schema.Tabler, on ...field.Expr) *skuInfoTabDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skuInfoTabDo) LeftJoin(table schema.Tabler, on ...field.Expr) *skuInfoTabDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skuInfoTabDo) RightJoin(table schema.Tabler, on ...field.Expr) *skuInfoTabDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skuInfoTabDo) Group(cols ...field.Expr) *skuInfoTabDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skuInfoTabDo) Having(conds ...gen.Condition) *skuInfoTabDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skuInfoTabDo) Limit(limit int) *skuInfoTabDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skuInfoTabDo) Offset(offset int) *skuInfoTabDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skuInfoTabDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *skuInfoTabDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skuInfoTabDo) Unscoped() *skuInfoTabDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skuInfoTabDo) Create(values ...*model.SkuInfoTab) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skuInfoTabDo) CreateInBatches(values []*model.SkuInfoTab, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skuInfoTabDo) Save(values ...*model.SkuInfoTab) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skuInfoTabDo) First() (*model.SkuInfoTab, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInfoTab), nil
	}
}

func (s skuInfoTabDo) Take() (*model.SkuInfoTab, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInfoTab), nil
	}
}

func (s skuInfoTabDo) Last() (*model.SkuInfoTab, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInfoTab), nil
	}
}

func (s skuInfoTabDo) Find() ([]*model.SkuInfoTab, error) {
	result, err := s.DO.Find()
	return result.([]*model.SkuInfoTab), err
}

func (s skuInfoTabDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SkuInfoTab, err error) {
	buf := make([]*model.SkuInfoTab, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skuInfoTabDo) FindInBatches(result *[]*model.SkuInfoTab, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skuInfoTabDo) Attrs(attrs ...field.AssignExpr) *skuInfoTabDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skuInfoTabDo) Assign(attrs ...field.AssignExpr) *skuInfoTabDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skuInfoTabDo) Joins(fields ...field.RelationField) *skuInfoTabDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skuInfoTabDo) Preload(fields ...field.RelationField) *skuInfoTabDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skuInfoTabDo) FirstOrInit() (*model.SkuInfoTab, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInfoTab), nil
	}
}

func (s skuInfoTabDo) FirstOrCreate() (*model.SkuInfoTab, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInfoTab), nil
	}
}

func (s skuInfoTabDo) FindByPage(offset int, limit int) (result []*model.SkuInfoTab, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s skuInfoTabDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skuInfoTabDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skuInfoTabDo) Delete(models ...*model.SkuInfoTab) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skuInfoTabDo) withDO(do gen.Dao) *skuInfoTabDo {
	s.DO = *do.(*gen.DO)
	return s
}