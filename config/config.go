package config

import (
	"github.com/phuslu/log"
	"github.com/spf13/viper"
)

type Config struct {
	Listen string `mapstructure:"LISTEN"`

	GinMode string `mapstructure:"GIN_MODE"`

	LogLevel   string `mapstructure:"LOG_LEVEL"`
	LogMaxsize int64  `mapstructure:"LOG_MAX_SIZE"`
	LogBackups int    `mapstructure:"LOG_BACKUPS"`

	DB_Driver string `mapstructure:"DB_DRIVER"`
	DB_DSN    string `mapstructure:"DB_DSN"`

	CookieSecret string `mapstructure:"COOKIE_SECRET"`
}

var config *Config

func Init() (err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.SetDefault("LISTEN", "0.0.0.0:8000")
	viper.SetDefault("GIN_MODE", "debug")
	viper.SetDefault("LOG_LEVEL", "DEBUG")
	viper.SetDefault("LOG_MAX_SIZE", 2147483648)
	viper.SetDefault("LOG_BACKUPS", 5)
	viper.SetDefault("DB_DRIVER", "sqlite")
	viper.SetDefault("DB_DSN", "./local.db")
	viper.SetDefault("COOKIE_SECRET", "5PYiBroBQVm06OEKpJAOI2KkaESmKWrWc7Ew")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Info().Err(err).Msg("read config file failed, use default config")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Info().Err(err).Msg("unmarshal config file failed")
		return
	}
	return
}

func GetConfig() *Config {
	return config
}
