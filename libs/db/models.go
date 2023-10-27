package db

import "time"

type MaterialData struct {
	MaterialCode           string `gorm:"column:material_code;primaryKey"`
	MaterialCodeRef        string `gorm:"column:material_code_ref"`
	OldMaterialCode        string `gorm:"column:old_material_code"`
	MaterialName           string `gorm:"column:material_name;not null"`
	Description            string `gorm:"not null;index"`
	MaterialType           string `gorm:"column:material_type;not null"`
	BaseMeasureUnit        string `gorm:"column:base_measure_unit;not null"`
	MaterialGroup          string `gorm:"column:material_group;not null"`
	MaterialGroupName      string `gorm:"column:material_group_name;not null"`
	MaterialStatus         string `gorm:"column:material_status;not null"`
	ProductGroup           string `gorm:"column:product_group;not null"`
	MachineType            string `gorm:"column:machine_type"`
	PurchaseType           string `gorm:"column:purchase_type;not null"`
	SpecialPurchaseType    string `gorm:"column:special_purchase_type;not null"`
	MrpController          string `gorm:"column:mrp_controller;not null"`
	HomemadeProductionDays uint   `gorm:"column:homemade_production_days"`
	PurchaseGroup          string `gorm:"column:purchase_group;not null"`

	CodeType       string `gorm:"column:code_type;not null"`
	MainCategory   string `gorm:"column:main_category;not null"`
	MediumCategory string `gorm:"column:medium_category;not null"`
	SmallCategory  string `gorm:"column:small_category"`
	Creator        string
	Modifier       string
	CreatedTime    time.Time `gorm:"column:created_time;autoCreateTime"`
	UpdatedTime    time.Time `gorm:"column:updated_time;autoUpdateTime;index"`
}

func (MaterialData) TableName() string {
	return "material_data"
}

type MaterialCodeSequence struct {
	MaterialCodeRef string `gorm:"column:material_code_ref;primaryKey"`
	Sequence        uint32 `gorm:"not null"`

	Creator     string
	Modifier    string
	CreatedTime time.Time `gorm:"column:created_time;autoCreateTime"`
	UpdatedTime time.Time `gorm:"column:updated_time;autoUpdateTime;index"`
}

func (MaterialCodeSequence) TableName() string {
	return "material_code_sequence"
}

type MaterialSyncLog struct {
	LogId    uint64    `gorm:"column:log_id;autoIncrement;primaryKey"`
	LogDate  time.Time `gorm:"column:updated_time;autoUpdateTime;index"`
	Request  string
	Response string
}

func (MaterialSyncLog) TableName() string {
	return "material_sync_log"
}
