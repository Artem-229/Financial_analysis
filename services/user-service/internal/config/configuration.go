package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server *Server `mapstructure:"Server"`
	Kafka  *Kafka  `mapstructure:"Kafka"`
	Cron   *Cron   `mapstructure:"Cron"`
}

type Server struct {
	Port string `mapstructure:"Port"`
}

type Kafka struct {
	Brokers       []string `mapstructure:"brokers"`
	UpstreamTopic string   `mapstructure:"UpstreamTopic"`
	HandleTopic   string   `mapstructure:"HandleTopic"`
}

type Cron struct {
	Processor *Processor `mapstructure:"Processor"`
}

type Processor struct {
	Spec string `mapstructure:"Spec"`
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
