package main

import (
	"log"

	"github.com/fathimasithara01/product-inventory-system/config"
	"github.com/fathimasithara01/product-inventory-system/database"
	"github.com/fathimasithara01/product-inventory-system/internal/middleware"
	"github.com/fathimasithara01/product-inventory-system/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	database.RunMigrations(db)

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.Use(gin.Logger(), gin.Recovery())

	routes.RegisterRoutes(r, db)

	log.Printf("Server running at http://localhost:%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
