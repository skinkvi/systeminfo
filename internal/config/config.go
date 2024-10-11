package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DBHost     string `yaml:"db_host"`
	DBPort     string `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
	DBURL      string `yaml:"db_url"`
	DBSSLMode  string `yaml:"db_sslmode"`
	GRPCPort   string `yaml:"grpc_port"`
	HTTPPort   string `yaml:"http_port"`
}

func LoadConfig(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
