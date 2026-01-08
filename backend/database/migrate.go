package database

import (
	"log"

	"github.com/fathimasithara01/product-inventory-system/internal/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Product{},
		&model.Variant{},
		&model.VariantOption{},
		&model.SubVariant{},
		&model.StockTransaction{},
	)

	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("Database migrations completed!")
}
