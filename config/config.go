package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func ReadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	configPath := os.Getenv("CONFIG_PATH")
	if configPath != "" {
		viper.AddConfigPath(configPath)
	} else {
		viper.AddConfigPath("./config")
		err := viper.ReadInConfig()

		if err != nil {
			log.Printf("error while reading config: %s", err.Error())
			viper.SetDefault("port", 8080)
		}
	}

	return nil
}
