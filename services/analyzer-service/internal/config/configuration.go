package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Kafka        *Kafka  `mapstructure:"Kafka"`
	Ollama       *Ollama `mapstructure:"Ollama"`
	SystemPrompt string  `mapstructure:"SystemPrompt"`
}

type Kafka struct {
	Brokers         []string      `mapstructure:"brokers"`
	UpstreamTopic   string        `mapstructure:"UpstreamTopic"`
	HandleTopic     string        `mapstructure:"HandleTopic"`
	ConsumerGroupID string        `mapstructure:"ConsumerGroupID"`
	BatchSize       int           `mapstructure:"BatchSize"`
	MaxWait         time.Duration `mapstructure:"MaxWait"`
}

type Ollama struct {
	Host    string        `mapstructure:"Host"`
	Model   string        `mapstructure:"Model"`
	Timeout time.Duration `mapstructure:"Timeout"`
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
		return nil, fmt.Errorf("unmarshalling config error: %w", err)
	}

	return &cfg, nil
}
