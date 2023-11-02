package main

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type HttpServer struct {
	Port            int           `mapstructure:"port"`
	ShutdownTimeout time.Duration `mapstructure:"shutdownTimeout"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type DependenciesConfig struct {
	APIKey       string `mapstructure:"api_key"`
	OtherSetting string `mapstructure:"other_setting"`
}

type Config struct {
	HttpServer   HttpServer         `mapstructure:"httpServer"`
	Database     DatabaseConfig     `mapstructure:"database"`
	Dependencies DependenciesConfig `mapstructure:"dependencies"`
}

func main() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file: %s", err))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("Error unmarshaling config: %s", err))
	}

	fmt.Printf("HTTP Server Port: %d\n", config.HttpServer.Port)
	fmt.Printf("HTTP Server Shutdown Timeout: %s\n", config.HttpServer.ShutdownTimeout)

	fmt.Printf("Database Host: %s\n", config.Database.Host)
	fmt.Printf("Database Port: %d\n", config.Database.Port)
	fmt.Printf("Database Username: %s\n", config.Database.Username)
	fmt.Printf("Database Password: %s\n", config.Database.Password)
	fmt.Printf("Database Name: %s\n", config.Database.DBName)

	fmt.Printf("API Key: %s\n", config.Dependencies.APIKey)
	fmt.Printf("Other Setting: %s\n", config.Dependencies.OtherSetting)
}
