package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
    Env string `yaml:"env" env-default:"local"`
    StoragePath string `yaml:"storage_path" env_required:"true"`
    TokenTTL time.Duration `yaml:"token_ttl" env_required:"true"`
    GRPC GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
    Port int `yaml:"port"`
    Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
    configPath := fetchConfigPath()

    if configPath == "" {
        panic("config path is empty")
    }

    if _, err := os.Stat(configPath); os.IsNotExist(err) {
        panic("config file doen not exist: " + configPath)
    }

    var cfg Config

    if err:=cleanenv.ReadConfig(configPath, &cfg); err != nil {
        panic("failed to red config: " + err.Error())
    }

    return &cfg;
}

func fetchConfigPath() string {
    var res string

    flag.StringVar(&res, "config", "", "path to config file")
    flag.Parse()

    if res == "" {
        res = os.Getenv("CONFIG_PATH")
    }

    return res
}
