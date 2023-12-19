package config

import (
	"time"
)

type Config struct {
	Env         string
	StoragePath string
	GRPC        GRPCConfig
}

type GRPCConfig struct {
	Port    int
	Timeout time.Duration
}

func New() *Config {
	cfg := Config{
		Env:         "local",
		StoragePath: "",
		GRPC: GRPCConfig{
			Port:    11111,
			Timeout: 5,
		},
	}

	return &cfg
}
