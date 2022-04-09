package config

import (
	_ "embed" //pkg.go.dev/embed
	"os"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	vpr *viper.Viper
}

func ConfigSetup(s string) (*Config, error) {
	config := Config{}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		zap.L().Error("Error loading config file", zap.Any("error", err.Error()))
		os.Exit(1)
	}
	// Look to see if a specific environment is configured
	if s == "" {
		s = "default"
	}
	config.vpr = viper.Sub(s)
	config.vpr.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.vpr.AutomaticEnv()
	return &config, nil
}

func (cfg *Config) GetString(k string) string {
	return cfg.vpr.GetString(k)
}

func (cfg *Config) GetInt(k string) int {
	return cfg.vpr.GetInt(k)
}

func (cfg *Config) GetBool(k string) bool {
	return cfg.vpr.GetBool(k)
}
