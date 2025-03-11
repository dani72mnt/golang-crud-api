package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

type ServerConfig struct {
	Port         string
	InternalPort string
	ExternalPort string
	RunMode      string
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func LoadConfig() *Config {
	var cfg Config

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("conf")
	v.AddConfigPath("../app/config")
	v.AutomaticEnv()

	// Debugging: Print where Viper is looking
	fmt.Println("Searching for config file in:", v.ConfigFileUsed())

	// Attempt to read the config
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return nil
	}

	// Unmarshal into the Config struct
	err := v.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("Error unmarshalling config: %v\n", err)
		return nil
	}

	fmt.Println("Config loaded successfully:", cfg)
	return &cfg
}
