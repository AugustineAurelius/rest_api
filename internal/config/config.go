package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"10s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	configPatch := "./../../config/local.yaml"
	if configPatch == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	//check if file exist

	if _, err := os.Stat(configPatch); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPatch)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPatch, &cfg); err != nil {
		log.Fatalf("Cannot read config: %s", err)
	}

	return &cfg
}
