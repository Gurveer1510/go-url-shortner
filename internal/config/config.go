package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PGHOST           string `mapstructure:"PGHOST"`
	PGDATABASE       string `mapstructure:"PGDATABASE"`
	PGUSER           string `mapstructure:"PGUSER"`
	PGPASSWORD       string `mapstructure:"PGPASSWORD"`
	PGSSLMODE        string `mapstructure:"PGSSLMODE"`
	PGCHANNELBINDING string `mapstructure:"PGCHANNELBINDING"`
	APP_PORT         string `mapstructure:"APP_PORT"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	envConfigFileName := ".env.local"
	viper.AutomaticEnv()
	viper.AddConfigPath("./.secrets")
	viper.SetConfigName(envConfigFileName)
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("config file not found. using env variables")
			return nil, err
		} else {
			return nil, fmt.Errorf("failed to read config file: %v", err)
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %v", err)
	}
	fmt.Println("config:", config)
	return config, nil
}
