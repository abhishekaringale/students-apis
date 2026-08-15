package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-reuired:"true"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configpath string
	configpath = os.Getenv("CONFIG_PATH")

	if configpath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()
		if *flags != "" {
			configpath = *flags
		}

		if configpath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configpath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configpath, &cfg)
	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	return &cfg
}
