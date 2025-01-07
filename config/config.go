package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL   string
	RedisURL      string
	ServerAddress string
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return &Config{
		DatabaseURL:   viper.GetString("database.url"),
		RedisURL:      viper.GetString("redis.url"),
		ServerAddress: viper.GetString("server.address"),
	}
}
