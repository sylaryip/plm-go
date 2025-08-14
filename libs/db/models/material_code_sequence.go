package models

import "time"

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
