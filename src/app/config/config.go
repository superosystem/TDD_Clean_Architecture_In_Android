package config

import (
	"github.com/spf13/viper"
	"log"
)

func GetEnvValue(key string) string {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading config: %s", err)
	}

	return viper.GetString(key)
}
