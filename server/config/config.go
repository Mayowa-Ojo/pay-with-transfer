package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	PRODUCTION_ENV  = "production"
	DEVELOPMENT_ENV = "dev"
)

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Environment string `envconfig:"ENV"`
	Port        string `envconfig:"PORT"`
}

func Load() (*Config, error) {
	if strings.EqualFold(os.Getenv("ENV"), DEVELOPMENT_ENV) {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
