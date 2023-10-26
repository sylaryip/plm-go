package config

import (
	"log"

	"github.com/spf13/viper"
)

func GetDatabaseConfig() map[string]any {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file, ", err)
	}

	return viper.GetStringMap("database")
}
