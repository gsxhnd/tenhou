package utils

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Mode             string `env:"MODE"`
	LogFileName      string `env:"LOG_FILE_NAME"`
	TenhouDBPath     string `env:"DB_PATH"`
	TenhouJsonDBPath string `env:"JSON_DB_PATH"`
	Listen           string `env:"LISTEN"`
}

func NewConfig() (*Config, error) {
	var c = Config{
		Mode:             "dev",
		LogFileName:      "./log/tenhou.log",
		TenhouDBPath:     "./data/tenhou.db",
		TenhouJsonDBPath: "./data/tenhou_json.db",
		Listen:           ":8080",
	}

	opts := env.Options{
		Prefix: "TENHOU_",
	}

	if err := env.ParseWithOptions(&c, opts); err != nil {
		return nil, err
	}

	fmt.Println("config:", c)

	return &c, nil
}
