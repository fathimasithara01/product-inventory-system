package repository

import (
	"context"

	"github.com/fathimasithara01/product-inventory-system/internal/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, tx *gorm.DB, product *model.Product) error
	List(ctx context.Context, offset, limit int) ([]model.Product, int64, error)
	FindByID(ctx context.Context, id string) (*model.Product, error)
	UpdateTotalStock(ctx context.Context, tx *gorm.DB, productID string) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) Create(ctx context.Context, tx *gorm.DB, product *model.Product) error {
	return tx.WithContext(ctx).Create(product).Error
}

func (r *productRepo) List(ctx context.Context, offset, limit int) ([]model.Product, int64, error) {
	var (
		products []model.Product
		total    int64
	)

	if err := r.db.WithContext(ctx).Model(&model.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Order("created_date DESC").
		Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *productRepo) FindByID(ctx context.Context, id string) (*model.Product, error) {
	var product model.Product
	if err := r.db.WithContext(ctx).First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) UpdateTotalStock(ctx context.Context, tx *gorm.DB, productID string) error {
	return tx.WithContext(ctx).Exec(`
		UPDATE products
		SET total_stock = (
			SELECT COALESCE(SUM(stock), 0)
			FROM sub_variants
			WHERE product_id = ?
		)
		WHERE id = ?;
	`, productID, productID).Error
}
