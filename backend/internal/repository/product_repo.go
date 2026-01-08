package repository

import (
	"context"

	"github.com/fathimasithara01/product-inventory-system/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, tx *gorm.DB, product *models.Product) error
	List(ctx context.Context, offset, limit int) ([]models.Product, int64, error)
	FindByID(ctx context.Context, id string) (*models.Product, error)
	UpdateTotalStock(ctx context.Context, tx *gorm.DB, productID string) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, tx *gorm.DB, product *models.Product) error {
	return tx.WithContext(ctx).Create(product).Error
}

func (r *productRepository) List(ctx context.Context, offset, limit int) ([]models.Product, int64, error) {
	var (
		products []models.Product
		total    int64
	)

	err := r.db.WithContext(ctx).Model(&models.Product{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Order("created_date DESC").
		Find(&products).Error

	return products, total, err
}

func (r *productRepository) FindByID(ctx context.Context, id string) (*models.Product, error) {
	var product models.Product
	err := r.db.WithContext(ctx).First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) UpdateTotalStock(
	ctx context.Context,
	tx *gorm.DB,
	productID string,
) error {
	return tx.WithContext(ctx).Exec(`
		UPDATE products
		SET total_stock = (
			SELECT COALESCE(SUM(stock), 0)
			FROM sub_variants
			WHERE product_id = ?
		)
		WHERE id = ?
	`, productID, productID).Error
}
