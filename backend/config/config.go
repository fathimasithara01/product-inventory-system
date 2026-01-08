package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/fathimasithara01/product-inventory-system/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    string
}

var (
	cfg  *Config
	once sync.Once
)

// LoadConfig reads environment variables (singleton)
func LoadConfig() *Config {
	once.Do(func() {
		cfg = &Config{
			DBHost:     getEnv("DB_HOST", "localhost"),
			DBPort:     getEnv("DB_PORT", "5432"),
			DBUser:     getEnv("DB_USER", "postgres"),
			DBPassword: getEnv("DB_PASSWORD", "@postgres"),
			DBName:     getEnv("DB_NAME", "product_inventory"),
			AppPort:    getEnv("APP_PORT", "8080"),
		}
	})
	return cfg
}

// ConnectDB connects to PostgreSQL and auto-migrates tables
func ConnectDB(cfg *Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")

	// ⚡ Auto-migrate all models
	err = db.AutoMigrate(
		&models.Product{},
		&models.Variant{},
		&models.VariantOption{},
		&models.SubVariant{},
		&models.StockTransaction{},
	)

	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	log.Println("Database tables migrated successfully!")
	return db
}

// getEnv helper
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
