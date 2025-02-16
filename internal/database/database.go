package database

import (
	"fmt"
	"log"
	"qr-code-generator/config"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func Connect() {
	once.Do(func() {
		config := config.LoadDatabaseConfig()
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)

		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatalf("Failed to get DB instance: %v", err)
		}
		if err := sqlDB.Ping(); err != nil {
			log.Fatalf("Failed to ping the database: %v", err)
		}

		log.Println("Database connection established successfully")
	})
}

func GetDB() *gorm.DB {
	return DB
}
