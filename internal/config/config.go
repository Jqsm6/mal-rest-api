package config

import (
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		Host         string        `yaml:"Host" env-default:"localhost"`
		Port         string        `yaml:"Port" env-default:"8080"`
		LoggingLevel string        `yaml:"LoggingLevel" env-default:""`
		ReadTimeout  time.Duration `yaml:"ReadTimeout" env-default:"15s"`
		WriteTimeout time.Duration `yaml:"WriteTimeout" env-default:"15s"`
	} `yaml:"server"`
	Redis struct {
		Host      string        `yaml:"Host" env-default:"localhost"`
		Port      string        `yaml:"Port" env-default:"6379"`
		Password  string        `yaml:"Password" env-default:""`
		Database  int           `yaml:"Database" env-default:"0"`
		CacheTime time.Duration `yaml:"CacheTime" env-default:"1h"`
	}
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{}

		err := cleanenv.ReadConfig("config.yml", config)
		if err != nil {
			help, _ := cleanenv.GetDescription(config, nil)
			println(help)
			panic(err)
		}
	})

	return config
}
