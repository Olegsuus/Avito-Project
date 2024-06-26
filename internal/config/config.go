package config

import (
	"github.com/spf13/viper"
	"log"
)

type ServerConfig struct {
	Port     int `yaml:"port"`
	PageSize int `yaml:"pageSize"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	PageSize  int    `yaml:"PageSize"`
	JWTSecret string `yaml:"JWTSecret"`
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable decode into struct, %v", err)
	}

	return &cfg
}
