// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSkuInventoryTab = "sku_inventory_tab"

// SkuInventoryTab mapped from table <sku_inventory_tab>
type SkuInventoryTab struct {
	ID        uint64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SkuID     string `gorm:"column:sku_id;not null" json:"sku_id"`
	Inventory uint32 `gorm:"column:inventory;not null" json:"inventory"`
	Damage    uint32 `gorm:"column:damage;not null" json:"damage"`
	Ctime     uint32 `gorm:"column:ctime;not null" json:"ctime"`
	Mtime     uint32 `gorm:"column:mtime;not null" json:"mtime"`
}

// TableName SkuInventoryTab's table name
func (*SkuInventoryTab) TableName() string {
	return TableNameSkuInventoryTab
}