package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitViperConfig() {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file, ", err)
	}
}

func GetDatabaseConfig() map[string]any {
	return viper.GetStringMap("database")
}

func GetServerConfig() map[string]any {
	return viper.GetStringMap("App")
}
