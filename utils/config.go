package utils

import "github.com/caarlos0/env/v11"

type Config struct {
	Mode             string
	LogFileName      string
	TenhouDBPath     string
	TenhouJsonDBPath string
	Listen           string
}

func NewConfig() (*Config, error) {
	var c = Config{
		Mode:             "dev",
		LogFileName:      "./log/tenhou.log",
		TenhouDBPath:     "./data/tenhou.db",
		TenhouJsonDBPath: "./data/tenhou_json.db",
		Listen:           ":8080",
	}

	if err := env.Parse(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
