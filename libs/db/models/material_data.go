package models

type MaterialData struct {
	MaterialCode string `gorm:"column:material_code;primaryKey"`
}

func (MaterialData) TableName() string {
	return "material_data"
}
