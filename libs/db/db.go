package db

import (
	"fmt"
	"go-plm/libs/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dbConfig := config.GetDatabaseConfig()

	dsn := fmt.Sprintf("%v", dbConfig["dsn"])
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&MaterialData{},
		&MaterialCodeSequence{},
		&MaterialSyncLog{},
	)
}
