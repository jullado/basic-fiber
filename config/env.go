package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func InitEnvironment() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Default --------------------------
	viper.SetDefault("app.port", 3000)
	viper.SetDefault("app.env", "production")
	viper.SetDefault("storage.path", "files")
	// -----------------------------------

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Load From Environment")
	}
}
