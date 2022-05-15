package util

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ConfigEnv() {
	viper.SetConfigFile(".env")

	viper.SetDefault("APP_HOST", "http://localhost")
	viper.SetDefault("APP_PORT", "8034")
	viper.SetDefault("ADM_USER", "admin")
	viper.SetDefault("ADM_PWD", "admin")
	viper.SetDefault("DB_NAME", "shortcatdb")
	viper.SetDefault("USE_DEFAULT_FRONTEND", true)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("[CONFIG] - Error: %s\n", err.Error()))
	}
}
