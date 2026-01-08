package handler

import (
	"net/http"

	"github.com/fathimasithara01/product-inventory-system/internal/service"
	"github.com/shopspring/decimal"
	"github.com/gin-gonic/gin"
)

// StockHandler handles stock in/out operations
type StockHandler struct {
	stockService service.StockService
}

func NewStockHandler(ss service.StockService) *StockHandler {
	return &StockHandler{stockService: ss}
}

type stockRequest struct {
	ProductID    string          `json:"product_id" binding:"required"`
	SubVariantID string          `json:"sub_variant_id" binding:"required"`
	Quantity     decimal.Decimal `json:"quantity" binding:"required"`
}

func (h *StockHandler) AddStock(c *gin.Context) {
	var req stockRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Quantity.LessThanOrEqual(decimal.Zero) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "quantity must be greater than zero"})
		return
	}

	if err := h.stockService.AddStock(
		c.Request.Context(),
		req.ProductID,
		req.SubVariantID,
		req.Quantity,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock added successfully"})
}

func (h *StockHandler) RemoveStock(c *gin.Context) {
	var req stockRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Quantity.LessThanOrEqual(decimal.Zero) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "quantity must be greater than zero"})
		return
	}

	if err := h.stockService.RemoveStock(
		c.Request.Context(),
		req.ProductID,
		req.SubVariantID,
		req.Quantity,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock removed successfully"})
}
