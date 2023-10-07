package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	PRODUCTION_ENV  = "production"
	DEVELOPMENT_ENV = "dev"
)

const DATABASE_DRIVER = "postgres"

type Config struct {
	App       AppConfig
	Database  DatabaseConfig
	Redis     RedisConfig
	Providers ProviderConfig
}

type AppConfig struct {
	Environment         string        `envconfig:"ENV"`
	Port                string        `envconfig:"PORT"`
	EphemeralAccountTTL time.Duration `envconfig:"EPHEMERAL_ACCOUNT_TTL"`
}

type DatabaseConfig struct {
	Host     string `envconfig:"PAY_DB_HOST"`
	Port     int    `envconfig:"PAY_DB_PORT"`
	User     string `envconfig:"PAY_DB_USER"`
	Password string `envconfig:"PAY_DB_PASSWORD"`
	Name     string `envconfig:"PAY_DB_NAME"`
	SSLMode  string `envconfig:"PAY_DB_SSL_MODE"`
	Schema   string `envconfig:"PAY_DB_SCHEMA"`
}

func (d *DatabaseConfig) GetURI() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		DATABASE_DRIVER,
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
		d.SSLMode,
	)
}

type RedisConfig struct {
	Host       string `envconfig:"PAY_REDIS_HOST"`
	Port       string `envconfig:"PAY_REDIS_PORT"`
	Password   string `envconfig:"PAY_REDIS_PASSWORD"`
	Username   string `envconfig:"PAY_REDIS_USERNAME"`
	Namespace  string `envconfig:"PAY_REDIS_NAMESPACE"`
	TLSEnabled bool   `envconfig:"PAY_REDIS_TLS_ENABLED"`
}

type ProviderConfig struct {
	Paystack PaystackConfig
}

type PaystackConfig struct {
	BaseURL   string `envconfig:"BASE_URL"`
	SecretKey string `envconfig:"SECRET_KEY"`
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
