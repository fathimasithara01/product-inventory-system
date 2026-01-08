package routes

import (
	"github.com/fathimasithara01/product-inventory-system/internal/handler"
	"github.com/fathimasithara01/product-inventory-system/internal/repository"
	"github.com/fathimasithara01/product-inventory-system/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	productRepo := repository.NewProductRepository(db)
	subVariantRepo := repository.NewSubVariantRepository(db)
	stockRepo := repository.NewStockRepository(db)
	variantRepo := repository.NewVariantRepository(db)

	productService := service.NewProductService(db, productRepo, variantRepo, subVariantRepo)
	stockService := service.NewStockService(db, subVariantRepo, stockRepo, productRepo)
	reportService := service.NewReportService(stockRepo)

	productHandler := handler.NewProductHandler(productService)
	stockHandler := handler.NewStockHandler(stockService)
	reportHandler := handler.NewReportHandler(reportService)

	api := r.Group("/api")
	{
		api.POST("/products", productHandler.CreateProduct)
		api.GET("/products", productHandler.ListProducts)

		api.POST("/stock/in", stockHandler.AddStock)
		api.POST("/stock/out", stockHandler.RemoveStock)

		api.GET("/stock/report", reportHandler.StockReport)
	}
}
