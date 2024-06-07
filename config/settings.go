package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DatabaseURL string `split_words:"true" default:"" envconfig:"DATABASE_URL"`

	DBType       string `split_words:"true" default:"postgres" envconfig:"DatabaseType"`
	DSN          string `split_words:"true" default:"postgres://postgres:postgres@localhost:5432/?sslmode=disable" envconfig:"DatabaseDsn"`
	MaxOpenConns int    `split_words:"true" default:"20" envconfig:"MaxOpenConns"`
	MaxIdleConns int    `split_words:"true" default:"10" envconfig:"MaxIdleConns"`
}

var config Config

func LoadConfig() {
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("Failed to parse configuration: %+v", err)
	}
	log.Infof("Config: %+v", config)
}

func GetConfig() Config {
	return config
}
