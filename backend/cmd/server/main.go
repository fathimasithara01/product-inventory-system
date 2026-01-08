package main

import (
	"log"

	"github.com/fathimasithara01/product-inventory-system/config"
	"github.com/fathimasithara01/product-inventory-system/internal/handler"
	"github.com/fathimasithara01/product-inventory-system/internal/middleware"
	"github.com/fathimasithara01/product-inventory-system/internal/repository"
	"github.com/fathimasithara01/product-inventory-system/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()

	db := config.ConnectDB(cfg)

	productRepo := repository.NewProductRepository(db)
	variantRepo := repository.NewVariantRepository(db)
	subRepo := repository.NewSubVariantRepository(db)
	stockRepo := repository.NewStockRepository(db)

	productService := service.NewProductService(db, productRepo, variantRepo, subRepo)
	stockService := service.NewStockService(db, subRepo, stockRepo, productRepo)
	reportService := service.NewReportService(stockRepo)

	productHandler := handler.NewProductHandler(productService)
	stockHandler := handler.NewStockHandler(stockService)
	reportHandler := handler.NewReportHandler(reportService)

	router := gin.Default()

	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())

	api := router.Group("/api")
	{
		api.POST("/products", productHandler.CreateProduct)
		api.GET("/products", productHandler.ListProducts)

		api.POST("/stock/in", stockHandler.AddStock)
		api.POST("/stock/out", stockHandler.RemoveStock)

		api.GET("/reports/stock", reportHandler.StockReport)
	}

	log.Printf("Server running on port %s", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
