// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"wechat-merchant-go/internal/dao/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := _gen_test_db.AutoMigrate(&model.ShopInfoTab{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.ShopInfoTab{}) fail: %s", err)
	}
}

func Test_shopInfoTabQuery(t *testing.T) {
	shopInfoTab := newShopInfoTab(_gen_test_db)
	shopInfoTab = *shopInfoTab.As(shopInfoTab.TableName())
	_do := shopInfoTab.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(shopInfoTab.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <shop_info_tab> fail:", err)
		return
	}

	_, ok := shopInfoTab.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from shopInfoTab success")
	}

	err = _do.Create(&model.ShopInfoTab{})
	if err != nil {
		t.Error("create item in table <shop_info_tab> fail:", err)
	}

	err = _do.Save(&model.ShopInfoTab{})
	if err != nil {
		t.Error("create item in table <shop_info_tab> fail:", err)
	}

	err = _do.CreateInBatches([]*model.ShopInfoTab{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <shop_info_tab> fail:", err)
	}

	_, err = _do.Select(shopInfoTab.ALL).Take()
	if err != nil {
		t.Error("Take() on table <shop_info_tab> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <shop_info_tab> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <shop_info_tab> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <shop_info_tab> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.ShopInfoTab{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <shop_info_tab> fail:", err)
	}

	_, err = _do.Select(shopInfoTab.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <shop_info_tab> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <shop_info_tab> fail:", err)
	}

	_, err = _do.Select(shopInfoTab.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <shop_info_tab> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <shop_info_tab> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <shop_info_tab> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <shop_info_tab> fail:", err)
	}

	_, err = _do.ScanByPage(&model.ShopInfoTab{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <shop_info_tab> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <shop_info_tab> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <shop_info_tab> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <shop_info_tab> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <shop_info_tab> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <shop_info_tab> fail:", err)
	}
}
