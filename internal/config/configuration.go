package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server *Server `mapstructure:"Server"`
}

type Server struct {
	Port string `mapstructure:"Port"`
}

func ReadConfig() (*Config, error) {
	viper.SetConfigName("configuration")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error with reading config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshalhing config error: %w", err)
	}

	return &cfg, nil
}
