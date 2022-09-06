package config

import (
	"fmt"
	validator "github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	App App `mapstructure:"app"`
}

type App struct {
	Version         string  `mapstructure:"version" validate:"required"`
	DevelopmentMode bool    `mapstructure:"development_mode"`
	Name            string  `mapstructure:"name" validate:"required"`
	Env             string  `mapstructure:"env" validate:"required"`
	ShutdownSec     float64 `mapstructure:"shutdown_sec" validate:"required"`
}

func ParseConfig() (*Config, error) {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath(".")            // Look for config in current directory
	config.AddConfigPath("config/")      // Optionally look for config in the working directory.
	config.AddConfigPath("../config/")   // Look for config needed for tests.
	config.AddConfigPath("../")          // Look for config needed for tests.
	config.AddConfigPath("../../config") // Look for config needed for tests.
	config.AddConfigPath("../../../config")

	config.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	config.AutomaticEnv()

	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		return nil, fmt.Errorf("read config err: %v", err)
	}

	var cfg Config

	err = config.Unmarshal(&cfg)
	if err != nil { // Handle errors reading the config file
		return nil, fmt.Errorf("Fatal error config file: %s \n", err)
	}

	// validate
	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		return nil, fmt.Errorf("fatal error validate config: %v", err)
	}

	return &cfg, nil

}
