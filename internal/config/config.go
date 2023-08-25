package config

import (
	"flag"
	"os"
)

type Config struct {
	Env string
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) ParseFlags() error {
	flag.StringVar(&cfg.Env, "env", cfg.GetOsVar("ENV", "development"), "Sets environment variable")

	return nil
}

func (cfg *Config) IsProd() bool {
	return cfg.Env == "production"
}

func (cfg *Config) IsDev() bool {
	return !cfg.IsProd()
}

func (cfg *Config) GetOsVar(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
