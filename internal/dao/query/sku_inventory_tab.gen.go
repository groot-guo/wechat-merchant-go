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

func newSkuInventoryTab(db *gorm.DB, opts ...gen.DOOption) skuInventoryTab {
	_skuInventoryTab := skuInventoryTab{}

	_skuInventoryTab.skuInventoryTabDo.UseDB(db, opts...)
	_skuInventoryTab.skuInventoryTabDo.UseModel(&model.SkuInventoryTab{})

	tableName := _skuInventoryTab.skuInventoryTabDo.TableName()
	_skuInventoryTab.ALL = field.NewAsterisk(tableName)
	_skuInventoryTab.ID = field.NewUint64(tableName, "id")
	_skuInventoryTab.SkuID = field.NewString(tableName, "sku_id")
	_skuInventoryTab.Inventory = field.NewUint32(tableName, "inventory")
	_skuInventoryTab.Damage = field.NewUint32(tableName, "damage")
	_skuInventoryTab.Ctime = field.NewUint32(tableName, "ctime")
	_skuInventoryTab.Mtime = field.NewUint32(tableName, "mtime")

	_skuInventoryTab.fillFieldMap()

	return _skuInventoryTab
}

type skuInventoryTab struct {
	skuInventoryTabDo skuInventoryTabDo

	ALL       field.Asterisk
	ID        field.Uint64
	SkuID     field.String
	Inventory field.Uint32
	Damage    field.Uint32
	Ctime     field.Uint32
	Mtime     field.Uint32

	fieldMap map[string]field.Expr
}

func (s skuInventoryTab) Table(newTableName string) *skuInventoryTab {
	s.skuInventoryTabDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skuInventoryTab) As(alias string) *skuInventoryTab {
	s.skuInventoryTabDo.DO = *(s.skuInventoryTabDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skuInventoryTab) updateTableName(table string) *skuInventoryTab {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint64(table, "id")
	s.SkuID = field.NewString(table, "sku_id")
	s.Inventory = field.NewUint32(table, "inventory")
	s.Damage = field.NewUint32(table, "damage")
	s.Ctime = field.NewUint32(table, "ctime")
	s.Mtime = field.NewUint32(table, "mtime")

	s.fillFieldMap()

	return s
}

func (s *skuInventoryTab) WithContext(ctx context.Context) *skuInventoryTabDo {
	return s.skuInventoryTabDo.WithContext(ctx)
}

func (s skuInventoryTab) TableName() string { return s.skuInventoryTabDo.TableName() }

func (s skuInventoryTab) Alias() string { return s.skuInventoryTabDo.Alias() }

func (s skuInventoryTab) Columns(cols ...field.Expr) gen.Columns {
	return s.skuInventoryTabDo.Columns(cols...)
}

func (s *skuInventoryTab) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skuInventoryTab) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id"] = s.ID
	s.fieldMap["sku_id"] = s.SkuID
	s.fieldMap["inventory"] = s.Inventory
	s.fieldMap["damage"] = s.Damage
	s.fieldMap["ctime"] = s.Ctime
	s.fieldMap["mtime"] = s.Mtime
}

func (s skuInventoryTab) clone(db *gorm.DB) skuInventoryTab {
	s.skuInventoryTabDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skuInventoryTab) replaceDB(db *gorm.DB) skuInventoryTab {
	s.skuInventoryTabDo.ReplaceDB(db)
	return s
}

type skuInventoryTabDo struct{ gen.DO }

func (s skuInventoryTabDo) Debug() *skuInventoryTabDo {
	return s.withDO(s.DO.Debug())
}

func (s skuInventoryTabDo) WithContext(ctx context.Context) *skuInventoryTabDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skuInventoryTabDo) ReadDB() *skuInventoryTabDo {
	return s.Clauses(dbresolver.Read)
}

func (s skuInventoryTabDo) WriteDB() *skuInventoryTabDo {
	return s.Clauses(dbresolver.Write)
}

func (s skuInventoryTabDo) Session(config *gorm.Session) *skuInventoryTabDo {
	return s.withDO(s.DO.Session(config))
}

func (s skuInventoryTabDo) Clauses(conds ...clause.Expression) *skuInventoryTabDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skuInventoryTabDo) Returning(value interface{}, columns ...string) *skuInventoryTabDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skuInventoryTabDo) Not(conds ...gen.Condition) *skuInventoryTabDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skuInventoryTabDo) Or(conds ...gen.Condition) *skuInventoryTabDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skuInventoryTabDo) Select(conds ...field.Expr) *skuInventoryTabDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skuInventoryTabDo) Where(conds ...gen.Condition) *skuInventoryTabDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skuInventoryTabDo) Order(conds ...field.Expr) *skuInventoryTabDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skuInventoryTabDo) Distinct(cols ...field.Expr) *skuInventoryTabDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skuInventoryTabDo) Omit(cols ...field.Expr) *skuInventoryTabDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skuInventoryTabDo) Join(table schema.Tabler, on ...field.Expr) *skuInventoryTabDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skuInventoryTabDo) LeftJoin(table schema.Tabler, on ...field.Expr) *skuInventoryTabDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skuInventoryTabDo) RightJoin(table schema.Tabler, on ...field.Expr) *skuInventoryTabDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skuInventoryTabDo) Group(cols ...field.Expr) *skuInventoryTabDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skuInventoryTabDo) Having(conds ...gen.Condition) *skuInventoryTabDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skuInventoryTabDo) Limit(limit int) *skuInventoryTabDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skuInventoryTabDo) Offset(offset int) *skuInventoryTabDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skuInventoryTabDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *skuInventoryTabDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skuInventoryTabDo) Unscoped() *skuInventoryTabDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skuInventoryTabDo) Create(values ...*model.SkuInventoryTab) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skuInventoryTabDo) CreateInBatches(values []*model.SkuInventoryTab, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skuInventoryTabDo) Save(values ...*model.SkuInventoryTab) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skuInventoryTabDo) First() (*model.SkuInventoryTab, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventoryTab), nil
	}
}

func (s skuInventoryTabDo) Take() (*model.SkuInventoryTab, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventoryTab), nil
	}
}

func (s skuInventoryTabDo) Last() (*model.SkuInventoryTab, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventoryTab), nil
	}
}

func (s skuInventoryTabDo) Find() ([]*model.SkuInventoryTab, error) {
	result, err := s.DO.Find()
	return result.([]*model.SkuInventoryTab), err
}

func (s skuInventoryTabDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SkuInventoryTab, err error) {
	buf := make([]*model.SkuInventoryTab, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skuInventoryTabDo) FindInBatches(result *[]*model.SkuInventoryTab, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skuInventoryTabDo) Attrs(attrs ...field.AssignExpr) *skuInventoryTabDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skuInventoryTabDo) Assign(attrs ...field.AssignExpr) *skuInventoryTabDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skuInventoryTabDo) Joins(fields ...field.RelationField) *skuInventoryTabDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skuInventoryTabDo) Preload(fields ...field.RelationField) *skuInventoryTabDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skuInventoryTabDo) FirstOrInit() (*model.SkuInventoryTab, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventoryTab), nil
	}
}

func (s skuInventoryTabDo) FirstOrCreate() (*model.SkuInventoryTab, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SkuInventoryTab), nil
	}
}

func (s skuInventoryTabDo) FindByPage(offset int, limit int) (result []*model.SkuInventoryTab, count int64, err error) {
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

func (s skuInventoryTabDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skuInventoryTabDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skuInventoryTabDo) Delete(models ...*model.SkuInventoryTab) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skuInventoryTabDo) withDO(do gen.Dao) *skuInventoryTabDo {
	s.DO = *do.(*gen.DO)
	return s
}
