package db

import (
	"fmt"
	"go-plm/libs/config"
	"go-plm/libs/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dbConfig := config.GetDatabaseConfig()

	dsn := fmt.Sprintf("%v", dbConfig["dsn"])
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&models.MaterialData{},
		&models.MaterialCodeSequence{},
		&models.MaterialSyncLog{},
	)
}
