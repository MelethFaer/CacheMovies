package database

import (
	"Cache/internal/models"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Connect *gorm.DB

type dbConfig struct {
	Host     string `env:"HOST"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	DBName   string `env:"DB_NAME"`
	Port     string `env:"PORT"`
	SSLMode  string `env:"SSL_MODE"`
	TimeZone string `env:"TIME_ZONE"`
}

func initConfig() *dbConfig {
	if err := godotenv.Load("pkg/database/dbConfig.env"); err != nil {
		log.Print("No .env file found")
	}

	config := &dbConfig{}
	if err := env.Parse(config); err != nil {
		log.Fatal(err)
	}

	return config
}

func Start() {
	config := initConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode, config.TimeZone)

	Connect, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err := Connect.AutoMigrate(&models.Movie{}); err != nil {
		log.Fatal(err)
	}

}
