package models

import "time"

type MaterialSyncLog struct {
	LogId    uint64    `gorm:"column:log_id;autoIncrement;primaryKey"`
	LogDate  time.Time `gorm:"column:updated_time;autoUpdateTime;index"`
	Request  string
	Response string
}

func (MaterialSyncLog) TableName() string {
	return "material_sync_log"
}
