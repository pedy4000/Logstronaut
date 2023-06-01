package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment        string `mapstructure:"ENVIRONMENT"`
	DBSource           string `mapstructure:"DB_ADDRESS"`
	ReceiverAddress    string `mapstructure:"RECEIVER_ADDRESS"`
	StorageDestination string `mapstructure:"STORAGE_DESTINATION"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	// using $(path)/app.env as the config file
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// overwite config with environment variables if they exist
	viper.AutomaticEnv()

	// read config file from disk
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// unmarshal config file into Config struct
	err = viper.Unmarshal(&config)
	return
}
