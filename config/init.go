package config

import (
	"log"

	"github.com/spf13/viper"
)

// Init ...
func Init() error {
	log.Println("Read config file ...")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
