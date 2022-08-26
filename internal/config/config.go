package config

import (
	"github.com/jinzhu/configor"
	"time"
)

type MainConfig struct {
	ApiConfig APIConfig
}

type APIConfig struct {
	Host         string        `default:"0.0.0.0" env:"API_HOST"`
	Port         uint          `default:"8080" env:"API_PORT"`
	ReadTimeout  time.Duration `default:"10s" env:"API_READ_TIMEOUT"`
	WriteTimeout time.Duration `default:"10s" env:"API_WRITE_TIMEOUT"`
}

func Load() (*MainConfig, error) {
	cfg := &MainConfig{}
	err := configor.Load(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
