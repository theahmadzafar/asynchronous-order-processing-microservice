package config

import (
	"asynchronous-order-processing-microservice/internal/services/notification"
	"asynchronous-order-processing-microservice/internal/services/persistance"
	"asynchronous-order-processing-microservice/internal/services/validation"
	"asynchronous-order-processing-microservice/internal/transport/http"

	"github.com/spf13/viper"
)

type Config struct {
	Server       http.Config         `yaml:"server"`
	Validation   validation.Config   `yaml:"validation"`
	Persistance  persistance.Config  `yaml:"persistance"`
	Notification notification.Config `yaml:"notification"`
}

func New() (*Config, error) {
	viper.Reset()
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := &Config{}
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
