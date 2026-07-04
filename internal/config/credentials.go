package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Credentials struct {
	Postgres *Postgres `mapstructure:"Postgres"`
	JWT      *JWT      `mapstructure:"JWT"`
}

type Postgres struct {
	Connstr string `mapstructure:"Connstring"`
}

type JWT struct {
	Secret string `mapstructure:"Secret"`
}

func ReadCredentials() (*Credentials, error) {
	viper.SetConfigName("credentials")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var credentials Credentials
	if err := viper.Unmarshal(&credentials); err != nil {
		return nil, fmt.Errorf("unable to decode config, %w", err)
	}

	return &credentials, nil
}
