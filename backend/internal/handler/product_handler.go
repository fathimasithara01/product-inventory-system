package handler

import (
	"net/http"
	"strconv"

	"github.com/fathimasithara01/product-inventory-system/internal/models"
	"github.com/fathimasithara01/product-inventory-system/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(ps service.ProductService) *ProductHandler {
	return &ProductHandler{productService: ps}
}

// POST /products
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req struct {
		Product     models.Product         `json:"product"`
		Variants    []models.Variant       `json:"variants"`
		Options     []models.VariantOption `json:"options"`
		SubVariants []models.SubVariant    `json:"sub_variants"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.productService.CreateProduct(
		c.Request.Context(),
		&req.Product,
		req.Variants,
		req.Options,
		req.SubVariants,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
	})
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	products, total, err := h.productService.ListProducts(
		c.Request.Context(),
		page,
		limit,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  products,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}


// “Handlers only validate input and format responses.
// All business logic is in the service layer, which improves testability and maintainability.”