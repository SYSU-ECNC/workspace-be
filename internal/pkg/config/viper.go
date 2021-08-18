package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("ecnc")
	viper.AutomaticEnv()

	viper.SetConfigName("config")      // name of config file (without extension)
	viper.SetConfigType("toml")        // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/ecnc/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.ecnc") // call multiple times to add many search paths
	viper.AddConfigPath(".")           // optionally look for config in the working directory
	err := viper.ReadInConfig()        // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func Get(key string) string {
	return viper.GetString(key)
}
