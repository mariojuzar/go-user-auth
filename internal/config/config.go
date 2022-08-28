package config

import (
	"github.com/jinzhu/configor"
	"time"
)

type MainConfig struct {
	ApiConfig  APIConfig
	MongoDB    MongoDBConfig
	AuthConfig AuthConfig
}

type APIConfig struct {
	Host         string        `default:"0.0.0.0" env:"API_HOST"`
	Port         uint          `default:"8080" env:"API_PORT"`
	ReadTimeout  time.Duration `default:"10s" env:"API_READ_TIMEOUT"`
	WriteTimeout time.Duration `default:"10s" env:"API_WRITE_TIMEOUT"`
}

type MongoDBConfig struct {
	Host     string `default:"0.0.0.0" env:"MONGODB_HOST"`
	User     string `env:"MONGODB_USER"`
	Password string `env:"MONGODB_PASSWORD"`
	Port     uint   `default:"27017" env:"MONGODB_PORT"`
	Database string `default:"appointment" env:"MONGODB_DATABASE"`
}

type AuthConfig struct {
	JwtSecret            string        `env:"JWT_SECRET"`
	AccessTokenDuration  time.Duration `env:"JWT_ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `env:"JWT_REFRESH_TOKEN_DURATION"`
}

func Load() (*MainConfig, error) {
	cfg := &MainConfig{}
	err := configor.Load(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
