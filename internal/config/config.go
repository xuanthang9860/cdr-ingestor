package config

import (
	"github.com/cc-integration-team/cc-pkg/v2/pkg/logger"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig           `mapstructure:"app"`
	Postgres PostgresConfig      `mapstructure:"postgres"`
	RabbitMQ RabbitMQConfig      `mapstructure:"rabbitmq"`
	Logger   logger.LoggerConfig `mapstructure:"logger"`
}

type AppConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type RabbitMQConfig struct {
	URL       string `mapstructure:"url"`
	QueueName string `mapstructure:"queue_name"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	v := viper.NewWithOptions(viper.KeyDelimiter(":"))
	v.AddConfigPath(path)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
